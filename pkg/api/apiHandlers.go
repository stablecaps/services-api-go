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

//	@Summary	Get a list of all services
//	@Id			1
//	@produce	application/json
//	@Param		name	query		string	true	"Input name"
//	@Success	200		{object}	GreeterResponse
//	@Failure	404		{object}	message
//	@Router		/services [get, head]
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
	successMsg := fmt.Sprintf("Successfully created service: %s", service.ServiceName)
	log.Println(successMsg)
	return WriteJson(writer, http.StatusCreated, successMsg)
}

//	@Summary	Get service by Id
//	@Id			3
//	@produce	application/json
//	@Param		name	query		string	true	"Input name"
//	@Success	200		{object}	ServiceResponse
//	@Failure	400		{object}	message
//	@Failure	404		{object}	"404 page not found"
//	@Failure	404		{object}	message
//	@Router		/services/id/{ServiceId:[0-9]+} [get, head]
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

//	@Summary	Delete service by Id
//	@Id			4
//	@produce	application/json
//	@Param		serviceId	query		int	true	"serviceId"
//	@Failure	400		{object}	"invalid serviceidstr $serviceIdStr"
//	@Failure	500		{object}	"Server Error: $err"
//	@Failure	404		{object}	"404 page not found"
//	@Failure	404		{object}	"Could not find serviceId: $serviceId"
//	@Success	200		{object}	"deleted $serviceId"
//	@Router		/services/id/{ServiceId:[0-9]+} [delete, head]
func (server *APIServer) handleDeleteServiceById(writer http.ResponseWriter, req *http.Request) error {
	serviceId, err := getServiceId(req)
	if err != nil {
		log.Printf("Error 400: %s", err)
		return WriteJson(writer, http.StatusBadRequest, err)
	}

	fmt.Printf("Deleting service with serviceId %d", serviceId)
	numDeleted, err := server.db.DeleteServiceById(serviceId)
	if err != nil {
		log.Printf("Error 500: %s", err)
		return WriteJson(writer, http.StatusInternalServerError, err)
	}

	if numDeleted == 0 {
		notFoundMsg := fmt.Sprintf("Could not find serviceId: %s", strconv.Itoa(97))
		log.Printf("Error 404: %s", err)
		return WriteJson(writer, http.StatusNotFound, map[string]string{"Could not find": notFoundMsg})
	}

	return WriteJson(writer, http.StatusOK, map[string]int{"deleted": serviceId})
}

//	@Summary	Get service versions by Id
//	@Id			5
//	@produce	application/json
//	@Param		serviceId	query	int	true	"serviceId"
//	@Failure	400		{object}	"invalid serviceidstr $serviceIdStr"
//	@Failure	404		{object}	"404 page not found"
//	@Failure	500		{object}	"Server Error: $err"
//	@Success	200		{object}	serviceVersions
//	@Router		/services/id/{ServiceId:[0-9]+} [get, head]
func (server *APIServer) handleGetServiceVersionsById(writer http.ResponseWriter, req *http.Request) error {
	serviceId, err := getServiceId(req)
	if err != nil {
		log.Printf("Error 400: %s", err)
		return WriteJson(writer, http.StatusBadRequest, err)
	}

	fmt.Printf("checking for serviceId %d", serviceId)
	serviceVersions, err := server.db.GetServiceVersionsById(serviceId)
	if err != nil {
		log.Printf("Error 500: %s", err)
		return WriteJson(writer, http.StatusInternalServerError, err)
	}

	// TODO: implement 404 for serviceId not found

	return WriteJson(writer, http.StatusOK, serviceVersions)
}

//	@Summary	Get service by name
//	@Id			6
//	@produce	application/json
//	@Param		serviceName	query	string	true	"serviceName"
//	@Failure	500		{object}	"$err"
//	@Failure	404		{object}	"404 page not found"
//	@Success	200		{object}	service
//	@Router		/services/name/{ServiceName:[a-zA-Z0-9]+} [get, head]
func (server *APIServer) handleGetServiceByName(writer http.ResponseWriter, req *http.Request) error {
	// TODO: validate mux var (400 error)
	serviceName := mux.Vars(req)["ServiceName"]
	log.Printf("Searching for service %s by name", serviceName)

	fmt.Printf("checking for serviceName %s", serviceName)
	service, err := server.db.GetServiceByName(serviceName)
	if err != nil {
		log.Printf("Error 500: %s", err)
		return WriteJson(writer, http.StatusInternalServerError, err)
	}

	return WriteJson(writer, http.StatusOK, service)
}

//	@Summary	Get API health
//	@Id			6
//	@produce	application/json
//	@Success	200		{object}	"service is up and running"
//	@Router		/services/name/{ServiceName:[a-zA-Z0-9]+} [get, head]
func (server *APIServer) handleGetHealth(writer http.ResponseWriter, req *http.Request) error {
	return WriteJson(writer, http.StatusOK, "service is up and running")
}


////////////////////////////////////////////////////////
////////////////////////////////////////////////////////
func WriteJson(writer http.ResponseWriter, status int, value any) error {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(status)

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
