package main

import (
	"fmt"
	"strconv"

	"github.com/fatih/color"
)

// Declaring Global variables
var baseURL string = "http://localhost:8969/"
var testsPassed int = 0
var testsFailed int = 0

func scoreGlobalTestsPassedandFailes(testCode, wantedCode int) {
	if testCode == wantedCode {
		color.Green("Test passed: %d", testCode)
		testsPassed ++
	} else {
		color.Red("Error!! Wanted %d, but got unexpected status: %d", wantedCode, testCode)
		testsFailed ++
	}
}

func printTestSeperator(testCategory string) {
	println("\n##################################")
	color.Yellow(fmt.Sprintf("   >> %s", testCategory))
}

func main() {


	// Run tests
	printTestSeperator("testGetHealthendpoint")
	testGetHealthendpoint()

	printTestSeperator("testListServices")
	testListServices()

	printTestSeperator("testGetServiceById")
	testGetServiceById()

	printTestSeperator("testGetServiceByName")
	testGetServiceByName()

	printTestSeperator("testGetServiceVersionsById")
	testGetServiceVersionsById()

	printTestSeperator("testDeleteServiceById")
	testDeleteServiceById()

	printTestSeperator("testDeleteServiceByIdError")
	testDeleteServiceByIdError()

	// TODO: test new service endpoint. kind of tested already with populate db
	// printTestSeperator("testCreateNewService")
	// testCreateNewService()

	println("\n\n")
	color.Red("Tests failed: %s", strconv.Itoa(testsFailed))
	color.Green("Tests passed: %s", strconv.Itoa(testsPassed))

}