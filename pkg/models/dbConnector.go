package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type Dbase interface {
	GetAllServices(string, string, int, int) ([]*Service, error)
	GetServiceByName(string) (*Service, error)
	GetServiceById(int) (*Service, error)
	DeleteServiceById(int) (int64, error)
	GetServiceVersionsById(int) (string, error)
	CreateNewService(*Service)  (*Service, error)
}

type PostgresDb struct {
	db *sql.DB
}

func NewPostgresDb(userName, dbName, password, sslmode string) (*PostgresDb, error) {
	log.Printf("hello")
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=%s", userName, dbName, password, sslmode)
	log.Printf("connStr: %s", connStr)


	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// set db connection related settings
	db.SetMaxOpenConns(20) // Sane default
	db.SetMaxIdleConns(0)
	db.SetConnMaxLifetime(time.Nanosecond)

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
	return db.CreateTable()
}