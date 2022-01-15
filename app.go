package main

import (
	"fmt"
	"os"
)

func main() {

	server, err := InitServer()
	if err != nil {
		fmt.Println("could not init the server :", err)
		os.Exit(1)
	}

	fmt.Println("Starting Application...")
	fmt.Println("Started")
	server.Run()

}
