package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/stablecaps/services-api-go/pkg/models"
	"github.com/tjarratt/babble"
)


func makeRandomName() string {
	babbler := babble.NewBabbler()
	babbler.Count = 1
	return babbler.Babble()
}

func makeRandomDescription(wordCount int) string {
	babbler := babble.NewBabbler()
	babbler.Separator = " "
	babbler.Count = wordCount
	return babbler.Babble()
}

func makeRandomService() []byte {
	randomName := makeRandomName()
	numWords := rand.Intn(10)
	radomDesc := makeRandomDescription(numWords)

	log.Printf("numWords is %d", numWords)
	log.Printf("randomName is %s", randomName)
	log.Printf("radomDesc is %s", radomDesc)

	body := []byte(fmt.Sprintf(`{
		"serviceName": "%s",
		"ServiceDescription": "%s"
	}`,randomName, radomDesc) )

	return body
}


func main() {
	createNewServiceUrl := "http://localhost:8969/services/new"

	log.Printf("Creating new service by posting data to: %s", createNewServiceUrl)

	numServices := 200

	for idx := 0; idx <= numServices; idx++ {
		log.Printf("Creating service no %d", idx)
		body := makeRandomService()
		submitPostRequest(createNewServiceUrl, body)
	}

}

func submitPostRequest(url string, reqBody []byte) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		log.Println("Error creating post request")
		os.Exit(42)
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println("Error making post request")
		os.Exit(42)
	}

	defer res.Body.Close()

	//
	post := &models.CreateServiceRequest{}
	derr := json.NewDecoder(res.Body).Decode(post)
	if derr != nil {
		log.Printf("Error decoding post response: %s", err)
		log.Println(res.Body)
		os.Exit(42)
	}

	if res.StatusCode != http.StatusCreated {
		log.Printf("Error unexpected status: %d", res.StatusCode)
		os.Exit(42)
	}

	fmt.Println("Id:", post.ServiceName)
	fmt.Println("Title:", post.ServiceDescription)
}