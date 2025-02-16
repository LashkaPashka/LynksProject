package main

import (
	"Lynks/user/internal/server"
	"fmt"
)

func main() {
	server, err := server.New()
	if err != nil {
		panic(err)
	}

	//server.NewLogger()
	
	fmt.Println("Server' running :8082")
	server.Run()
}