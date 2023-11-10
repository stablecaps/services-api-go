package main

import (
	"strconv"

	"github.com/fatih/color"
	_ "github.com/stablecaps/services-api-go/internal/dbtools"
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
// func submitDeleteRequest() {
// 	fmt.Println("\n~~~~~~~~~~~~~~~~~~~~")
// 	// 8fmt.Printf("Running test %d: -  %s", idx, testName)


// 	// testTime := time.Now().UTC()
// 	body := dbtools.MakeExplicitService("testDeleter", "a service we will test delete", "v1,v2,v3")
// 	createNewServiceUrl := "http://localhost:8969/services/new"
// 	dbtools.SubmitExplicitPostRequest(createNewServiceUrl, body)
// }

func main() {

	// Run tests
	testListServices()

	// println("\n#######################")
	// println("#######################")
	// println("#######################")

	testGetServiceById()

	// Ran out of time for further tests
	// submitDeleteRequest()
	println("\n\n")
	color.Red("Tests failed: %s", strconv.Itoa(testsFailed))
	color.Green("Tests passed: %s", strconv.Itoa(testsPassed))

}