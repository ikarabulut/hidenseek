package util

import (
	"encoding/json"
	"log"
	"net/http"
)

type SecretResponse struct {
	Id     string
	Status int
	Secret string
}

func (response *SecretResponse) createHeaders(w http.ResponseWriter) {
	w.WriteHeader(response.Status)
	w.Header().Set("Content-Type", "application/json")
}

func (response *SecretResponse) createResponseBody(w http.ResponseWriter) {
	jsonResp, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	_, err = w.Write(jsonResp)
	if err != nil {
		log.Fatalf("Error writing to response. Err: %s", err)
	}
}

func (response *SecretResponse) BuildResponse(w http.ResponseWriter) {
	response.createHeaders(w)
	response.createResponseBody(w)
}
