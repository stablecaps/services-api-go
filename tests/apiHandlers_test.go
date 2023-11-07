package main

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"
)

// func (server *APIServer) handleCreateNewService(writer http.ResponseWriter, req *http.Request) error {
// 	if req.Method != "POST" {
// 		return WriteJson(writer, http.StatusPreconditionFailed, fmt.Errorf("create service should use post"))
// 	}

// 	log.Println("Creating new service")
// 	createServReq := new(CreateServiceRequest)
// 	log.Printf("createServReq: %s", createServReq)

// 	err := decodeJSONBody(writer, req, &createServReq)
//     if err != nil {
//         var mr *malformedRequest
//         if errors.As(err, &mr) {
//             http.Error(writer, mr.msg, mr.status)
//         } else {
//             log.Print(err.Error())
//             http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
//         }
//         return err
//     }
// 	// if err := json.NewDecoder(req.Body).Decode(createServReq); err != nil {
// 	// 	log.Printf("Error decoding json: %s", err)
// 	// 	return WriteJson(writer, http.StatusBadRequest, err)
// 	// }

// 	service := NewService(createServReq.ServiceName, createServReq.ServiceDescription)
// 	if err := server.db.CreateNewService(service); err != nil {
// 		log.Printf("Error creating service: %s", err)
// 		return WriteJson(writer, http.StatusBadRequest, err)
// 	}

// 	return WriteJson(writer, http.StatusOK, service)
// }

// func TestHandleCreateNewService(t *testing.T) {
// 	go main()

// 	response, err := http.Get("http://localhost:8969/services/new")
// 	if err != nil {
// 		t.Errorf("expected no errors, but got %v", err)
// 	}

// 	if response.StatusCode != http.StatusOK {
// 		t.Errorf("expected 200 statuscode, but got %v", response.StatusCode)
// 	}

// 	responseBody := ""
// 	json.NewDecoder(response.Body).Decode(&responseBody)
// 	response.Body.Close()

// 	if responseBody != "service is up and running" {
// 		t.Errorf(`expected message to be "service is up and running", but got %v`, responseBody)
// 	}

// 	os.Interrupt.Signal()
// }

func TestHealthCheck(t *testing.T) {
	go main()

	response, err := http.Get("http://localhost:8969/health")
	if err != nil {
		t.Errorf("expected no errors, but got %v", err)
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("expected 200 statuscode, but got %v", response.StatusCode)
	}

	responseBody := ""
	json.NewDecoder(response.Body).Decode(&responseBody)
	response.Body.Close()

	if responseBody != "service is up and running" {
		t.Errorf(`expected message to be "service is up and running", but got %v`, responseBody)
	}

	os.Interrupt.Signal()
}
