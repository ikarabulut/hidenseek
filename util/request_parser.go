package util

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ikarabulut/hidenseek/types"
)

func ParseBody(r *http.Request) (requestModel types.SecretRequestModel) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	err := json.Unmarshal(body, &requestModel)
	if err != nil {
		fmt.Println("Error parsing request", err)
	}

	return requestModel
}
