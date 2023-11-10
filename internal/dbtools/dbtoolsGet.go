package dbtools

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func SubmitGetRequest(baseURL, endpoint string, paramMap map[string]string) int {

	u, _ := url.ParseRequestURI(baseURL)
	u.Path = endpoint

	// add params
	if len(paramMap) > 0 {
	params := url.Values{}
		for key, val := range paramMap {
			fmt.Println(key, val)
			params.Add(key, val)
		}
		u.RawQuery = params.Encode()
	}

	urlStr := fmt.Sprintf("%v", u)
	log.Printf("Calling services with the following url: %s", urlStr)

	resp, err := http.Get(urlStr)

	if err != nil {
		log.Printf("Request Failed: %s", err)
		return 0
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return 0
	}

	// Log the request body
	bodyString := string(body)
	log.Print(bodyString)

	return resp.StatusCode
}