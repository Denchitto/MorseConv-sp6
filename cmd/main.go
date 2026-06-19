package main

import (
	"MorseConv-sp6/internal/server"
	"log"
)

func main() {
	var logger *log.Logger
	newServer := server.MakeRouter(logger)
	err := newServer.Server.ListenAndServe()
	if err != nil {
		newServer.Logger.Fatal(err)
	}
}
