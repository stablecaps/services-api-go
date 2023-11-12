package main

import (
	"fmt"

	"github.com/stablecaps/services-api-go/internal/dbtools"
)


func testListServices() {
	// list services pathetic tests
	endpoint := "/services"


	testNameSlice := []string{
		// Test 400s
		"badLimit", "outOfRangeLimit", "badOffset", "outOfRangeOffset", "badOrderColName", "badOrderDir",
		// Test orderDir 200
		"orderDirasc", "orderDirdesc",
		// Test orderColName 200
		"orderColNameServiceId", "orderColNameServiceName", "orderColNameServiceDescription", "orderColNameServiceVersions", "orderColNameCreatedAt",
	}
	wantedCodes := []int{400, 400, 400, 400, 400, 400, 200, 200, 200, 200, 200, 200, 200}
	paramMapList := []map[string]string{
		// Test 400s
		{
			"limit": "fake",
			"offset": "0",
			"orderColName": "serviceName",
			"orderDir": "asc",
		},
		{
			"limit": "-10",
			"offset": "0",
			"orderColName": "serviceName",
			"orderDir": "asc",
		},
		{
			"limit": "4",
			"offset": "fake",
			"orderColName": "serviceName",
			"orderDir": "asc",
		},
		{
			"limit": "4",
			"offset": "-10",
			"orderColName": "serviceName",
			"orderDir": "asc",
		},
		{
			"limit": "4",
			"offset": "0",
			"orderColName": "fake",
		},
		{
			"limit": "4",
			"offset": "0",
			"orderColName": "serviceName",
			"orderDir": "fake",
		},
		// Test orderColName 200
		{
			"limit": "4",
			"offset": "0",
			"orderColName": "serviceId",
			"orderDir": "asc",
		},
		{
			"limit": "4",
			"offset": "0",
			"orderColName": "serviceName",
			"orderDir": "asc",
		},
		{
			"limit": "4",
			"offset": "0",
			"orderColName": "serviceDescription",
			"orderDir": "asc",
		},
		{
			"limit": "4",
			"offset": "0",
			"orderColName": "serviceVersions",
			"orderDir": "asc",
		},
		{
			"limit": "4",
			"offset": "0",
			"orderColName": "createdAt",
			"orderDir": "asc",
		},
		// Test orderDir 200
		{
			"limit": "4",
			"offset": "0",
			"orderColName": "serviceName",
			"orderDir": "asc",
		},

		{
			"limit": "4",
			"offset": "0",
			"orderColName": "serviceName",
			"orderDir": "asc",
		},
	}


	for idx, testName := range testNameSlice {
		fmt.Println("\n~~~~~~~~~~~~~~~~~~~~")
		fmt.Printf("Running test %d: -  %s\n", idx, testName)

		resp, _ := dbtools.MakeHttpRequestWrapper(baseURL, endpoint, "GET", paramMapList[idx], nil)

		scoreGlobalTestsPassedandFailes(resp.StatusCode, wantedCodes[idx])
	}
}
