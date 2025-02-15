package links

import "math/rand"

type Links struct {
	ID int
	Url string
	ShortUrl string
}

func NewLink(url string) *Links{
	link := &Links{
		Url: url,
	}
	link.Hash()

	return link
}

var letterRandom = []rune("abcdefghijklmnopqrstuvwxyzABCEFGHIJKLMNOPQRSTUVWXYZ")

func (l *Links) Hash() {
	l.ShortUrl = GenerateHash(10)
}

func GenerateHash(n int) string {
	var hash string
	for i := 0; i < n; i++ {
		hash += string(letterRandom[rand.Intn(len(letterRandom)-1)])
	}

	return hash
}