package server

import (
	"log"
	"net/http"
	"sync"

	"github.com/ikarabulut/hidenseek/util"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	buildResponse(w, 200, "health-check")
}

func createSecret(secretsPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestModel := ParseBody(r)
		secretHex := util.CreateMd5Hex(requestModel.PlainText)
		fStore := util.FileStore{
			Mu:              sync.Mutex{},
			SecretsFilePath: secretsPath,
			Store:           make(map[string]string),
		}

		fStore.Write(requestModel.PlainText, secretHex)

		buildResponse(w, 200, secretHex)
	}
}

func RunServer(secretsFilePath string) {
	mux := http.NewServeMux()

	health := http.HandlerFunc(healthCheckHandler)
	secret := http.HandlerFunc(createSecret(secretsFilePath))

	mux.Handle("/health", health)
	mux.Handle("/secret", secret)

	log.Print("Listening...")
	http.ListenAndServe(":3000", mux)
}
