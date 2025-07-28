package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
)

func PrettyPrint(s interface{}) {
	b, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		fmt.Println("PrettyPrint error:", err)
		return
	}
	fmt.Println(string(b))
}

func MakeJSONRequest(method, url string, payload interface{}) (*http.Request, *httptest.ResponseRecorder, error) {
	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, nil, err
	}

	req := httptest.NewRequest(method, url, bytes.NewBuffer(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	// Recorder captures the response
	w := httptest.NewRecorder()

	return req, w, nil
}

// MakeGETRequestWithQuery creates a GET request with query parameters for testing
func MakeGETRequestWithQuery(basePath string, queryParams map[string]string) (*http.Request, *httptest.ResponseRecorder) {
	u, _ := url.Parse(basePath)
	q := u.Query()

	for key, value := range queryParams {
		q.Set(key, value)
	}

	u.RawQuery = q.Encode()
	req := httptest.NewRequest(http.MethodGet, u.String(), nil)
	w := httptest.NewRecorder()
	return req, w
}
