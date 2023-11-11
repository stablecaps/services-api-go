package main

import (
	"fmt"

	"github.com/stablecaps/services-api-go/internal/dbtools"
)

func testGetServiceByName() {
	baseURL := "http://localhost:8969/"
	listEndpoint := "/services/name/"

	goodService := dbtools.CreateExplicitService(dbtools.MakeRandomName(), dbtools.MakeRandomDescription(4), "v1,v2,v3")

	testNameSlice := []string{
		// Test serviceId 400s
		"badServiceName", "nonExistentName",
		// Test serviceId 200
		"goodServiceId",
	}
	wantedCodes := []int{500, 500, 200}
	paramMapList := map[string]string{}
	serviceIdList := []string{"10", "NonExistantService", goodService.ServiceName}


	for idx, testName := range testNameSlice {
		fmt.Println("\n~~~~~~~~~~~~~~~~~~~~")
		fmt.Printf("Running test %d: -  %s\n", idx, testName)
		fullListEndpoint := fmt.Sprintf("%s%s", listEndpoint, serviceIdList[idx])

		_, respStatusCode := dbtools.SubmitGetRequest(baseURL, fullListEndpoint, paramMapList)
		if respStatusCode == wantedCodes[idx] {
			globalTestCounterPass(respStatusCode)
		} else {
			globalTestCounterFail(respStatusCode)
		}

	}
}