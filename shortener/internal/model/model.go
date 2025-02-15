package model

import (
	"context"
	"math/rand"
)
type myString string
const (
	HostAPI myString = "hostAPI"
)


type Links struct {
	ID int
	Url string
	ShortUrl string
}

func NewLink(ctx context.Context, url string) *Links{
	link := &Links{
		Url: url,
	}
	link.Hash(ctx)

	return link
}

var letterRandom = []rune("abcdefghijklmnopqrstuvwxyzABCEFGHIJKLMNOPQRSTUVWXYZ")

func (l *Links) Hash(ctx context.Context) {
	l.ShortUrl = ctx.Value(HostAPI).(string) + GenerateHash(10)
}

func GenerateHash(n int) string {
	var hash string
	for i := 0; i < n; i++ {
		hash += string(letterRandom[rand.Intn(len(letterRandom)-1)])
	}

	return hash
}