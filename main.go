package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/ikarabulut/hidenseek/handlers"
)

func main() {
	dataFilePath := os.Getenv("DATA_FILE_PATH")
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	if dataFilePath == "" {
		fmt.Println("You have not specified a DATA_FILE_PATH, a default secrets file will be created in your applications root")
		dataFilePath = "/secrets.json"
		var sb strings.Builder
		sb.WriteString(wd)
		sb.WriteString(dataFilePath)
		dataFilePath = sb.String()
	}

	handleFilePath(dataFilePath)

	mux := http.NewServeMux()
	handlers.SetupHandlers(mux, dataFilePath)

	log.Print("Listening...")
	http.ListenAndServe(":3000", mux)
}

func handleFilePath(filePath string) (secretsFile os.FileInfo, err error) {
	file, err := os.Stat(filePath)
	if errors.Is(err, os.ErrNotExist) {
		file, err := createFile(filePath)
		return file, err
	}
	return file, err
}

func createFile(filePath string) (file os.FileInfo, err error) {
	os.Create(filePath)
	return
}
