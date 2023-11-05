package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateService(*Service) error
	DeleteService(string) error
	UpdateService(*Service) error
	GetServiceById(string) (*Service, error)
}

type PostgreStore struct {
	db *sql.DB
}

func NewPostgreStore() (*PostgreStore, error) {
	// TODO: make password secret
	connStr := "user=postgres dbname=postgres password=mysecretpassword sslmode=disable"

	db, error := sql.Open("postgres", connStr)
	if error != nil {
		log.Fatal(error)
		return nil, error
	}

	if error := db.Ping(); error != nil {
		return nil, error
	}

	log.Println("Connected to db..")
	return &PostgreStore{
		db: db,
	}, nil

}

func (db *PostgreStore) Init() error {

}

func (db *PostgreStore) CreateAccountTable() error {
	query := "create table "

}








func (db PostgreStore) ListAllServices(*Service) error {
	return nil
}

func (db PostgreStore) GetServiceById(ServiceId string) (*Service, error) {
	return nil, nil
}

func (db PostgreStore) GetServiceVersions(*Service) error {
	return nil
}

func (db PostgreStore) CreateService(*Service) error {
	return nil
}

func (db PostgreStore) UpdateService(*Service) error {
	return nil
}

func (db PostgreStore) DeleteService(ServiceId string) error {
	return nil
}

