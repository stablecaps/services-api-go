package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/lib/pq"
)

type Dbase interface {
	GetAllServices() ([]*Service, error)
	GetServiceById(int) (*Service, error)
	CreateService(*Service)  error
	UpdateService(*Service) error
	DeleteService(string) error
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

func (db *PostgresDb) CreateServiceTable() error {
	query := `create table if not exists services (
		ServiceId serial primary key,
		ServiceName varchar(50),
		ServiceDescription varchar(200),
		ServiceVersions varchar(50),
		CreatedAt timestamp
	)
	`
	_, err := db.db.Exec(query)
	return err
}


func (db *PostgresDb) GetAllServices() ([]*Service, error) {
	log.Println("Looking up services in DB")

	query := `select * from services`
	rows, err := db.db.Query(query)
	if err != nil {
		log.Printf("Error: %s", err)
		return nil, err
	}

	serviceSlice := []*Service{}
	for rows.Next() {
		service, err := scanService(rows)
		serviceSlice = append(serviceSlice, service)
		if err !=nil {
			log.Printf("Error: %s", err)
			return nil, err
		}

		serviceSlice = append(serviceSlice, service)
	}
	log.Println("DB lookup sucessful")
	return serviceSlice, nil
}

func (db *PostgresDb) GetServiceById(ServiceId int) (*Service, error) {
	fmt.Println("ServiceId: " + strconv.Itoa(ServiceId))
	query := `select * from services where ServiceId = $1`
	rows, err := db.db.Query(
		query,
		ServiceId,
	)
	if err != nil {
		log.Printf("Error: %s", err)
		return nil, err
	}
	for rows.Next() {
		return scanService(rows)
	}
	return nil, fmt.Errorf("service with ServiceId %d not found", ServiceId)
}

// func (db *PostgresDb) GetServiceByName(ServiceName string) (*Service, error) {
// 	return nil
// }

func (db *PostgresDb) GetServiceVersions(*Service) error {
	return nil
}

func (db *PostgresDb) CreateService(service *Service) error {
	log.Println("Creating new service in DB")

	query := `insert into services
	(ServiceName, ServiceDescription, ServiceVersions, CreatedAt)
	values ($1, $2, $3, $4)`

	resp, err := db.db.Query(
		query,
		service.ServiceName,
		service.ServiceDescription,
		service.ServiceVersions,
		service.CreatedAt,
	)

	if err != nil {
		log.Printf("Error: %s", err)
		return err
	}

	fmt.Printf("%+v\n", resp)
	log.Printf("Created new service <%s> in DB", service.ServiceName)
	return nil
}

func (db *PostgresDb) UpdateService(*Service) error {
	return nil
}

func (db *PostgresDb) DeleteService(ServiceId string) error {
	return nil
}

func scanService(rows *sql.Rows) (*Service, error) {
	service := new(Service)
	err := rows.Scan(
		&service.ServiceId,
		&service.ServiceName,
		&service.ServiceDescription,
		&service.ServiceVersions,
		&service.CreatedAt,
	)
	return service, err
}
