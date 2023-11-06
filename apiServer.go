package main

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