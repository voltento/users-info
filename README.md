The application provides http interface for storing/modifying users' data.

## Quick start
One command start:
- `make app-run` run localy with infrastructure in docker
- `make app-run-docker` run in docker
- `make test` run all tests both unit and functional. It starts the app in docker

### API
- Run swagger `make swagger-ui`
- Open `http://localhost:8084/`
- Explore  `./swagger.yml`

### Assemptations
- name, second name and email can't be longer 40 symbols
- country code can't be longer than 5 symbols

### Ideas for improvements
- add country code set and country code checking
- notify mechanism: use kafka or another centralized system for scalability
- add validation for country code


