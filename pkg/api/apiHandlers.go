package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/stablecaps/services-api-go/pkg/models"
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

func (server *APIServer) handleGetAllServices(writer http.ResponseWriter, req *http.Request) error {

	strLimit := req.URL.Query().Get("limit")
	log.Printf("strLimit is: %s", strLimit)
	limit := 10
    // with a value as -1 for gorms Limit method, we'll get a request without limit as default
    if strLimit != "" {
		var err error
        limit, err = strconv.Atoi(strLimit)
        if err != nil || limit < -1 {
			return WriteJson(writer, http.StatusBadRequest, "limit query parameter is not a valid number")
        }

    }
	log.Printf("limit is: %d", limit)


    strOffset := req.URL.Query().Get("offset")
	log.Printf("strOffset is: %s", strOffset)
	offset := 0
    if strOffset != "" {
		var err error
        offset, err = strconv.Atoi(strOffset)
        if err != nil || offset < -1 {
			WriteJson(writer, http.StatusBadRequest, "offset query parameter is not a valid number")
        }
    }
	log.Printf("offset is: %d", offset)

	// pageSize := 10
	serviceSlice, err := server.db.GetAllServices(limit, offset)
	if err != nil {
		log.Printf("Error: %s", err)
		return WriteJson(writer, http.StatusBadRequest, err)
	}
	return WriteJson(writer, http.StatusOK, serviceSlice)
}

func (server *APIServer) handleCreateNewService(writer http.ResponseWriter, req *http.Request) error {

	log.Println("Creating new service")
	createServReq := new(models.CreateServiceRequest)
	log.Printf("createServReq: %s", createServReq)


	err := decodeJSONBody(writer, req, &createServReq)
	jsonInputCheck(err,  writer)

	service := models.NewService(createServReq.ServiceName, createServReq.ServiceDescription)
	if err := server.db.CreateNewService(service); err != nil {
		log.Printf("Error creating service: %s", err)
		return WriteJson(writer, http.StatusBadRequest, err)
	}
	log.Printf("Successfully created service: %s", err)
	return WriteJson(writer, http.StatusCreated, service)
}

func (server *APIServer) handleGetServiceById(writer http.ResponseWriter, req *http.Request) error {
	serviceId, err := getServiceId(req)
	if err != nil {
		return WriteJson(writer, http.StatusBadRequest, err)
	}


	fmt.Printf("checking for serviceId %d", serviceId)
	service, err := server.db.GetServiceById(serviceId)
	if err != nil {
		log.Printf("Error: %s", err)
		return WriteJson(writer, http.StatusNotFound, "Error: serviceId not found")
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
