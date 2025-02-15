package main

import (
	"Lynks/shortener/internal/server"
)

func main(){
	server, err := server.New()
	if err != nil {
		panic(err)
	}

	// logger
	server.NewLogger()
	server.Run()
}