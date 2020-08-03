App architecture

## Quick start
One command start:
- run localy with infrastructure in docker: `make app-run`
- run in docker: `make app-run-docker`

### Commands
- `make test` run tests
- `make test-cover` show test coverage

### API
- Run swagger `make swagger-ui`
- Open `http://localhost:8084/`
- Explore  `./swagger.yml`

### TODO:
- notify mechanism: use kafka or another centralized system
- Add logger into all


### Assemptations
- name, second name and email can't be longer 40 symbols
- country code can't be longer than 5 symbols

### Ideas for improvements
- use migrations system f.i. https://github.com/pressly/goose
- add country code set and country code checking

