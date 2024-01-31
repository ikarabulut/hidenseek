package handlers

import (
	"net/http"
)

func SetupHandlers(mux *http.ServeMux, secretsFilePath string) {
	mux.HandleFunc("/health", HealthCheckHandler)
	mux.Handle("/secret", SecretHandler{SecretsPath: secretsFilePath})
}
