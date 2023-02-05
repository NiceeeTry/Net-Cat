package main

import (
	// "net/client"
	"fmt"
	"net/TCPChat/server"
	"os"
)

func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Println("Listening on the port :8081")
		err := server.Server("8081")
		if err != nil {
			os.Exit(1)
		}
	} else if len(os.Args[1:]) == 1 {
		fmt.Println("Listening on the port :", os.Args[1])
		err := server.Server(os.Args[1])
		if err != nil {
			os.Exit(1)
		}
	} else {
		fmt.Println("[USAGE]: ./TCPChat $port")
	}

	// client.Client()
}
