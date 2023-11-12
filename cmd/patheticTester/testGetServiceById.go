package main

import (
	"fmt"

	"github.com/stablecaps/services-api-go/internal/dbtools"
)

func testGetServiceById() {
	endpoint := "/services/id/"


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
		fmt.Printf("Running test %d: -  %s\n", idx, testName)
		fullEndpoint := fmt.Sprintf("%s%s", endpoint, serviceIdList[idx])

		resp, _ := dbtools.MakeHttpRequestWrapper(baseURL, fullEndpoint, "GET", paramMapList, nil)
		scoreGlobalTestsPassedandFailes(resp.StatusCode, wantedCodes[idx])

	}
}