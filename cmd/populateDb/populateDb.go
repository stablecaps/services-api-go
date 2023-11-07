package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/stablecaps/services-api-go/pkg/models"
)





func main() {
	createNewServiceUrl := "http://localhost:8969/services/new"

	log.Printf("Creating new service by posting data to: %s", createNewServiceUrl)



	log.Println("Creating new service")

	body := []byte(`{
		"serviceName": "",
		"ServiceDescription": ""
	}`)

	req, err := http.NewRequest("POST", createNewServiceUrl, bytes.NewBuffer(body))
	if err != nil {
		log.Println("Error creating post request")
		panic(err)
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println("Error making post request")
		panic(err)
	}

	defer res.Body.Close()

	//
	post := &models.CreateServiceRequest{}
	derr := json.NewDecoder(res.Body).Decode(post)
	if derr != nil {
		log.Println("Error decoding post response")
		panic(derr)
	}

	if res.StatusCode != http.StatusCreated {
		log.Println("Error unexpected status")
		panic(res.Status)
	}

	fmt.Println("Id:", post.ServiceName)
	fmt.Println("Title:", post.ServiceDescription)
}