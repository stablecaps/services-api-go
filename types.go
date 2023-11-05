package main

import (
	"math/rand"
	"strconv"
	"time"
)

type CreateServiceRequest struct {
	ServiceName string
	ServiceDescription string
}

type Service struct {
	ServiceId int
	ServiceName string
	ServiceDescription string
  	// ServiceVersions pq.StringArray
	CreatedAt time.Time
}

func makeRandomVersionsSlice(max int) []string {
	num_versions := rand.Intn(max)
	if num_versions == 0 {
		num_versions = 1
	}

	versions_array := make([]string, num_versions)
	for idx := range versions_array {
		versions_array[idx] = "v" + strconv.Itoa(1+idx)
	}

	return versions_array
}

func NewService(ServiceName, ServiceDescription string) *Service {
	return &Service{
		ServiceName: ServiceName,
		ServiceDescription: ServiceDescription,
		// ServiceVersions: makeRandomVersionsSlice(5),
		// TODO: check if uuid already exists
		CreatedAt: time.Now().UTC(),
	}
}