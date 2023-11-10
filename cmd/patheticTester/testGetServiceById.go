package main

import (
	"fmt"

	"github.com/stablecaps/services-api-go/internal/dbtools"
)

// TODO move this to tests
// TODO: make these functions more modular
func testGetServiceById() {
	baseURL := "http://localhost:8969/"
	listEndpoint := "/services/id/"


	testNameSlice := []string{
		// Test serviceId 400s
		"badServiceId", "outOfRangeServiceId",
		// Test serviceId 200
		"goodServiceId",
	}
	wantedCodes := []int{404, 404, 200}
	paramMapList := map[string]string{}
	serviceIdList := []string{"fake", "-10", "5",}


	for idx, testName := range testNameSlice {
		fmt.Println("\n~~~~~~~~~~~~~~~~~~~~")
		fmt.Printf("Running test %d: -  %s", idx, testName)
		fullListEndpoint := fmt.Sprintf("%s%s", listEndpoint, serviceIdList[idx])

		respStatusCode := dbtools.SubmitGetRequest(baseURL, fullListEndpoint, paramMapList)
		if respStatusCode == wantedCodes[idx] {
			globalTestCounterPass(respStatusCode)
		} else {
			globalTestCounterFail(respStatusCode)
		}

	}

}