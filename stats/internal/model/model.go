package model

type Stats struct {
	ID int `json:"id"`
	Url string `json:"url"`
	Clicks uint `json:"click"`
	Average_length int `json:"average_length"`
}