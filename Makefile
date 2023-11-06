build:
	go build -o bin/servicesapi

run:
	./bin/servicesapi

test:
	go test $(go list ./... | grep -v /data/) -coverprofile .testCoverage.txt