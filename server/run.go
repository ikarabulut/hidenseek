package server

import (
	"log"
	"net/http"
	"encoding/json"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Status OK"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}

func RunServer() {
	mux := http.NewServeMux()

	health := http.HandlerFunc(healthCheckHandler)

	mux.Handle("/health", health)

	log.Print("Listening...")
	http.ListenAndServe(":3000", mux)
}	