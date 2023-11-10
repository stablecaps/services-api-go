package main

import (
	"fmt"
	"log"
	"net/url"
	"strconv"
)

type test_struct struct {
    Test string
}


func main() {
	baseURL := "http://localhost:8969/"
	resource := "/services"

	numServices := 500

	page := 1
	limit := 5
	offsetInc := (limit * page) - limit //(page - 1) * limit
	for offset := 0; offset <= numServices; offset = offset + offsetInc {
		fmt.Printf("offset & page: %d & %d\n", offset, page)
		params := url.Values{}
		params.Add("limit", strconv.Itoa(limit))
		params.Add("offset", strconv.Itoa(offset))

		u, _ := url.ParseRequestURI(baseURL)
		u.Path = resource
		u.RawQuery = params.Encode()
		urlStr := fmt.Sprintf("%v", u)

		log.Printf("Listing services by using limit & offset to %s", urlStr)
		page  = page + 1

	}

}
