package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"slices"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/stablecaps/services-api-go/pkg/models"
)

type ApiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
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

func validateLimit(strLimit string) int {
	// strLimit := req.URL.Query().Get("limit")
	log.Printf("strLimit is: %s", strLimit)
	limit := -1
    if strLimit != "" {
		var err error
        limit, err = strconv.Atoi(strLimit)
        if err != nil || limit < -1 {
			fmt.Printf("limit query param invalid: %s. Aborting..", err)
			// return WriteJson(writer, http.StatusBadRequest, fmt.Sprintf("limit query param invalid: %s. Must be an int", strLimit))
			return -1
        }

    }
	log.Printf("limit is: %d", limit)
	return limit
}

func validateOffset(strOffset string) int {
	log.Printf("strOffset is: %s", strOffset)
	offset := -1
    if strOffset != "" {
		var err error
        offset, err = strconv.Atoi(strOffset)
        if err != nil || offset < -1 {
			fmt.Printf("offset query param invalid: %s. Aborting..", err)
			return -1
        }
    }
	log.Printf("offset is: %d", offset)
	return offset
}

func validateColName(strOrderColName string) (string, error) {
	log.Printf("strOrderColName is: %s", strOrderColName)
	orderColName := "serviceId"
    if strOrderColName != "" {
		var existingColumnNames = []string{"serviceid","servicename","servicedescription","serviceversions","createdat"}
		if slices.Contains(existingColumnNames, strOrderColName) {
			orderColName = strOrderColName
		} else {
			strExistingColumnNames := strings.Join(existingColumnNames[:], ", ")
			colNameErr := fmt.Errorf("orderColName query param invalid: %s. Must be one of %s", strOrderColName, strExistingColumnNames)
			return "", colNameErr
		}
    }
	return orderColName, nil
}

