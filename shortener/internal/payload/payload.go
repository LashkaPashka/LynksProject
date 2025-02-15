package payload

type LinkRequest struct {
	Destination string `json:"destination"`
}

type LinkResponse struct {
	ShortUrl string `json:"shortUrl"`
	Destination string `json:"destination"`
}