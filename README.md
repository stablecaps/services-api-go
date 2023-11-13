# Dashboard Services Go API

## Overview
This is a timed interview test to create an api which retrieves a list of services written in Go. See docs directory for full brief. I have not had any golang coding experience, so this was a crash course in learning Golang over ~12 days and has been quite fun!


## Design Process.

1. I started writing an initial swagger spec using swaggerhub. this can be found `swagger/initial_ServicesDashboardAPI-1.0.0-unresolved.yaml`

2. I then looked at several crash courses on creating APIs in golang to get to grips with how to code the app.

3. Wrote the API and tried to ensure that responses did not contain information that would reveal the inner workings of the backend.

4. Made the api fail incorrect requests with useful messages fed back to user vs falling back to a default. This was to promote correct api usage.


## Usage instructions
1. Clone this repo using `git clone github.com:stablecaps/services-api-go.git`
2. There is a Makefile which can be utilised `./Makefile`
3. Make a mount point for postgres db `mkdir -p /tmp/data`
4. Start the database by first running `docker-compose up`. (Only the db is handled at the moment)
5. Create secrets file: `cp config_dev_secrets.env.template config_dev_secrets.env` and edit values in it so that API can access the DB. Also rename the DB secrets file `cp config_postgres_secrets.env.template config_postgres_secrets.env`. Note `DB_PASSWORD` must equal `POSTGRES_PASSWORD`
6. Build the project: `make build`
7. Start API server: `make run`
8. Test & populate DB: `make runpopdb`
9. Test List services & get service by id: `make pattest`

## Available endpoints:

Note: All endpoints check that  `Content-Type` is `application/json`

1. List services: `/services?limit=4&offset=0`. Can also order by column name using `/services?limit=4&offset=0&orderColName=serviceDescription&orderDir=desc`

2. Create a new service `/services/new`. Request body example:
```
{
    "ServiceName": "My New Service",
    "ServiceDescription": "Service to list coffee granules left in the jar"
}
```

3. Get service by integer serviceId `/services/versions/{serviceId:[0-9]+}`

4. Delete service by integer serviceId `/services/versions/{serviceId:[0-9]+}`

5. Get service by serviceName `/services/name/{ServiceName:[a-zA-Z0-9]+}`

6. Get service versions by integer serviceId `/services/versions/{serviceId:[0-9]+}`

7. Check App health `/health`


## Notes & things left to do
1. As time was short, I didn't want to waste any by potentially falling down a rabbit hole with the golang testing framework. Thus, I created a quick & dirty method to test API functionality instead. I will have a look at moving these to `go test`.

2. I've only tested HTTP response codes for the various endpoints so far. Still need to examine and test the response content.

3. I didn't spend too much time worrying about performance, rate limiting, etc as I think this can be handled by other infrastructure components like the API gateway. Set some sane defaults with respect to database connections though. I understand there is some performance benefit in allowing some idle connections. Ideally, this would be empirically tested under load.

4. Spent some time trying to create auto documentation via swaggo, but it needs a bit more TLC to get it over the line.

5. The docker-compose file only handles the DB at the moment. I would want to expand this so that it also handles a dockerised verion of the app andf integrates with Kong API gateway.

6. Written terraform code to deploy it to a production environment.

7. Improved config reading so that the app gets secrets via environment variables.

8. Support different API versions in the Dashboard-API itself.
