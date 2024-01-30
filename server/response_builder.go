package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func createHeaders(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
}

func createResponseBody(w http.ResponseWriter, statusCode int, secretHash string) {
	resp := make(map[string]string)
	resp["id"] = secretHash
	resp["status"] = http.StatusText(statusCode)

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	_, err = w.Write(jsonResp)
	if err != nil {
		log.Fatalf("Error writing to response. Err: %s", err)
	}
}

func buildResponse(w http.ResponseWriter, statusCode int, secretHash string) {
	createHeaders(w, statusCode)
	createResponseBody(w, statusCode, secretHash)
}
