package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/fatih/color"
)

// Declaring Global variables
var testsPassed int = 0
var testsFailed int = 0

// TODO: make these functions more modular
func submitGetRequest(testName, baseURL, endpoint string, paramMap map[string]string, idx, wantedRespCode int) {
	fmt.Println("\n~~~~~~~~~~~~~~~~~~~~")
	fmt.Printf("Running test %d: -  %s", idx, testName)

	u, _ := url.ParseRequestURI(baseURL)
	u.Path = endpoint

	// add params
	if len(paramMap) > 0 {
	params := url.Values{}
		for key, val := range paramMap {
			fmt.Println(key, val)
			params.Add(key, val)
		}
		u.RawQuery = params.Encode()
	}

	urlStr := fmt.Sprintf("%v", u)
	log.Printf("Calling services with the following url: %s", urlStr)

	resp, err := http.Get(urlStr)

	if err != nil {
		log.Printf("Request Failed: %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return
	}

	// Log the request body
	bodyString := string(body)
	log.Print(bodyString)

	if resp.StatusCode != wantedRespCode {
		color.Red("Error!! unexpected status: %d", resp.StatusCode)
		testsFailed ++
	} else {
		color.Green("Test passed")
		testsPassed ++
	}
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

	println("\n#######################")
	println("#######################")
	println("#######################")

	testGetServiceById()

	// Ran out of time for further tests
	// submitDeleteRequest()
	println("\n\n")

	color.Red("Tests failed: %s", strconv.Itoa(testsFailed))
	color.Green("Tests passed: %s", strconv.Itoa(testsPassed))

}