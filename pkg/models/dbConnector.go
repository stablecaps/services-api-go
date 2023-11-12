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

// investigate prepared statements - looks like we prob do not need to do this
// https://www.reddit.com/r/golang/comments/6wll4z/lots_of_prepared_statements_how_do_i_deal_with/
func NewPostgresDb(userName, dbName, password, sslmode string, maxOpenConns, maxIdleConns int) (*PostgresDb, error) {
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=%s", userName, dbName, password, sslmode)
	log.Printf("connStr: %s", connStr)


	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// set db connection related settings
	// https://www.alexedwards.net/blog/configuring-sqldb
	db.SetMaxOpenConns(maxOpenConns) // Sane default
	db.SetMaxIdleConns(maxIdleConns) // try 2 for some performance gains
	db.SetConnMaxLifetime(5*time.Minute)

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