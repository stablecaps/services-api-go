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
3. Start the database by first running `docker-compose up`. (Only the db is handled at the moment)
4. Create secrets file: `cp config_dev_secrets.env.template config_dev_secrets.env` and edit values in it so that API can access the DB.
5. Build the project: `make build`
6. Start API server: `make run`
7. Test & populate DB: `make runpopdb`
8. Test pagination (offset & limit): `make runpage`
9. Test List services & get service by id: `make pattest`



## Things to do


2. Have not finished all the testing modules. I didn't have enough time to get to grips with the golang testing framework, so instead created a quick & dirty method to test API functionality.

3. I didn't spend too much time worrying about performance, rate limiting, etc as I think this can be handled by other infrastructure components like an API gateway.

4. Spent some time trying to create auto documentation via swaggo, but it needs a bit more TLC to get it over the line.

5. Submitted a fully working docker-compose file. Just handles the DB at the mo

6. Written terraform code to deploy it to a production environment.

7. I realise there are probably much better go programmers out there than me so my chances in this are slim. Thus, I mainly used this as an opportunity to learn about golang!