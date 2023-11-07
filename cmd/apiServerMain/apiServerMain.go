package main

import (
	"log"

	"github.com/stablecaps/services-api-go/pkg/api"
	"github.com/stablecaps/services-api-go/pkg/models"
)

func main() {

	log.Println("Starting up database..")
	dataBase, err := models.NewPostgresDb()
	if err != nil {
		log.Fatal(err)
	}

	if err := dataBase.Init(); err != nil {
		log.Fatal(err)
	}

	log.Println("Starting up server..")
	server := api.NewAPIServer(":8969", dataBase)
	server.Run()

}

