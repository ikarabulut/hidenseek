package server

import (
	"log"
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	BuildResponse(w, 200)
	return
}

func RunServer() {
	mux := http.NewServeMux()

	health := http.HandlerFunc(HealthCheckHandler)

	mux.Handle("/health", health)

	log.Print("Listening...")
	http.ListenAndServe(":3000", mux)
}	