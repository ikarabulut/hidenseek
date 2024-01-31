package util

import (
	"encoding/json"
	"net/http"

	"github.com/ikarabulut/hidenseek/types"
)

func ParseBody(w http.ResponseWriter, r *http.Request) (requestModel types.SecretRequestModel) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	err := json.Unmarshal(body, &requestModel)
	if err != nil {
		http.Error(w, "Error reading body", http.StatusInternalServerError)
	}
	return requestModel
}
