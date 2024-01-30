package handlers

import (
	"net/http"
	"sync"

	"github.com/ikarabulut/hidenseek/util"
)

type SecretHandler struct {
	secretsPath string
}

func (sHandler SecretHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		sHandler.createSecret(w, r)
	case "GET":
		sHandler.getSecret(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (sHandler SecretHandler) getSecret(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("hash")
	if id == "" {
		http.Error(w, "No Secret ID specified", http.StatusBadRequest)
		return
	}
	fStore := util.FileStore{
		Mu:              sync.Mutex{},
		SecretsFilePath: sHandler.secretsPath,
		Store:           make(map[string]string),
	}

	secret, err := fStore.Read(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := util.SecretResponse{
		Id:     id,
		Status: 200,
		Secret: secret,
	}

	response.BuildResponse(w)
}

func (sHandler SecretHandler) createSecret(w http.ResponseWriter, r *http.Request) {
	requestModel := util.ParseBody(r)
	secretHex := util.CreateMd5Hex(requestModel.PlainText)
	fStore := util.FileStore{
		Mu:              sync.Mutex{},
		SecretsFilePath: sHandler.secretsPath,
		Store:           make(map[string]string),
	}

	fStore.Write(requestModel.PlainText, secretHex)

	response := util.SecretResponse{
		Id:     secretHex,
		Status: 200,
	}
	response.BuildResponse(w)
}
