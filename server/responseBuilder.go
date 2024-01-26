package server

import (
	"net/http"
	"encoding/json"
	"log"
)

// has to write the headers status code
// has to set the content type
// ++ any other headers
// has to create the response body. In this case JSON



func Create200ResponseBody(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	resp := make(map[string]string)
	resp["message"] = "Status OK"

	jsonResp, err := json.Marshal(resp)
		if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
}