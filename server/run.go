package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/ikarabulut/hidenseek/util"
)

type FileStore struct {
	mu              sync.Mutex
	secretsFilePath string
	Store           map[string]string
}

func (c *FileStore) updateMap(secret string, hash string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	f, err := os.Open(c.secretsFilePath)
	if err != nil {
		fmt.Println(err)
	}
	jsonData, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}
	if len(jsonData) != 0 {
		json.Unmarshal(jsonData, &c.Store)
	}

	c.Store[hash] = secret
	j, err := json.Marshal(c.Store)

	f, err = os.Create(c.secretsFilePath)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	_, err = f.Write(j)

}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	BuildResponse(w, 200, "health-check")
	return
}

func SecretHandler(secretsFilePath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open(secretsFilePath)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		fileC := FileStore{
			mu:              sync.Mutex{},
			secretsFilePath: secretsFilePath,
			Store:           make(map[string]string),
		}

		requestModel := ParseBody(r)
		secretHex := util.CreateMd5Hex(requestModel.PlainText)
		fileC.updateMap(requestModel.PlainText, secretHex)
		BuildResponse(w, 200, secretHex)
	}
}

func RunServer(secretsFilePath string) {
	mux := http.NewServeMux()

	health := http.HandlerFunc(HealthCheckHandler)
	secret := http.HandlerFunc(SecretHandler(secretsFilePath))

	mux.Handle("/health", health)
	mux.Handle("/secret", secret)

	log.Print("Listening...")
	http.ListenAndServe(":3000", mux)
}
