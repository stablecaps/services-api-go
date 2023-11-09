package models

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/lib/pq"
)

func (db *PostgresDb) CreateTable() error {
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
func (db *PostgresDb) GetAllServices(orderColName, orderDirection string, limit, offset int) ([]*Service, error) {

	log.Println("Looking up services in DB")

	// check to see if this is good against sql injection
	// example: SELECT * FROM services ORDER BY serviceId ASC LIMIT 4 OFFSET 0
	query := fmt.Sprintf("SELECT * FROM services ORDER BY %s %s LIMIT %d OFFSET %d", orderColName, orderDirection, limit, offset)

	log.Printf("SELECT * FROM services ORDER BY %s %s LIMIT %d OFFSET %d",
		orderColName,
		orderDirection,
		limit,
		offset)

	rows, err := db.db.Query(query)
	if err != nil {
		log.Printf("Error: %s", err)
		return nil, err
	}
	defer rows.Close()

	serviceSlice := []*Service{}
	println()
	for rows.Next() {
		println(" >>>rows", rows)
		service, err := scanService(rows)
		if err !=nil {
			log.Printf("Error: %s", err)
			return nil, err
		}

		serviceSlice = append(serviceSlice, service)
	}

	for i, v := range serviceSlice {
		fmt.Printf("%d - %v\n", i, v.ServiceName)
	}


	log.Println("DB lookup sucessful")
	return serviceSlice, nil
}

func (db *PostgresDb) CreateNewService(service *Service) error {
	log.Println("Creating new service in DB")

	query := `insert into services
	(ServiceName, ServiceDescription, ServiceVersions, CreatedAt)
	values ($1, $2, $3, $4)`

	row, err := db.db.Query(
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
	defer row.Close()

	fmt.Printf("%+v\n", row)
	log.Printf("Created new service %s in DB", service.ServiceName)
	return nil
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

func (db *PostgresDb) DeleteServiceById(ServiceId int) (int64, error) {
	log.Println("Deleting new service in DB")

	query := `delete from services where ServiceId = $1`
	res, err := db.db.Exec(
		query,
		ServiceId,
	)
	if err != nil {
		log.Printf("Error: %s", err)
		return 0, err
	}

	numDeleted, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error: %s", err)
		return 0, err
	}
	log.Printf("Number of rows deleted: %d", numDeleted)

	return numDeleted, err
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
	println(">>> xxx", service.ServiceName)
	return service, err
}
