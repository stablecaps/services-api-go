package main

import "log"

func main() {
	log.Println("Starting up server..")
	server := NewAPIServer(":8969")
	server.Run()

}

