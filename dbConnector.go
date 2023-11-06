package main

import (
	"database/sql"
	"log"
)
type Dbase interface {
	GetAllServices() ([]*Service, error)
	GetServiceByName(string) (*Service, error)
	GetServiceById(int) (*Service, error)
	DeleteServiceById(int) error
	GetServiceVersionsById(int) (string, error)
	CreateService(*Service)  error
}

type PostgresDb struct {
	db *sql.DB
}

func NewPostgresDb() (*PostgresDb, error) {
	// TODO: make password secret
	connStr := "user=postgres dbname=postgres password=mysecretpassword sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Printf("Error: %s", err)
		return nil, err
	}

	log.Println("Connected to db..")
	return &PostgresDb{
		db: db,
	}, nil

}

func (db *PostgresDb) Init() error {
	return db.CreateServiceTable()
}