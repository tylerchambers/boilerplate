package main

import (
	"github.com/tylerchambers/boilerplate/app"
)

func main() {
	server, _ := app.NewServer("8080")

	server.Run()
}
