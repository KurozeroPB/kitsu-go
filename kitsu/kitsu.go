package kitsu

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

const (
	baseURL   = "https://kitsu.io/api/edge"
	userAgent = "kitsu.go/v0.0.5 - (github.com/KurozeroPB/kitsu.go)"
)

func executeRequest(request *http.Request, expectedStatus int) ([]byte, error) {
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != expectedStatus {
		return nil, fmt.Errorf(
			"Expected status %d; Got %d \nResponse: %#v",
			expectedStatus,
			response.StatusCode,
			buf.String(),
		)
	}
	return buf.Bytes(), nil
}

func newRequest(method string, url string) *http.Request {
	return newUARequest(method, url, userAgent)
}

func newUARequest(method string, url string, ua string) *http.Request {
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("User-Agent", ua)
	request.Header.Set("Accept", "application/vnd.api+json")
	request.Header.Set("Content-Type", "application/vnd.api+json")
	return request
}

func safeGET(url string, expectedStatus int) ([]byte, error) {
	byt, err := executeRequest(newRequest("GET", url), expectedStatus)
	if err != nil {
		return nil, err
	}
	return byt, nil
}

func get(url string) ([]byte, error) {
	byt, err := safeGET(url, 200)
	if err != nil {
		return nil, err
	}
	return byt, nil
}
