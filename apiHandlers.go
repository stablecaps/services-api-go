package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ApiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

////////////////////////////////////////////////////////
////////////////////////////////////////////////////////
func makeHTTPHandleFunc(f ApiFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		if err := f(writer, req); err != nil {
			log.Printf("Error: %s", err)
			WriteJson(writer, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func (server *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/services", makeHTTPHandleFunc(server.handleGetAllServices))
	router.HandleFunc("/services/new", makeHTTPHandleFunc(server.handleCreateNewService))
	router.HandleFunc("/services/versions/{ServiceId}", makeHTTPHandleFunc(server.handleGetServiceVersionsById))
	router.HandleFunc("/services/id/{ServiceId}", makeHTTPHandleFunc(server.handleServiceId))
	router.HandleFunc("/services/name/{ServiceName}", makeHTTPHandleFunc(server.handleGetServiceByName))
	router.HandleFunc("/health", makeHTTPHandleFunc(server.handleGetHealth))

	log.Println("API server running on port: ", server.listenAddr)
	log.Fatal(http.ListenAndServe(server.listenAddr, router))
}

func (server *APIServer) handleGetAllServices(writer http.ResponseWriter, req *http.Request) error {
	serviceSlice, err := server.db.GetAllServices()
	if err != nil {
		log.Printf("Error: %s", err)
		return WriteJson(writer, http.StatusBadRequest, err)
	}
	return WriteJson(writer, http.StatusOK, serviceSlice)
}

func (server *APIServer) handleCreateNewService(writer http.ResponseWriter, req *http.Request) error {
	log.Println("req.Method", req.Method)
	if req.Method != "POST" {
		return WriteJson(writer, http.StatusPreconditionFailed, "create service should use post")
	}

	log.Println("Creating new service")
	createServReq := new(CreateServiceRequest)
	log.Printf("createServReq: %s", createServReq)

	err := decodeJSONBody(writer, req, &createServReq)
    if err != nil {
        var mr *malformedRequest
        if errors.As(err, &mr) {
            http.Error(writer, mr.msg, mr.status)
        } else {
            log.Print(err.Error())
            http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
        }
        return err
    }

	service := NewService(createServReq.ServiceName, createServReq.ServiceDescription)
	if err := server.db.CreateNewService(service); err != nil {
		log.Printf("Error creating service: %s", err)
		return WriteJson(writer, http.StatusBadRequest, err)
	}

	return WriteJson(writer, http.StatusOK, service)
}

func (server *APIServer) handleServiceId(writer http.ResponseWriter, req *http.Request) error {
	if req.Method == "GET" {
		return server.handleGetServiceById(writer, req)
	}

	if req.Method == "DELETE" {
		return server.handleDeleteServiceById(writer, req)
	}

	return fmt.Errorf("unsupported method: %s", req.Method)
}

func (server *APIServer) handleGetServiceById(writer http.ResponseWriter, req *http.Request) error {
	serviceId, err := getServiceId(req)
	if err != nil {
		return WriteJson(writer, http.StatusBadRequest, err)
	}

	fmt.Printf("checking for serviceId %d", serviceId)
	service, err := server.db.GetServiceVersionsById(serviceId)
	if err != nil {
		log.Printf("Error: %s", err)
		return WriteJson(writer, http.StatusNotFound, err)
	}

	return WriteJson(writer, http.StatusOK, service)
}

func (server *APIServer) handleDeleteServiceById(writer http.ResponseWriter, req *http.Request) error {
	serviceId, err := getServiceId(req)
	if err != nil {
		return WriteJson(writer, http.StatusBadRequest, err)
	}

	fmt.Printf("Deleting service with serviceId %d", serviceId)
	if err := server.db.DeleteServiceById(serviceId); err != nil {
		log.Printf("Error: %s", err)
		return WriteJson(writer, http.StatusNotFound, err)
	}

	return WriteJson(writer, http.StatusOK, map[string]int{"deleted": serviceId})
}

func (server *APIServer) handleGetServiceVersionsById(writer http.ResponseWriter, req *http.Request) error {
	serviceId, err := getServiceId(req)
	if err != nil {
		return err
	}

	fmt.Printf("checking for serviceId %d", serviceId)
	serviceVersions, err := server.db.GetServiceVersionsById(serviceId)
	if err != nil {
		log.Printf("Error: %s", err)
		return WriteJson(writer, http.StatusNotFound, err)
	}

	return WriteJson(writer, http.StatusOK, serviceVersions)
}

func (server *APIServer) handleGetServiceByName(writer http.ResponseWriter, req *http.Request) error {
	serviceName := mux.Vars(req)["ServiceName"]
	log.Printf("Searching for service %s  by name", serviceName)

	fmt.Printf("checking for serviceName %s", serviceName)
	service, err := server.db.GetServiceByName(serviceName)
	if err != nil {
		log.Printf("Error: %s", err)
		return WriteJson(writer, http.StatusNotFound, err)
	}

	return WriteJson(writer, http.StatusOK, service)
}

func (server *APIServer) handleGetHealth(writer http.ResponseWriter, req *http.Request) error {
	return WriteJson(writer, http.StatusOK, "service is up and running")
}
////////////////////////////////////////////////////////
////////////////////////////////////////////////////////
func WriteJson(writer http.ResponseWriter, status int, value any) error {
	writer.WriteHeader(status)
	writer.Header().Add("Content-Type", "application/json")

	jsonResponse := json.NewEncoder(writer).Encode(value)
	return jsonResponse
}

func getServiceId(req *http.Request) (int, error) {
	serviceIdStr := mux.Vars(req)["ServiceId"]
	serviceId, err := strconv.Atoi(serviceIdStr)
	if err != nil {
		return serviceId, fmt.Errorf("invalid serviceidstr %s", serviceIdStr)
	}

	return serviceId, nil
}
