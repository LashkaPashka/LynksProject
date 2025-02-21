package main

import (
	"Stats/internal/server"
)

func main(){
	server, err := server.New()
	if err != nil {
		panic(err)
	}

	server.NewLogger()
	server.KafkaConsumer()
	

	server.Run()
}