package main

import (
	"fmt"
	"log"
	"os"

	"github.com/stablecaps/services-api-go/pkg/api"
	"github.com/stablecaps/services-api-go/pkg/config"
	"github.com/stablecaps/services-api-go/pkg/models"
)

func main() {

	log.Println("Reading config..")


	config, err := config.Readconfig("config_dev_secrets", "env")
    if err != nil {
        fmt.Printf("Error: cannot read config: %s", err)
		os.Exit(42)
    }

	log.Printf("Config: %s", config.DBUser)

	log.Println("Starting up database..")
	dataBase, err := models.NewPostgresDb(config.DBUser, config.DBName, config.DBPassword, config.DBSSLmode)
	if err != nil {
		fmt.Printf("Error connecting to database: %s", err)
		os.Exit(42)
	}

	if err := dataBase.Init(); err != nil {
		fmt.Printf("Error creating DB table: %s", err)
		os.Exit(42)
	}

	log.Println("Starting up server..")
	port := fmt.Sprintf(":%s",  config.APIPort)
	server := api.NewAPIServer(port, dataBase)
	server.Run()

}

