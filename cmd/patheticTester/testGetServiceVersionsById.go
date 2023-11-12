package main

import (
	"fmt"
	"strconv"

	"github.com/stablecaps/services-api-go/internal/dbtools"
)

func testGetServiceVersionsById() {
	endpoint := "/services/id/"

	expectedServiceVersions := "v1,v2,v3,v4,v5"
	postedServiceData := dbtools.CreateExplicitService(dbtools.MakeRandomName(), dbtools.MakeRandomDescription(4), expectedServiceVersions)

	testNameSlice := []string{
		// Test serviceId 500s
		"NonExistantService",
		// Test serviceId 200
		"goodService",
	}
	wantedCodes := []int{500, 200}
	paramMapList := map[string]string{}
	serviceIdList := []string{"9999999999999", strconv.Itoa(postedServiceData.ServiceId)}


	for idx, testName := range testNameSlice {
		fmt.Println("\n~~~~~~~~~~~~~~~~~~~~")
		fmt.Printf("Running test %d: -  %s\n", idx, testName)
		fullEndpoint := fmt.Sprintf("%s%s", endpoint, serviceIdList[idx])

		resp, _ := dbtools.MakeHttpRequestWrapper(baseURL, fullEndpoint, "GET", paramMapList, nil)

		scoreGlobalTestsPassedandFailes(resp.StatusCode, wantedCodes[idx])
	}
}