package main

import (
	"memCache/internal/server"
	"fmt"
)

func main() {
	server, err := server.New()
	if err != nil {
		panic(err)
	}

	fmt.Println("Running redis server on port 8084")
	server.Run()

}