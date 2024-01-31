package handlers

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/ikarabulut/hidenseek/types"
	"github.com/ikarabulut/hidenseek/util"
)

type SecretHandler struct {
	SecretsPath string
	Password    string
	Salt        string
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
		SecretsFilePath: sHandler.SecretsPath,
		Store:           make(map[string][]byte),
	}
	fStore.InitCrypto(sHandler.Password, sHandler.Salt)

	secret, err := fStore.Read(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResp, err := json.Marshal(types.GetSecretResponse{Secret: secret})
	if err != nil {
		http.Error(w, "Error writing body", http.StatusInternalServerError)
	}
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func (sHandler SecretHandler) createSecret(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength == 0 {
		http.Error(w, "Missing body param", http.StatusBadRequest)
		return
	}

	requestModel := util.ParseBody(w, r)
	if requestModel.PlainText == "" {
		http.Error(w, "Invalid body param", http.StatusBadRequest)
		return
	}

	secretHex := util.CreateMd5Hex(requestModel.PlainText)
	fStore := util.FileStore{
		Mu:              sync.Mutex{},
		SecretsFilePath: sHandler.SecretsPath,
		Store:           make(map[string][]byte),
	}
	fStore.InitCrypto(sHandler.Password, sHandler.Salt)

	fStore.Write(requestModel.PlainText, secretHex)

	jsonResp, err := json.Marshal(types.CreateSecretResponse{Id: secretHex})
	if err != nil {
		http.Error(w, "Error creating hash", http.StatusInternalServerError)
	}
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}
