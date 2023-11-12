package main

import (
	"log"

	dbtools "github.com/stablecaps/services-api-go/internal/dbtools"
)

// Declaring Global variables
var baseURL string = "http://localhost:8969/"


func main() {
	endpoint := "/services/new"

	log.Printf("Creating new service by posting data to: %s%s", baseURL, endpoint)

	numServices := 50
	paramMapList := map[string]string{}


	for idx := 0; idx <= numServices; idx++ {
		log.Printf("Creating service no %d", idx)

		body := dbtools.MakeRandomService()
		dbtools.MakeHttpRequestWrapper(baseURL, endpoint, "POST", paramMapList, body)

	}

}

