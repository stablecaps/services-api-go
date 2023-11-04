build:
	go build -o bin/servicesapi

run:
	./bin/servicesapi

test:
	go test -v ./...