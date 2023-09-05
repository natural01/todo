package main

import (
	"ToDo/pkg"
	"fmt"
)

func main() {
	handler := new(pkg.Handler)
	server := new(pkg.Server)
	err := server.Run(handler.NewHandler())
	if err != nil {
		fmt.Printf("Server not running (%s)", err.Error())
	}
}
