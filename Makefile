build:
	go build -C cmd/apiServerMain -o ../../bin/servicesapi
	go build -C cmd/populateDb -o ../../bin/populateDb

run:
	./bin/servicesapi

runpopdb:
	./bin/populateDb

doc:
	# annoyingly broken
	swag init --dir cmd/apiServerMain/ --output swagger --exclude data/ -g apiServerMain.go --parseDependency --parseDependencyLevel 3

test:
	go test $(go list ./... | grep -v /data/) -coverprofile .testCoverage.txt