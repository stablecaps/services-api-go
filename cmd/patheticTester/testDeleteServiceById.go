package main

import (
	"fmt"

	"github.com/stablecaps/services-api-go/internal/dbtools"
)


func testDeleteRequestById() {
	fmt.Println("\n~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("Posting sample service to test API delete endpoint")

	body := dbtools.MakeExplicitServiceJson("testDeleter2", "a service we will test delete", "v1,v2,v3")
	createNewServiceUrl := "http://localhost:8969/services/new"

	postedServiceData, _ := dbtools.SubmitPostRequest(createNewServiceUrl, body)

	fmt.Println("\n~~~~~~~~~~~~~~~~~~~~")
	fmt.Printf("Now Deleting ServiceId %d via API delete endpoint\n", postedServiceData.ServiceId)

	deleteNewServiceUrl := fmt.Sprintf("http://localhost:8969/services/id/%d", postedServiceData.ServiceId)
	deletedServiceData, respStatusCode := dbtools.SubmitDeleteRequest(deleteNewServiceUrl, nil)
	fmt.Printf("deletedServiceData: %v\n", deletedServiceData)

	if respStatusCode == 200 {
		globalTestCounterPass(respStatusCode)
	} else {
		globalTestCounterFail(respStatusCode)
	}
}