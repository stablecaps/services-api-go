package main

import (
	"fmt"
	"log"
	"os"

	"github.com/stablecaps/services-api-go/pkg/api"
	"github.com/stablecaps/services-api-go/pkg/config"
	"github.com/stablecaps/services-api-go/pkg/models"
	_ "github.com/stablecaps/services-api-go/swagger"
)

//	@title			Service Catalog Dashboard API
//	@version		1.0
//	@description	Services API for Dashboard widget
//	@termsOfService	http://swagger.io/terms/

//	@securityDefinitions.basic	BasicAuth

//	@contact.name	Giri
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	log.Println("Reading config..")


	config, err := config.Readconfig("config_app_secrets", "env")
    if err != nil {
        fmt.Printf("Error: cannot read config: %s", err)
		os.Exit(42)
    }

	log.Printf("Config: %s", config.DBUser)

	log.Println("Starting up database..")
	dataBase, err := models.NewPostgresDb(config.DBUser,
										  config.DBName,
										  config.DBPassword,
										  config.DBSSLmode,
										  config.DBMaxOpenConns,
										  config.DBMaxOpenConns)
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
	server.Run(port)

}

