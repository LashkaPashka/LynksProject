package payload

type LinkCachePayload struct {
	Url string `json:"url"`
	ShortUrl string `json:"short_url"`
}

type ResponseLinkPayload struct {
	Url string `json:"url"`
	ShortUrl string `json:"short_url"`
}