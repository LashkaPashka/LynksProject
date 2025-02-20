package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetCache(hash string) (map[string]string, error) {
	var result map[string]string

	resp, err := http.Get(fmt.Sprintf("http://memcache:8084/Get/%s", hash))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal([]byte(string(body)), &result); err != nil {
		return nil, err
	}

	return result, nil
}