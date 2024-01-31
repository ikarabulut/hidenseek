package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"io"
	"log"
	"os"
	"sync"

	"golang.org/x/crypto/scrypt"
)

type FileStore struct {
	Mu              sync.Mutex
	SecretsFilePath string
	GCM             cipher.AEAD
	Nonce           []byte
	Store           map[string][]byte
}

func (fStore *FileStore) Write(secret string, hash string) error {
	fStore.Mu.Lock()
	defer fStore.Mu.Unlock()

	err := fStore.ReadFromFile()
	if err != nil {
		return err
	}

	fStore.Store[hash] = fStore.encrypt(secret)

	return fStore.WriteToFile()

}

func (fStore *FileStore) Read(id string) (string, error) {
	fStore.Mu.Lock()
	defer fStore.Mu.Unlock()

	err := fStore.ReadFromFile()
	if err != nil {
		return "", err
	}

	data, err := fStore.decrypt(fStore.Store[id])
	if err != nil {
		log.Println(err)
		return "", err
	}
	delete(fStore.Store, id)
	fStore.WriteToFile()

	return string(data), nil
}

func (fStore *FileStore) ReadFromFile() error {
	file, err := os.Open(fStore.SecretsFilePath)
	if err != nil {
		log.Fatal(err)
	}
	jsonData, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	if len(jsonData) != 0 {
		return json.Unmarshal(jsonData, &fStore.Store)
	}
	return nil
}

func (fStore *FileStore) WriteToFile() error {
	var file *os.File
	jsonData, err := json.Marshal(fStore.Store)
	if err != nil {
		return err
	}
	file, err = os.Create(fStore.SecretsFilePath)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(jsonData)
	return err
}

func (fStore *FileStore) InitCrypto(password, salt string) error {
	key, err := scrypt.Key([]byte(password), []byte(salt), 32768, 8, 1, 32)
	if err != nil {
		return err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}
	nonce := make([]byte, gcm.NonceSize())
	fStore.GCM = gcm
	fStore.Nonce = nonce
	return nil
}

func (fStore *FileStore) encrypt(data string) []byte {
	if _, err := io.ReadFull(rand.Reader, fStore.Nonce); err != nil {
		log.Fatal(err)
	}
	encryptedData := fStore.GCM.Seal(fStore.Nonce, fStore.Nonce, []byte(data), nil)
	return encryptedData
}

func (fStore *FileStore) decrypt(encData []byte) ([]byte, error) {
	nonce := encData[:fStore.GCM.NonceSize()]
	encData = encData[fStore.GCM.NonceSize():]
	data, err := fStore.GCM.Open(nil, nonce, encData, nil)
	if err != nil {
		return nil, err
	}
	return data, err
}
