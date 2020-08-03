CWD=`pwd`

# app

app-run-docker: db-run
	docker-compose -f ci/docker-compose.yml up -d --build app

app-run: db-run app-build-localy
	./users-info --config config/config.json

app-build-localy:
	go build -o users-info app/app.go

down:
	docker-compose -f ci/docker-compose.yml down  || exit 1

# test
test: test-unit test-functional

# unit tests
test-unit:
	go test -v ./...

# functional tests
test-functional: app-run-docker
	cd test && go test

# tools
golint:
	 golangci-lint run

precommit: golint test

swagger-ui:
	docker-compose -f ci/docker-compose.yml up -d swaggerui

# db
migrations_path="migrations"
db_creds="host=localhost port=5432 dbname=users-info user=users-info password=users-info"

db-run: db-up db-create

db-connect:
	 psql ${db_creds}

db-migrate: db-create

db-migrate-test: db-create
	 psql ${db_creds} -f ${migrations_path}/add-test-data.sql

db-up:
	docker-compose -f ci/docker-compose.yml up -d postgres

db-create: db-ready
	 psql ${db_creds} -f ${migrations_path}/set-up.sql

db-ready:
	docker-compose -f ci/docker-compose.yml up start_dependencies



