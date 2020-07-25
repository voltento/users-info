
# tools
golint:
	 golangci-lint run

# db
migrations_path="migrations"
db_creds="host=localhost port=5432 dbname=users-info user=users-info password=users-info"

db-run: db-up

db-connect:
	 psql ${db_creds}

db-migrate: db-create

db-migrate-test: db-create
	 psql ${db_creds} -f ${migrations_path}/add-test-data.sql

db-up: db-down
	docker-compose -f ci/docker-compose-run-app.yml up -d

db-down:
	docker-compose -f ci/docker-compose-run-app.yml down

db-create:
	 psql ${db_creds} -f ${migrations_path}/set-up.sql



