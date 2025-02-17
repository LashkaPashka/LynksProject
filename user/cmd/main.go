package main

import (
	"Lynks/user/internal/server"
)

func main() {
	server, err := server.New()
	if err != nil {
		panic(err)
	}

	server.NewLogger()
	server.Run()
}