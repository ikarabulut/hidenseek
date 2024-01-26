package main

import (
	"fmt"
	"os"
    "errors"

    "github.com/ikarabulut/hidenseek/server"
)


func main() {
    dataFilePath := os.Getenv("DATA_FILE_PATH"); if dataFilePath == "" {
        fmt.Println("You have not specified a DATA_FILE_PATH, a default secrets file will be created in your applications root")
        dataFilePath = "./secrets.json"
    }

    handleFilePath(dataFilePath)

    server.RunServer()
}

func handleFilePath(filePath string) (secretsFile os.FileInfo, err error){
    file, err := os.Stat(filePath); if errors.Is(err, os.ErrNotExist) {
        file, err := createFile(filePath)
        return file, err
    }
    return file, err
}

func createFile(filePath string) (file os.FileInfo, err error){
    os.Create(filePath)
    return
}
