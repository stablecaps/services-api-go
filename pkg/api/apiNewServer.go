package api

import "github.com/stablecaps/services-api-go/pkg/models"

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
