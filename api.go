package main

import "net/http"

type APIServer struct {
	listenerAddr string
}

func NewAPIServer(listenerAddr string) *APIServer {
	return &APIServer{
		listenerAddr: listenerAddr,
	}
}

func (srv *APIServer) handleService(writer http.ResponseWriter, req *http.Request) error {
	return nil
}

func (srv *APIServer) handleListAllServices(writer http.ResponseWriter, req *http.Request) error {
	return nil
}

func (srv *APIServer) handleGetService(writer http.ResponseWriter, req *http.Request) error {
	return nil
}

func (srv *APIServer) handleCreateService(writer http.ResponseWriter, req *http.Request) error {
	return nil
}

func (srv *APIServer) handleDeleteService(writer http.ResponseWriter, req *http.Request) error {
	return nil
}