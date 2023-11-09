package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/fatih/color"
)


func submitGetRequest(testName, baseURL, endpoint string, paramMap map[string]string, idx, wantedRespCode int) {
	fmt.Println("\n~~~~~~~~~~~~~~~~~~~~")
	fmt.Printf("Running test %d: -  %s", idx, testName)
	// add params
	params := url.Values{}
	for key, val := range paramMap {
		fmt.Println(key, val)
		params.Add(key, val)
	}

	u, _ := url.ParseRequestURI(baseURL)
	u.Path = endpoint
	u.RawQuery = params.Encode()
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
		//os.Exit(42)
	} else {
		color.Green("Test passed")
	}
}


func main() {

	// list services pathetic tests
	baseURL := "http://localhost:8969/"
	listEndpoint := "/services"


	testNameSlice := []string{
		// Test 400s
		"badLimit", "outOfRangeLimit", "badOffset", "outOfRangeOffset", "badOrderColName", "badOrderDir",
		// Test orderDir 200
		"orderDirasc", "orderDirdesc",
		// Test orderColName 200
		"orderColNameServiceId", "orderColNameServiceName", "orderColNameServiceDescription", "orderColNameServiceVersions", "orderColNameCreatedAt",
	}
	wantedCodes := []int{400, 400, 400, 400, 400, 400, 200, 200, 200, 200, 200, 200, 200}
	paramMapList := []map[string]string{
		// Test 400s
		{
			"limit": "fake",
			"offset": "0",
			"orderColName": "serviceName",
			"orderDir": "asc",
    	},
		{
			"limit": "-10",
			"offset": "0",
			"orderColName": "serviceName",
			"orderDir": "asc",
    	},
		{
			"limit": "4",
			"offset": "fake",
			"orderColName": "serviceName",
			"orderDir": "asc",
    	},
		{
			"limit": "4",
			"offset": "-10",
			"orderColName": "serviceName",
			"orderDir": "asc",
    	},
		{
			"limit": "4",
			"offset": "0",
			"orderColName": "fake",
    	},
		{
			"limit": "4",
			"offset": "0",
			"orderColName": "serviceName",
			"orderDir": "fake",
    	},
		// Test orderColName 200
		{
			"limit": "4",
			"offset": "0",
			"orderColName": "serviceId",
			"orderDir": "asc",
    	},
		{
			"limit": "4",
			"offset": "0",
			"orderColName": "serviceName",
			"orderDir": "asc",
    	},
		{
			"limit": "4",
			"offset": "0",
			"orderColName": "serviceDescription",
			"orderDir": "asc",
    	},
		{
			"limit": "4",
			"offset": "0",
			"orderColName": "serviceVersions",
			"orderDir": "asc",
    	},
		{
			"limit": "4",
			"offset": "0",
			"orderColName": "createdAt",
			"orderDir": "asc",
    	},
		// Test orderDir 200
		{
			"limit": "4",
			"offset": "0",
			"orderColName": "serviceName",
			"orderDir": "asc",
    	},

		{
			"limit": "4",
			"offset": "0",
			"orderColName": "serviceName",
			"orderDir": "asc",
    	},
	}


	for idx, testName := range testNameSlice {
		submitGetRequest(testName, baseURL, listEndpoint, paramMapList[idx], idx, wantedCodes[idx])
	}

}