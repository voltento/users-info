App architecture

## Quick start
One command start:
- `make app-run` run localy with infrastructure in docker
- `make app-run-docker` run in docker
- `make test` run all tests both unit and functional. It starts the app in docker

### API
- Run swagger `make swagger-ui`
- Open `http://localhost:8084/`
- Explore  `./swagger.yml`

### TODO:
- add functional tests
- notify mechanism: use kafka or another centralized system
- add validation for country code

### Assemptations
- name, second name and email can't be longer 40 symbols
- country code can't be longer than 5 symbols

### Ideas for improvements
- use migrations system f.i. https://github.com/pressly/goose
- add country code set and country code checking

