package main

import (
	"log"

	dbtools "github.com/stablecaps/services-api-go/internal/populateDb"
)




func main() {
	createNewServiceUrl := "http://localhost:8969/services/new"

	log.Printf("Creating new service by posting data to: %s", createNewServiceUrl)

	numServices := 200


	for idx := 0; idx <= numServices; idx++ {
		log.Printf("Creating service no %d", idx)

		body := dbtools.MakeRandomService()
		dbtools.SubmitPostRequest(createNewServiceUrl, body)
	}

}

