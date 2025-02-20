package client

import (
	"Stats/internal/model"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Client(stat *model.Stats, hash string) error {
	url := "http://memcache:8084/Save"

	data, err := json.Marshal(&RequestClientPayload{
		Url: stat.Url,
		ShortUrl: fmt.Sprintf("http://memcache:8081/%s", hash),
	})
	if err != nil {
		return err
	}
	
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(body))
	return nil
}	
