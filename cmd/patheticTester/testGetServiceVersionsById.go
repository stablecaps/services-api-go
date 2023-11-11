package main

import (
	"fmt"
	"strconv"

	"github.com/stablecaps/services-api-go/internal/dbtools"
)

func testGetServiceVersionsById() {
	baseURL := "http://localhost:8969/"
	listEndpoint := "/services/id/"

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
		fullListEndpoint := fmt.Sprintf("%s%s", listEndpoint, serviceIdList[idx])

		_, respStatusCode := dbtools.SubmitGetRequest(baseURL, fullListEndpoint, paramMapList)

		if respStatusCode == wantedCodes[idx] {
			globalTestCounterPass(respStatusCode)
		} else {
			globalTestCounterFail(respStatusCode)
		}

	}
}