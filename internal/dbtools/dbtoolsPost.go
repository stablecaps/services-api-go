package dbtools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/stablecaps/services-api-go/pkg/models"
)

// PrettyPrint to print struct in a readable way
func PrettyPrint(i interface{}) string {
    s, _ := json.MarshalIndent(i, "", "\t")
    return string(s)
}

func MakeExplicitServiceJson(name, description, versions string) []byte {

	log.Printf("name is %s", name)
	log.Printf("description is %s", description)

	body := []byte(fmt.Sprintf(`{
		"serviceName": "%s",
		"serviceDescription": "%s",
		"serviceVersions": "%s"
	}`, name, description, versions) )

	return body
}

func SubmitPostRequest(url string, reqBody []byte) models.Service {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		log.Println("Error creating post request")
		os.Exit(42)
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error making post request")
		os.Exit(42)
	}

	defer resp.Body.Close()

	body, derr := io.ReadAll(resp.Body)
	if derr != nil {
		log.Printf("Error decoding post respponse: %s", err)
		log.Println(resp.Body)
		os.Exit(42)
	}
	fmt.Printf("string body: %s", string(body))

	if resp.StatusCode != http.StatusCreated {
		log.Printf("Error unexpected status: %d", resp.StatusCode)
		os.Exit(42)
	}

	// Unmarshal JSON to Go struct
	var result models.Service
	// Parse []byte to go struct pointer
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
		os.Exit(42)
	}
	fmt.Println(PrettyPrint(result))

	fmt.Println("ServiceId:", result.ServiceId)
	fmt.Println("Name:", result.ServiceName)
	fmt.Println("Descripton:", result.ServiceDescription)

	// os.Exit(42)

	return result
}

