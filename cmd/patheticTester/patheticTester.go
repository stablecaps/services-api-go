package main

import (
	"strconv"

	"github.com/fatih/color"
)

// Declaring Global variables
var testsPassed int = 0
var testsFailed int = 0


func globalTestCounterPass(statusCode int)  {
	color.Green("Test passed: %d", statusCode)
	testsPassed ++
}


func globalTestCounterFail(statusCode int) {
	color.Red("Error!! unexpected status: %d", statusCode)
	testsFailed ++
}


func main() {


	// Run tests
	// testListServices()

	// println("\n#######################")
	// println("#######################")
	// println("#######################")

	testGetServiceById()

	// println("\n#######################")
	// println("#######################")
	// println("#######################")
	testDeleteRequestById()

	println("\n\n")
	color.Red("Tests failed: %s", strconv.Itoa(testsFailed))
	color.Green("Tests passed: %s", strconv.Itoa(testsPassed))

}