package main

import (
	"fmt"
	"gopidgeon/internal/server"
)

func main() {
	fmt.Println("Starting Go Pidgeon Server...")

	server.StartServer()
}
