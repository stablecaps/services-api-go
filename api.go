package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func WriteJson(writer http.ResponseWriter, status int, value any) error {
	writer.WriteHeader(status)
	writer.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(writer).Encode(value)
}

type ApiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

func makeHTTPHandleFunc(f ApiFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		if err := f(writer, req); err != nil {
			WriteJson(writer, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
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

func (server *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/services", makeHTTPHandleFunc(server.handleService))
	router.HandleFunc("/services/{ServiceId}", makeHTTPHandleFunc(server.handleGetServiceById))

	log.Println("API server running on port: ", server.listenAddr)
	http.ListenAndServe(server.listenAddr, router)
}

func (server *APIServer) handleService(writer http.ResponseWriter, req *http.Request) error {
	if req.Method == "GET" {
		return server.handleGetAllServices(writer, req)
	}

	if req.Method == "POST" {
		return server.handleCreateService(writer, req)
	}

	if req.Method == "DELETE" {
		return server.handleDeleteService(writer, req)
	}

	return fmt.Errorf("unsupported method: <%s>", req.Method)
}

func (server *APIServer) handleGetAllServices(writer http.ResponseWriter, req *http.Request) error {
	serviceSlice, err := server.db.GetAllServices()
	if err != nil {
		return err
	}
	return WriteJson(writer, http.StatusOK, serviceSlice)

}

func (server *APIServer) handleGetServiceById(writer http.ResponseWriter, req *http.Request) error {

	service := NewService("TestService", "A test service to play with")

	return WriteJson(writer, http.StatusOK, service)
}

func (server *APIServer) handleGetServiceVersions(writer http.ResponseWriter, req *http.Request) error {
	return nil
}

func (server *APIServer) handleCreateService(writer http.ResponseWriter, req *http.Request) error {

	log.Println("Creating new service")
	createServReq := new(CreateServiceRequest)
	if err := json.NewDecoder(req.Body).Decode(createServReq); err != nil {
		return err
	}

	service := NewService(createServReq.ServiceName, createServReq.ServiceDescription)
	if err := server.db.CreateService(service); err != nil {
		return err
	}

	return WriteJson(writer, http.StatusOK, service)
}

func (server *APIServer) handleDeleteService(writer http.ResponseWriter, req *http.Request) error {
	return nil
}