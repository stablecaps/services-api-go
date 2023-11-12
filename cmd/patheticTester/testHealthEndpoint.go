package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/stablecaps/services-api-go/internal/dbtools"
)

func testGetHealthendpoint() {
	healthEndpoint := "/health"


	testNameSlice := []string{
		// Test health 200
		"goodCall",
	}
	wantedCodes := []int{200,}
	paramMapList := map[string]string{}


	for idx, testName := range testNameSlice {
		fmt.Println("\n~~~~~~~~~~~~~~~~~~~~")
		fmt.Printf("Running test %d: -  %s\n", idx, testName)

		resp, _ := dbtools.MakeHttpRequestWrapper(baseURL, healthEndpoint, "GET", paramMapList, nil)
		scoreGlobalTestsPassedandFailes(resp.StatusCode, wantedCodes[idx])
	}
}

func TestHealthCheck(t *testing.T) {

	response, err := http.Get("http://localhost:8969/health")
	if err != nil {
		t.Errorf("expected no errors, but got %v", err)
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("expected 200 statuscode, but got %v", response.StatusCode)
	}

	responseBody := ""
	json.NewDecoder(response.Body).Decode(&responseBody)
	response.Body.Close()

	if responseBody != "service is up and running" {
		t.Errorf(`expected message to be "service is up and running", but got %v`, responseBody)
	}

	os.Interrupt.Signal()
}
