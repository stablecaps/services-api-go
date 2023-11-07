build:
	go build -C cmd/apiServerMain -o ../../bin/servicesapi

run:
	./bin/servicesapi

test:
	go test $(go list ./... | grep -v /data/) -coverprofile .testCoverage.txt