func validateOrderDir(strOrderDir string) (string, error) {
	log.Printf("strOrderDir is: %s", strOrderDir)
	orderDir := "asc"
    if strOrderDir != "" {
		existingDirectionNames := []string{"asc","desc"}
		if slices.Contains(existingDirectionNames, strOrderDir) {
			orderDir = strOrderDir
		} else {
			strExistingDirectionNames := strings.Join(existingDirectionNames[:], ", ")
			directionNamesErr := fmt.Errorf("orderDir query param invalid: %s. Must be one of %s", strOrderDir, strExistingDirectionNames)
			return "", directionNamesErr
		}
    }
	log.Printf("orderDir is: %s", orderDir)
	return orderDir, nil
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
//	@Failure	500		message	"Server Error: $err"
//	@Failure	400		message	"limit query param invalid: $err"
//	@Failure	400		message	"offset query param invalid: $err"
//	@Success	200		message	model.Service
//	@Router		/services [get, head]
func (server *APIServer) handleGetAllServices(writer http.ResponseWriter, req *http.Request) error {
	//https://stackoverflow.com/questions/57776448/stop-processing-of-http-request-in-go

	strLimit := req.URL.Query().Get("limit")
	limit := validateLimit(strLimit)
	if limit == -1 {
		return WriteJson(writer, http.StatusBadRequest, fmt.Sprintf("Error 400: limit query param invalid: %s. Must be an int", strLimit))
	}

    strOffset := req.URL.Query().Get("offset")
	offset := validateOffset(strOffset)
	if limit == -1 {
		return WriteJson(writer, http.StatusBadRequest, fmt.Sprintf("Error 400: offset query param invalid: %s. Must be an int.", strOffset))
	}

	strOrderColName := req.URL.Query().Get("orderColName")
	orderColName, err := validateColName(strings.ToLower(strOrderColName))
	if err != nil {
		return WriteJson(writer, http.StatusBadRequest, fmt.Sprintf("Error 400: %s", err))
	}

	strOrderDir := req.URL.Query().Get("orderDir")
	orderDir, err := validateOrderDir(strings.ToLower(strOrderDir))
	if err != nil {
		return WriteJson(writer, http.StatusBadRequest, fmt.Sprintf("Error 400: %s", err))
	}

	//
	serviceSlice, err := server.db.GetAllServices(orderColName, orderDir, limit, offset)
	if err != nil {
		err500 := fmt.Sprintf("Server Error: %s", err)
		log.Println(err500)
		return WriteJson(writer, http.StatusInternalServerError, err500)
	}
	return WriteJson(writer, http.StatusOK, serviceSlice)
}

//	@Summary	Create new service
//	@Id			2
//	@produce	application/json
//	@Param		mode.CreateServiceRequest	body message	true
//	@Failure	500		message	"Server Error: $err"
//	@Failure	404		message	"404 page not found"
//	@Success	201		message	model.Service
//	@Router		/services/new [post, head]
func (server *APIServer) handleCreateNewService(writer http.ResponseWriter, req *http.Request) error {

	log.Println("Creating new service")
	createServReq := new(models.CreateServiceRequest)
	log.Printf("createServReq: %s", createServReq)

	// TODO: add this to swagger notation
	err := decodeJSONBody(writer, req, &createServReq)
	jsonInputCheck(err,  writer)

	service := models.NewService(createServReq.ServiceName, createServReq.ServiceDescription)

	if err := server.db.CreateNewService(service); err != nil {
		err500 := fmt.Sprintf("Server Error: %s", err)
		log.Println(err500)
		return WriteJson(writer, http.StatusInternalServerError, err500)
	}
	successMsg := fmt.Sprintf("Successfully created service: %s", service.ServiceName)
	log.Println(successMsg)
	return WriteJson(writer, http.StatusCreated, service)
}

//	@Summary	Get service by Id
//	@Id			3
//	@produce	application/json
//	@Param		serviceId	query		int	true	"serviceId"
//	@Failure	500		message	"Server Error: $err"
//	@Failure	400		message	"Error 400: $err"
//	@Failure	404		message	"404 page not found"
//	@Success	200		message	model.Service
//	@Router		/services/id/{ServiceId:[0-9]+} [get, head]
func (server *APIServer) handleGetServiceById(writer http.ResponseWriter, req *http.Request) error {
	serviceId, err := getServiceId(req)
	if err != nil {
		err400 := fmt.Sprintf("Error 400: %s", err)
		log.Println(err400)
		return WriteJson(writer, http.StatusBadRequest, err)
	}


	fmt.Printf("checking for serviceId %d", serviceId)
	service, err := server.db.GetServiceById(serviceId)
	if err != nil {
		err500 := fmt.Sprintf("Server Error: %s", err)
		log.Println(err500)
		return WriteJson(writer, http.StatusInternalServerError, err500)
	}

	return WriteJson(writer, http.StatusOK, service)
}

//	@Summary	Delete service by Id
//	@Id			4
//	@produce	application/json
//	@Param		serviceId	query		int	true	"serviceId"
//	@Failure	500		message	"Server Error: $err"
//	@Failure	400		message	"Error 400: $err"
//	@Failure	404		message	"404 page not found"
//	@Failure	404		message	"Could not find serviceId: $serviceId"
//	@Success	200		message	"deleted $serviceId"
//	@Router		/services/id/{ServiceId:[0-9]+} [delete, head]
func (server *APIServer) handleDeleteServiceById(writer http.ResponseWriter, req *http.Request) error {
	serviceId, err := getServiceId(req)
	if err != nil {
		err400 := fmt.Sprintf("Error 400: %s", err)
		log.Println(err400)
		return WriteJson(writer, http.StatusBadRequest, err)
	}

	fmt.Printf("Deleting service with serviceId %d", serviceId)
	numDeleted, err := server.db.DeleteServiceById(serviceId)
	if err != nil {
		err500 := fmt.Sprintf("Server Error: %s", err)
		log.Println(err500)
		return WriteJson(writer, http.StatusInternalServerError, err500)
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
//	@Failure	400		message	"Error 400: $err"
//	@Failure	404		message	"404 page not found"
//	@Failure	500		message	"Server Error: $err"
//	@Success	200		message	serviceVersions
//	@Router		/services/id/{ServiceId:[0-9]+} [get, head]
func (server *APIServer) handleGetServiceVersionsById(writer http.ResponseWriter, req *http.Request) error {
	serviceId, err := getServiceId(req)
	if err != nil {
		err400 := fmt.Sprintf("Error 400: %s", err)
		log.Println(err400)
		return WriteJson(writer, http.StatusBadRequest, err)
	}

	fmt.Printf("checking for serviceId %d", serviceId)
	serviceVersions, err := server.db.GetServiceVersionsById(serviceId)
	if err != nil {
		err500 := fmt.Sprintf("Server Error: %s", err)
		log.Println(err500)
		return WriteJson(writer, http.StatusInternalServerError, fmt.Sprintf("Server Error: %s", err500))
	}

	// TODO: implement 404 for serviceId not found

	return WriteJson(writer, http.StatusOK, serviceVersions)
}

//	@Summary	Get service by name
//	@Id			6
//	@produce	application/json
//	@Param		serviceName	query	string	true	"serviceName"
//	@Failure	500		message	"Server Error: $err"
//	@Failure	404		message	"404 page not found"
//	@Success	200		message	service
//	@Router		/services/name/{ServiceName:[a-zA-Z0-9]+} [get, head]
func (server *APIServer) handleGetServiceByName(writer http.ResponseWriter, req *http.Request) error {
	// TODO: validate mux var (400 error)
	serviceName := mux.Vars(req)["ServiceName"]
	log.Printf("Searching for service %s by name", serviceName)

	fmt.Printf("checking for serviceName %s", serviceName)
	service, err := server.db.GetServiceByName(serviceName)
	if err != nil {
		err500 := fmt.Sprintf("Server Error: %s", err)
		log.Println(err500)
		return WriteJson(writer, http.StatusInternalServerError, err500)
	}

	return WriteJson(writer, http.StatusOK, service)
}

//	@Summary	Get API health
//	@Id			6
//	@produce	application/json
//	@Success	200		message	"service is up and running"
//	@Router		/health [get, head]
func (server *APIServer) handleGetHealth(writer http.ResponseWriter, req *http.Request) error {
	return WriteJson(writer, http.StatusOK, "service is up and running")
}
