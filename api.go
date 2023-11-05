package main

import (
	"encoding/json"
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

type APIServer struct {
	listenAddr string
	db 	Dbase
}

func NewAPIServer(listenAddr string, db Dbase) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		db: db,
	}
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
	router.HandleFunc("/services/new", makeHTTPHandleFunc(server.handleCreateService))
	router.HandleFunc("/services/versions/{ServiceId}", makeHTTPHandleFunc(server.handleGetServiceVersionsById))
	router.HandleFunc("/services/id/{ServiceId}", makeHTTPHandleFunc(server.handleServiceId))
	router.HandleFunc("/services/name/{ServiceName}", makeHTTPHandleFunc(server.handleGetServiceByName))


	log.Println("API server running on port: ", server.listenAddr)
	http.ListenAndServe(server.listenAddr, router)
}

func (server *APIServer) handleGetAllServices(writer http.ResponseWriter, req *http.Request) error {
	serviceSlice, err := server.db.GetAllServices()
	if err != nil {
		log.Printf("Error: %s", err)
		return err
	}
	return WriteJson(writer, http.StatusOK, serviceSlice)

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
		return err
	}

	fmt.Printf("checking for serviceId %d", serviceId)
	service, err := server.db.GetServiceVersionsById(serviceId)
	if err != nil {
		log.Printf("Error: %s", err)
		return err
	}

	return WriteJson(writer, http.StatusOK, service)
}

func (server *APIServer) handleDeleteServiceById(writer http.ResponseWriter, req *http.Request) error {
	serviceId, err := getServiceId(req)
	if err != nil {
		return err
	}

	fmt.Printf("Deleting service with serviceId %d", serviceId)
	if err := server.db.DeleteServiceById(serviceId); err != nil {
		log.Printf("Error: %s", err)
		return err
	}

	return WriteJson(writer, http.StatusOK, map[string]int{"deleted": serviceId})
}

func (server *APIServer) handleCreateService(writer http.ResponseWriter, req *http.Request) error {

	log.Println("Creating new service")
	createServReq := new(CreateServiceRequest)

	if err := json.NewDecoder(req.Body).Decode(createServReq); err != nil {
		log.Printf("Error decoding json")
		return err
	}

	service := NewService(createServReq.ServiceName, createServReq.ServiceDescription)
	if err := server.db.CreateService(service); err != nil {
		log.Printf("Error creating service: %s", err)
		return err
	}

	return WriteJson(writer, http.StatusOK, service)
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
		return err
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
		return err
	}

	return WriteJson(writer, http.StatusOK, service)
}

////////////////////////////////////////////////////////
////////////////////////////////////////////////////////
func WriteJson(writer http.ResponseWriter, status int, value any) error {
	writer.WriteHeader(status)
	writer.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(writer).Encode(value)
}


func getServiceId(req *http.Request) (int, error) {
	serviceIdStr := mux.Vars(req)["ServiceId"]
	serviceId, err := strconv.Atoi(serviceIdStr)
	if err != nil {
		return serviceId, fmt.Errorf("invalid serviceidstr %s", serviceIdStr)
	}

	return serviceId, nil
}