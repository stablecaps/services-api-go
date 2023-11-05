package main

import "log"

func main() {
	log.Println("Starting up database..")
	store, error := NewPostgreStore()
	if error != nil {
		log.Fatal(error)
	}

	log.Println("Starting up server..")
	server := NewAPIServer(":8969", store)
	server.Run()

}

