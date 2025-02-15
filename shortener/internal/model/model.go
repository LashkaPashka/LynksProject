package model

import (
	"math/rand"
)

type Links struct {
	ID int `json:"id"`
	Url string `json:"url"`
	Hash string `json:"hash"`
}

func NewLink(url string) *Links{
	link := &Links{
		Url: url,
	}
	link.RHash()

	return link
}

var letterRandom = []rune("abcdefghijklmnopqrstuvwxyzABCEFGHIJKLMNOPQRSTUVWXYZ")

func (l *Links) RHash() {
	l.Hash = GenerateHash(10)
}

func GenerateHash(n int) string {
	var hash string
	for i := 0; i < n; i++ {
		hash += string(letterRandom[rand.Intn(len(letterRandom)-1)])
	}

	return hash
}