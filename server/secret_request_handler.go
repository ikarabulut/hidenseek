package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SecretRequestModel struct {
	PlainText string `json:"plain_text"`
}

func ParseBody(r *http.Request) (requestModel SecretRequestModel) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	err := json.Unmarshal(body, &requestModel)
	if err != nil {
		fmt.Println("Error parsing request", err)
	}

	return requestModel
}
