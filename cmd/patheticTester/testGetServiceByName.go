package main

import (
	"fmt"

	"github.com/stablecaps/services-api-go/internal/dbtools"
)

func testGetServiceByName() {
	endpoint := "/services/name/"

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
		fullEndpoint := fmt.Sprintf("%s%s", endpoint, serviceIdList[idx])

		resp, _ := dbtools.MakeHttpRequestWrapper(baseURL, fullEndpoint, "GET", paramMapList, nil)

		scoreGlobalTestsPassedandFailes(resp.StatusCode, wantedCodes[idx])

	}
}