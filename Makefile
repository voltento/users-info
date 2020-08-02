CWD=`pwd`


# app

app-run: db-run
	docker-compose -f ci/docker-compose.yml up app

# tools
golint:
	 golangci-lint run

swagger-ui:
	docker-compose -f ci/docker-compose.yml up -d swaggerui

# db
migrations_path="migrations"
db_creds="host=localhost port=5432 dbname=users-info user=users-info password=users-info"

db-run: db-up db-create

db-down:
	docker-compose -f ci/docker-compose.yml down  || exit 1

db-connect:
	 psql ${db_creds}

db-migrate: db-create

db-migrate-test: db-create
	 psql ${db_creds} -f ${migrations_path}/add-test-data.sql

db-up: db-down
	docker-compose -f ci/docker-compose.yml up -d postgres

db-create: db-ready
	 psql ${db_creds} -f ${migrations_path}/set-up.sql

db-ready:
	docker-compose -f ci/docker-compose.yml up start_dependencies

# test
test:
	go test -v ./...

test-cover:
	go test -coverprofile=coverage.out -v ./... && \
    go tool cover -html=coverage.out



