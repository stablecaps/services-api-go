build:
	go build -C cmd/apiServerMain -o ../../bin/servicesapi
	go build -C cmd/populateDb -o ../../bin/populateDb
	go build -C cmd/paginateTest -o ../../bin/paginateTest
	go build -C cmd/patheticTester -o ../../bin/patheticTester

run:
	./bin/servicesapi

runpopdb:
	./bin/populateDb

runpage:
	./bin/paginateTest

doc:
	# annoyingly broken
	swag init --dir cmd/apiServerMain/ --output swagger --exclude data/ -g apiServerMain.go --parseDependency --parseDependencyLevel 3

test:
	go test $(go list ./... | grep -v /data/) -coverprofile .testCoverage.txt

pattest:
	./bin/patheticTester