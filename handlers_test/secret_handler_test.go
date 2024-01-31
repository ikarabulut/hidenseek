package handlersTest

import (
	"fmt"
	"io"
	"log"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ikarabulut/hidenseek/handlers"
)

func TestHandlers_CreateSecret(t *testing.T) {
	tempDir := filepath.Join(t.TempDir(), "test_data.json")
	_, err := os.Create(tempDir)
	if err != nil {
		log.Println(err)
	}
	testSetup := []struct {
		requestBody            string
		expectedHTTPStatusCode int
		expectedBody           string
	}{
		{ // When proper body is passed
			requestBody:            "{\"plain_text\": \"super secret\"}",
			expectedHTTPStatusCode: 200,
			expectedBody:           fmt.Sprintf("{\"Id\":\"%s\"}", "5f1903f5f2cb32acb4c1dcae9e30d374"),
		},
	}

	for _, tc := range testSetup {
		req := httptest.NewRequest("POST", "/secret", strings.NewReader(tc.requestBody))
		w := httptest.NewRecorder()

		sHandler := handlers.SecretHandler{
			SecretsPath: tempDir,
		}
		sHandler.ServeHTTP(w, req)
		response := w.Result()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			t.Fatal(err)
		}
		if response.StatusCode != tc.expectedHTTPStatusCode {
			t.Errorf("Expected Response Status to be: %d, Got: %d", tc.expectedHTTPStatusCode, response.StatusCode)

		}
		if string(body) != tc.expectedBody {
			t.Errorf("Expected Response body to be: %s, Got: %s", tc.expectedBody, string(body))
		}
	}
}
