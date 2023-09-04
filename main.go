package main

import "fmt"

func main() {
	handler := new(Handler)
	server := new(Server)
	err := server.Run(handler.NewHandler())
	if err != nil {
		fmt.Printf("Server not running (%s)", err.Error())
	}
}
