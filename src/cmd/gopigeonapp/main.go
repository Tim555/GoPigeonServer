package main

import (
	"fmt"
	"gopigeon/internal/server"
)

func main() {
	fmt.Println("Starting Go Pidgeon Server...")

	server.StartServer()
}
