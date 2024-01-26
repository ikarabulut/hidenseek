package server

import (
	"log"
	"net/http"

	"github.com/ikarabulut/hidenseek/util"
)



func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	BuildResponse(w, 200, "health-check")
	return
}

func SecretHandler(w http.ResponseWriter, r *http.Request) {
	requestModel := ParseBody(r)
	secretHex := util.CreateMd5Hex(requestModel.PlainText)
	BuildResponse(w, 200, secretHex)
	return
}

func RunServer() {
	mux := http.NewServeMux()

	health := http.HandlerFunc(HealthCheckHandler)
	secret := http.HandlerFunc(SecretHandler)

	mux.Handle("/health", health)
	mux.Handle("/secret", secret)

	log.Print("Listening...")
	http.ListenAndServe(":3000", mux)
}	