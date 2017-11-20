package kitsu

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

const (
	baseURL   = "https://kitsu.io/api/edge"
	userAgent = "kitsu.go/v0.0.6 - (github.com/KurozeroPB/kitsu.go)"
)

func get(url string) ([]byte, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("User-Agent", userAgent)
	request.Header.Set("Accept", "application/vnd.api+json")
	request.Header.Set("Content-Type", "application/vnd.api+json")

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

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("Expected status %d; Got %d\nResponse: %#v", 200, response.StatusCode, buf.String())
	}

	return buf.Bytes(), nil
}
