package main

import (
	"fmt"
	"strconv"

	"github.com/fatih/color"
	"github.com/stablecaps/services-api-go/internal/dbtools"
)

// Declaring Global variables
var testsPassed int = 0
var testsFailed int = 0


func globalTestCounterPass(statusCode int)  {
	color.Green("Test passed: %d", statusCode)
	testsPassed ++
}


func globalTestCounterFail(statusCode int) {
	color.Red("Error!! unexpected status: %d", statusCode)
	testsFailed ++
}

// TODO: make these functions more modular
// func submitDeleteRequest(testName, baseURL, endpoint string, paramMap map[string]string, idx, wantedRespCode int) {
func submitDeleteRequest() {
	fmt.Println("\n~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("Posting sample service to test API delete endpoint")

	body := dbtools.MakeExplicitServiceJson("testDeleter2", "a service we will test delete", "v1,v2,v3")
	createNewServiceUrl := "http://localhost:8969/services/new"

	postedServiceData := dbtools.SubmitPostRequest(createNewServiceUrl, body)

	fmt.Println("\n~~~~~~~~~~~~~~~~~~~~")
	fmt.Printf("Now Deleting ServiceId %d via API delete endpoint", postedServiceData.ServiceId)


	// deletedServiceData := dbtools.SubmitDeleteRequest()
	// fmt.Printf("deletedServiceData: %v", deletedServiceData)


}

func main() {


	// Run tests
	// testListServices()

	// println("\n#######################")
	// println("#######################")
	// println("#######################")

	// testGetServiceById()

	// println("\n#######################")
	// println("#######################")
	// println("#######################")
	submitDeleteRequest()

	println("\n\n")
	color.Red("Tests failed: %s", strconv.Itoa(testsFailed))
	color.Green("Tests passed: %s", strconv.Itoa(testsPassed))

}