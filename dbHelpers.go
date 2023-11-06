package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/lib/pq"
)

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

////////////////////////////////////////////////////////
////////////////////////////////////////////////////////
func (db *PostgresDb) GetAllServices() ([]*Service, error) {
	log.Println("Looking up services in DB")

	query := `select * from services`
	rows, err := db.db.Query(query)
	if err != nil {
		log.Printf("Error: %s", err)
		return nil, err
	}
	defer rows.Close()

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

func (db *PostgresDb) GetServiceByName(ServiceName string) (*Service, error) {
	fmt.Println("\nServiceName: " + ServiceName)
	query := `select * from services where ServiceName = $1`
	rows, err := db.db.Query(
		query,
		ServiceName,
	)
	if err != nil {
		log.Printf("Error: %s", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		return scanService(rows)
	}
	return nil, fmt.Errorf("service with ServiceName %s not found", ServiceName)
}

func (db *PostgresDb) GetServiceById(ServiceId int) (*Service, error) {
	fmt.Println("\nServiceId: " + strconv.Itoa(ServiceId))
	query := `select * from services where ServiceId = $1`
	rows, err := db.db.Query(
		query,
		ServiceId,
	)
	if err != nil {
		log.Printf("Error: %s", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		return scanService(rows)
	}
	return nil, fmt.Errorf("service with ServiceId %d not found", ServiceId)
}

func (db *PostgresDb) DeleteServiceById(ServiceId int) error {
	log.Println("Deleting new service in DB")

	query := `delete from services where ServiceId = $1`
	_, err := db.db.Query(
		query,
		ServiceId,
	)
	return err
}

func (db *PostgresDb) GetServiceVersionsById(ServiceId int) (string, error) {
	fmt.Println("\nRetrieving version info for ServiceId: " + strconv.Itoa(ServiceId))
	query := `select ServiceVersions from services where ServiceId = $1`
	row := db.db.QueryRow(
		query,
		ServiceId,
	)
	if err := row.Err(); err != nil {
		return "", err
	}


	var serviceVersions string
	err := row.Scan(&serviceVersions)
	if err == sql.ErrNoRows {
		return "", fmt.Errorf("service with ServiceId %d not found", ServiceId)
	} else if err != nil {
		log.Printf("Error: %s", err)
		return "", err
	}

	fmt.Println("versions: ", serviceVersions)
	return serviceVersions, nil
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
	log.Printf("Created new service %s in DB", service.ServiceName)
	return nil
}

////////////////////////////////////////////////////////
////////////////////////////////////////////////////////
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
