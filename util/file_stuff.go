package util

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"sync"
)

type FileStore struct {
	Mu              sync.Mutex
	SecretsFilePath string
	Store           map[string]string
}

func (fStore *FileStore) Write(secret string, hash string) error {
	fStore.Mu.Lock()
	defer fStore.Mu.Unlock()

	err := fStore.ReadFromFile()
	if err != nil {
		return err
	}

	fStore.Store[hash] = secret

	return fStore.WriteToFile()

}

func (fStore *FileStore) Read(id string) (string, error) {
	fStore.Mu.Lock()
	defer fStore.Mu.Unlock()

	err := fStore.ReadFromFile()
	if err != nil {
		return "", err
	}

	data := fStore.Store[id]
	delete(fStore.Store, id)
	fStore.WriteToFile()

	return data, nil
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
