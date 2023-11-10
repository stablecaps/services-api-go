package models

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func makeRandomVersionsSlice(max int) string {
	num_versions := rand.Intn(max)
	if num_versions == 0 {
		num_versions = 1
	}

	versions_array := make([]string, num_versions)
	for idx := range versions_array {
		versions_array[idx] = "v" + strconv.Itoa(1+idx)
	}

	return strings.Join(versions_array[:], ",")
}

// TODO: validate for empty strings
type CreateServiceRequest struct {
	ServiceName string `json:"serviceName" validate:"required"`
	ServiceDescription string `json:"serviceDescription" validate:"required"`
}

// Primarily for testing
type CreateExplicitServiceRequest struct {
	ServiceId int `json:"serviceId"`
	ServiceName string `json:"serviceName" validate:"required"`
	ServiceDescription string `json:"serviceDescription" validate:"required"`
	ServiceVersions string `json:"serviceVersions" validate:"required"`
	CreatedAt string `json:"createdAt" validate:"required"`
}

// Response
type Service struct {
	ServiceId int `json:"serviceId"`
	ServiceName string `json:"serviceName"`
	ServiceDescription string `json:"serviceDescription"`
  	ServiceVersions string `json:"serviceVersions"`
	CreatedAt time.Time `json:"createdAt"`
}


func NewService(ServiceName, ServiceDescription string) *Service {
	return &Service{
		ServiceName: ServiceName,
		ServiceDescription: ServiceDescription,
		ServiceVersions: makeRandomVersionsSlice(5),
		CreatedAt: time.Now().UTC(),
	}
}

// Primarily for testing
func NewExplicitService(ServiceName, ServiceDescription, ServiceVersions string, CreatedAt time.Time) *Service {
	return &Service{
		ServiceName: ServiceName,
		ServiceDescription: ServiceDescription,
		ServiceVersions: ServiceVersions,
		CreatedAt: time.Now().UTC(),
	}
}