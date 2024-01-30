package server

import (
	"log"
	"net/http"
	"sync"

	"github.com/ikarabulut/hidenseek/util"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := secretResponse{
		Status: 200,
	}
	response.buildResponse(w)
}

type secretHandler struct {
	secretsPath string
}

func (sHandler secretHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

func (sHandler secretHandler) getSecret(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("hash")
	log.Println(id)
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

	response := secretResponse{
		Id:     id,
		Status: 200,
		Secret: secret,
	}

	response.buildResponse(w)
}

func (sHandler secretHandler) createSecret(w http.ResponseWriter, r *http.Request) {
	requestModel := ParseBody(r)
	secretHex := util.CreateMd5Hex(requestModel.PlainText)
	fStore := util.FileStore{
		Mu:              sync.Mutex{},
		SecretsFilePath: sHandler.secretsPath,
		Store:           make(map[string]string),
	}

	fStore.Write(requestModel.PlainText, secretHex)

	response := secretResponse{
		Id:     secretHex,
		Status: 200,
	}
	response.buildResponse(w)
}

func RunServer(secretsFilePath string) {
	mux := http.NewServeMux()

	health := http.HandlerFunc(healthCheckHandler)
	sHandler := secretHandler{secretsPath: secretsFilePath}

	mux.Handle("/health", health)
	mux.Handle("/secret", sHandler)

	log.Print("Listening...")
	http.ListenAndServe(":3000", mux)
}
