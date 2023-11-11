package dbtools

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/stablecaps/services-api-go/pkg/models"
)


func SubmitDeleteRequest(url string, reqBody []byte) models.Service {

	resp, body := MakeHttpRequestWrapper(url, "DELETE", nil)

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

	return result
}
