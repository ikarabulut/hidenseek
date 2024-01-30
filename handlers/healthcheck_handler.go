package handlers

import (
	"net/http"

	"github.com/ikarabulut/hidenseek/util"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := util.SecretResponse{
		Status: 200,
	}
	response.BuildResponse(w)
}
