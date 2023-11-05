package main

import "log"

func main() {
	log.Println("Starting up database..")
	dataBase, err := NewPostgresDb()
	if err != nil {
		log.Fatal(err)
	}

	if err := dataBase.Init(); err != nil {
		log.Fatal(err)
	}

	log.Println("Starting up server..")
	server := NewAPIServer(":8969", dataBase)
	server.Run()

}

