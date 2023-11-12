package main

import (
	"fmt"

	"github.com/stablecaps/services-api-go/internal/dbtools"
)



func testDeleteServiceById() {
	postedServiceData := dbtools.CreateExplicitService(dbtools.MakeRandomName(), dbtools.MakeRandomDescription(4), "v1,v2,v3")

	fmt.Println("\n~~~~~~~~~~~~~~~~~~~~")
	fmt.Printf("Now Deleting ServiceId %d via API delete endpoint\n", postedServiceData.ServiceId)

	fullEndpoint := fmt.Sprintf("/services/id/%d", postedServiceData.ServiceId)
	paramMapList := map[string]string{}
	resp, _ := dbtools.MakeHttpRequestWrapper(baseURL, fullEndpoint, "DELETE", paramMapList, nil)

scoreGlobalTestsPassedandFailes(resp.StatusCode, 200)
}

func testDeleteServiceByIdError() {
	endpoint := "/services/id/"


	testNameSlice := []string{
		// Test serviceId 400s
		"badServiceId", "outOfRangeServiceId",
		// Test serviceId 200
		"goodServiceIdOutOfRange",
	}
	wantedCodes := []int{404, 404, 500}
	paramMapList := map[string]string{}
	serviceIdList := []string{"fake", "-10", "99999999"}


	for idx, testName := range testNameSlice {
		fmt.Println("\n~~~~~~~~~~~~~~~~~~~~")
		fmt.Printf("Running test %d: -  %s\n", idx, testName)
		fullEndpoint := fmt.Sprintf("%s%s", endpoint, serviceIdList[idx])

		resp, _ := dbtools.MakeHttpRequestWrapper(baseURL, fullEndpoint, "GET", paramMapList, nil)

		fmt.Printf("resp.StatusCode: %d\n", resp.StatusCode)
		scoreGlobalTestsPassedandFailes(resp.StatusCode, wantedCodes[idx])

	}
}