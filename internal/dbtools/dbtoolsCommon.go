package dbtools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"slices"
	"strings"

	"github.com/stablecaps/services-api-go/pkg/models"
)

func readHttpBody(resp *http.Response, httpVerb string) []byte {
	body, derr := io.ReadAll(resp.Body)
	if derr != nil {
		log.Printf("Error decoding %s response: %s", httpVerb, derr)
		log.Println(resp.Body)
		os.Exit(42)
	}
	fmt.Printf("string body: %s", string(body))

	return body
}

func MakeHttpRequestWrapper(baseURL, endpoint, httpVerb string, paramMap map[string]string, reqBody []byte) (*http.Response, []byte) {
	allowedHttpVerbs := []string{"GET", "POST", "DELETE"}
	if !slices.Contains(allowedHttpVerbs, httpVerb) {
		log.Printf("HTTP verb %s not allowed!", httpVerb)
		log.Printf("Must be one of %s", strings.Join(allowedHttpVerbs[:], ", "))
		os.Exit(42)
	}

	myUrl, _ := url.ParseRequestURI(baseURL)
	myUrl.Path = endpoint

	// add params
	if len(paramMap) > 0 {
		params := url.Values{}
		for key, val := range paramMap {
			fmt.Println(key, val)
			params.Add(key, val)
		}
		myUrl.RawQuery = params.Encode()
	}
	myUrlStr := fmt.Sprintf("%v", myUrl)

	var req *http.Request
	var rerr error
	if reqBody != nil {
		req, rerr = http.NewRequest(httpVerb, myUrlStr, bytes.NewBuffer(reqBody))
	} else {
		req, rerr = http.NewRequest(httpVerb, myUrlStr, nil)
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

	body := readHttpBody(resp, httpVerb)
	return resp, body
}

func MakeExplicitServiceJson(name, description, versions string) []byte {
	log.Printf("name is %s", name)
	log.Printf("description is %s", description)

	body := []byte(fmt.Sprintf(`{
		"serviceName": "%s",
		"serviceDescription": "%s",
		"serviceVersions": "%s"
	}`, name, description, versions))

	return body
}

func getServiceFromPostRequest(url string, reqBody []byte) models.Service {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		log.Println("Error creating post request")
		os.Exit(42)
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error making post request")
		os.Exit(42)
	}

	defer resp.Body.Close()

	body, derr := io.ReadAll(resp.Body)
	if derr != nil {
		log.Printf("Error decoding post respponse: %s", err)
		log.Println(resp.Body)
		os.Exit(42)
	}
	fmt.Printf("string body: %s", string(body))

	if resp.StatusCode != http.StatusCreated {
		log.Printf("Error unexpected status: %d", resp.StatusCode)
		os.Exit(42)
	}

	// Unmarshal JSON to Go struct
	var result models.Service
	// Parse []byte to go struct pointer
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
		os.Exit(42)
	}
	fmt.Println(PrettyPrint(result))

	fmt.Println("ServiceId:", result.ServiceId)
	fmt.Println("Name:", result.ServiceName)
	fmt.Println("Descripton:", result.ServiceDescription)

	return result
}

func CreateExplicitService(serviceName, serviceDescriptions, serviceVersions string) models.Service {
	fmt.Println("\n~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("Posting sample service to test API delete endpoint")

	// baseURL = "http://localhost:8969/"
	body := MakeExplicitServiceJson(serviceName, serviceDescriptions, serviceVersions)
	createNewServiceUrl := "http://localhost:8969/services/new"

	postedServiceData := getServiceFromPostRequest(createNewServiceUrl, body)

	return postedServiceData
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
