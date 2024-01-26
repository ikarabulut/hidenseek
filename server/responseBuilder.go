package server

import (
	"net/http"
	"encoding/json"
	"log"
)

func CreateHeaders(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	return
}

func CreateResponseBody(w http.ResponseWriter, statusCode int) {
	resp := make(map[string]string)
	resp["status"] = http.StatusText(statusCode)

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	_, err = w.Write(jsonResp)
	if err != nil {
		log.Fatalf("Error writing to response. Err: %s", err)
	}
	return
}

func BuildResponse(w http.ResponseWriter, statusCode int) {
	CreateHeaders(w, statusCode)
	CreateResponseBody(w, statusCode)
	return
}