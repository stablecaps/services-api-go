package main

import (
	"fmt"

	"github.com/stablecaps/services-api-go/internal/dbtools"
)



func testDeleteserviceById() {
	postedServiceData := dbtools.CreateExplicitService(dbtools.MakeRandomName(), dbtools.MakeRandomDescription(4), "v1,v2,v3")


	fmt.Println("\n~~~~~~~~~~~~~~~~~~~~")
	fmt.Printf("Now Deleting ServiceId %d via API delete endpoint\n", postedServiceData.ServiceId)

	deleteNewServiceUrl := fmt.Sprintf("http://localhost:8969/services/id/%d", postedServiceData.ServiceId)
	respStatusCode := dbtools.SubmitDeleteRequest(deleteNewServiceUrl, nil)

	if respStatusCode == 200 {
		globalTestCounterPass(respStatusCode)
	} else {
		globalTestCounterFail(respStatusCode)
	}
}

func testDeleteserviceByIdError() {
	baseURL := "http://localhost:8969/"
	listEndpoint := "/services/id/"


	testNameSlice := []string{
		// Test serviceId 400s
		"badServiceId", "outOfRangeServiceId",
		// Test serviceId 200
		"goodServiceIdOutOfRange",
	}
	wantedCodes := []int{404, 404, 500}
	serviceIdList := []string{"fake", "-10", "99999999"}


	for idx, testName := range testNameSlice {
		fmt.Println("\n~~~~~~~~~~~~~~~~~~~~")
		fmt.Printf("Running test %d: -  %s\n", idx, testName)
		fullDeleteEndpoint := fmt.Sprintf("%s%s%s", baseURL, listEndpoint, serviceIdList[idx])

		respStatusCode  := dbtools.SubmitDeleteRequest(fullDeleteEndpoint, nil)
		fmt.Printf("respStatusCode: %d\n", respStatusCode)
		if respStatusCode == wantedCodes[idx] {
			globalTestCounterPass(respStatusCode)
		} else {
			globalTestCounterFail(respStatusCode)
		}

	}
}