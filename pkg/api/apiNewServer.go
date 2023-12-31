package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stablecaps/services-api-go/pkg/models"

	// _ "github.com/swaggo/http-swagger/example/gorilla/docs" // docs is generated by Swag CLI, you have to import it.
	httpSwagger "github.com/swaggo/http-swagger/v2"
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

func (server *APIServer) Run(port string) {
	router := mux.NewRouter()
	router.HandleFunc("/services", makeHTTPHandleFunc(server.handleGetAllServices)).Methods("GET", "HEAD")
	router.HandleFunc("/services/new", makeHTTPHandleFunc(server.handleCreateNewService)).Methods("POST", "HEAD")
	router.HandleFunc("/services/versions/{ServiceId:[0-9]+}", makeHTTPHandleFunc(server.handleGetServiceVersionsById)).Methods("GET", "HEAD")
	router.HandleFunc("/services/id/{ServiceId:[0-9]+}", makeHTTPHandleFunc(server.handleGetServiceById)).Methods("GET", "HEAD")
	router.HandleFunc("/services/id/{ServiceId:[0-9]+}", makeHTTPHandleFunc(server.handleDeleteServiceById)).Methods("DELETE", "HEAD")
	router.HandleFunc("/services/name/{ServiceName:[a-zA-Z0-9]+}", makeHTTPHandleFunc(server.handleGetServiceByName)).Methods("GET", "HEAD")
	router.HandleFunc("/health", makeHTTPHandleFunc(server.handleGetHealth)).Methods("GET", "HEAD")

	// docs
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		// httpSwagger.URL(fmt.Sprintf("http://localhost:%s/swagger/swagger.json", port)), // buggy
		httpSwagger.URL("http://localhost:8969/swagger/swagger.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)


	log.Println("API server running on port: ", server.listenAddr)
	log.Fatal(http.ListenAndServe(server.listenAddr, router))
}