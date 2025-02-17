package main

import (
	"Lynks/stats/internal/server"
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