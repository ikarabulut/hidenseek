package server

import (
	"log"
	"net/http"
	"sync"

	"github.com/ikarabulut/hidenseek/util"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	BuildResponse(w, 200, "health-check")
	return
}

func CreateSecret(secretsPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestModel := ParseBody(r)
		secretHex := util.CreateMd5Hex(requestModel.PlainText)
		fStore := util.FileStore{
			Mu:              sync.Mutex{},
			SecretsFilePath: secretsPath,
			Store:           make(map[string]string),
		}

		fStore.Write(requestModel.PlainText, secretHex)

		BuildResponse(w, 200, secretHex)
	}
}

func RunServer(secretsFilePath string) {
	mux := http.NewServeMux()

	health := http.HandlerFunc(HealthCheckHandler)
	secret := http.HandlerFunc(CreateSecret(secretsFilePath))

	mux.Handle("/health", health)
	mux.Handle("/secret", secret)

	log.Print("Listening...")
	http.ListenAndServe(":3000", mux)
}
