package payload

type LinkRequest struct {
	Destination string `json:"destination"`
}

type LinkResponse struct {
	ShortUrl string `json:"shortUrl"`
	Destination string `json:"destination"`
}

type ClientPayload struct {
	Url string `json:"url"`
	ShortUrl string `json:"short_url"`
}