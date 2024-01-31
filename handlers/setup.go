package handlers

import (
	"net/http"
)

func SetupHandlers(mux *http.ServeMux, secretsFilePath string, password string, salt string) {
	mux.HandleFunc("/health", HealthCheckHandler)
	mux.Handle("/secret", SecretHandler{SecretsPath: secretsFilePath, Password: password, Salt: salt})
}
