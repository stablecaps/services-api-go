package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)


func submitGetRequest(testName, baseURL, endpoint string, paramMap map[string]string, wantedRespCode int) {
	fmt.Println("\n~~~~~~~~~~~~~~~~~~~~")
	fmt.Printf("Running test: %s", testName)
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
		log.Printf("Error unexpected status: %d", resp.StatusCode)
		os.Exit(42)
	} else {
		println("Test passed")
	}
}


func main() {

	// list services pathethic tests
	baseURL := "http://localhost:8969/"
	listEndpoint := "/services"


	testNameSlice := []string{"badOffset", "badOrderDir"}
	wantedCodes := []int{400,400}

	paramMapList := []map[string]string{
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
	}


	for idx, testName := range testNameSlice {
		submitGetRequest(testName, baseURL, listEndpoint, paramMapList[idx], wantedCodes[idx])
	}
	// params := url.Values{}
	// limit := "4"
	// offset := "0"
	// params.Add("limit", limit)
	// params.Add("offset", offset)
	// params.Add("orderColName", "serviceName")

	// u, _ := url.ParseRequestURI(baseURL)
	// u.Path = resource
	// u.RawQuery = params.Encode()
	// urlStr := fmt.Sprintf("%v", u)
	// log.Printf("Calling services with the following url: %s", urlStr)

	// resp, err := http.Get(urlStr)

	// if err != nil {
	// 	log.Printf("Request Failed: %s", err)
	// 	return
	//  }
	//  defer resp.Body.Close()
	//  body, err := io.ReadAll(resp.Body)
	//  if err != nil {
	// 	log.Printf("Reading body failed: %s", err)
	// 	return
	//  }
	//  // Log the request body
	//  bodyString := string(body)
	//  log.Print(bodyString)

	//  if resp.StatusCode != 200 {
	// 	log.Printf("Error unexpected status: %d", resp.StatusCode)
	// 	os.Exit(42)
	// }
}