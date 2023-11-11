package dbtools

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func MakeHttpRequestWrapper(url, httpVerb string, reqBody []byte) (*http.Response, []byte) {
	// httpVerb examples: GET, POST, DELETE

	var req *http.Request
	var rerr error
	if reqBody != nil {
		req, rerr = http.NewRequest(httpVerb, url, bytes.NewBuffer(reqBody))
	} else {
		req, rerr = http.NewRequest(httpVerb, url, nil)
	}
	if rerr != nil {
		log.Printf("Error creating %s request", httpVerb)
		os.Exit(42)
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error making %s request", httpVerb)
		os.Exit(42)
	}

	defer resp.Body.Close()

	body, derr := io.ReadAll(resp.Body)
	if derr != nil {
		log.Printf("Error decoding %s response: %s", httpVerb, err)
		log.Println(resp.Body)
		os.Exit(42)
	}
	fmt.Printf("string body: %s", string(body))

	return resp, body
}