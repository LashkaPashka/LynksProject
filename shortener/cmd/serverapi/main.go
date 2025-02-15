package main

import "Lynks/shortener/internal/server"


func main(){
	server, err := server.New()
	if err != nil {
		panic(err)
	}

	server.Run()
}