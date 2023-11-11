package dbtools

import (
	"encoding/json"
	"fmt"
	"log"
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

func SubmitPostRequest(url string, reqBody []byte) (models.Service, int) {

	resp, body := MakeHttpRequestWrapper(url, "POST", reqBody)

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

	return result, resp.StatusCode
}

