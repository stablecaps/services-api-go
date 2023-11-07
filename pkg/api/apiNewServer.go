package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stablecaps/services-api-go/pkg/models"
)

type APIServer struct {
	listenAddr string
	db         models.Dbase
}

func NewAPIServer(listenAddr string, db models.Dbase) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		db:         db,
	}
}

func (server *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/services", makeHTTPHandleFunc(server.handleGetAllServices)).Methods("GET")
	router.HandleFunc("/services/new", makeHTTPHandleFunc(server.handleCreateNewService)).Methods("POST")
	router.HandleFunc("/services/versions/{ServiceId:[0-9]+}", makeHTTPHandleFunc(server.handleGetServiceVersionsById)).Methods("GET")
	router.HandleFunc("/services/id/{ServiceId:[0-9]+}", makeHTTPHandleFunc(server.handleGetServiceById)).Methods("GET")
	router.HandleFunc("/services/id/{ServiceId:[0-9]+}", makeHTTPHandleFunc(server.handleDeleteServiceById)).Methods("DELETE")
	router.HandleFunc("/services/name/{ServiceName:[a-zA-Z0-9]+}", makeHTTPHandleFunc(server.handleGetServiceByName)).Methods("GET")
	router.HandleFunc("/health", makeHTTPHandleFunc(server.handleGetHealth)).Methods("GET")

	log.Println("API server running on port: ", server.listenAddr)
	log.Fatal(http.ListenAndServe(server.listenAddr, router))
}