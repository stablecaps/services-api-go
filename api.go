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
	writer.Header().Set("Content0-Type", "application/json")
	return json.NewEncoder(writer).Encode(value)
}

type ApiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

func makeHTTPHandleFunc(f ApiFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		if error := f(writer, req); error != nil {
			WriteJson(writer, http.StatusBadRequest, ApiError{Error: error.Error()})
		}
	}
}

type APIServer struct {
	listenerAddr string
}

func NewAPIServer(listenerAddr string) *APIServer {
	return &APIServer{
		listenerAddr: listenerAddr,
	}
}

func (server *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/service", makeHTTPHandleFunc(server.handleService))

	log.Println("API server running on port: ", server.listenerAddr)
	http.ListenAndServe(server.listenerAddr, router)
}

func (server *APIServer) handleService(writer http.ResponseWriter, req *http.Request) error {
	if req.Method == "GET" {
		return server.handleGetService(writer, req)
	} else if req.Method == "POST" {
		return server.handleGetService(writer, req)
	} else if req.Method == "DELETE" {
		return server.handleGetService(writer, req)
	}

	return fmt.Errorf("unsupported method: <%s>", req.Method)
}

func (server *APIServer) handleListAllServices(writer http.ResponseWriter, req *http.Request) error {
	return nil
}

func (server *APIServer) handleGetService(writer http.ResponseWriter, req *http.Request) error {
	return nil
}

func (server *APIServer) handleGetServiceVersions(writer http.ResponseWriter, req *http.Request) error {
	return nil
}

func (server *APIServer) handleCreateService(writer http.ResponseWriter, req *http.Request) error {
	return nil
}

func (server *APIServer) handleDeleteService(writer http.ResponseWriter, req *http.Request) error {
	return nil
}