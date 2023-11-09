package main

import "fmt"

// TODO move this to tests
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
	// paramMapList := []map[string]string{}
	serviceIdList := []string{"fake", "-10", "5",}


	for idx, testName := range testNameSlice {
		fullListEndpoint := fmt.Sprintf("%s%s", listEndpoint, serviceIdList[idx])

		submitGetRequest(testName, baseURL, fullListEndpoint, map[string]string{}, idx, wantedCodes[idx])
	}

}