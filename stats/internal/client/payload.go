package client

type RequestClientPayload struct {
	Url string `json:"url"`
	ShortUrl string `json:"short_url"`
}