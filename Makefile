build:
	go build -C cmd/apiServerMain -o ../../bin/servicesapi
	go build -C cmd/populateDb -o ../../bin/populateDb

run:
	./bin/servicesapi

runpopdb:
	./bin/populateDb

test:
	go test $(go list ./... | grep -v /data/) -coverprofile .testCoverage.txt