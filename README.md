App architecture

### Commands
- `make db-run`- run database
- `make db-migrate` run migration
- `make db-migrate-test` run migration for testing
- `make test` run tests
- `make test-cover` show test cevarage

### API
- Run swagger `make swagger-ui`
- Open `http://localhost:8084/`
- Explore  `./swagger.yml`

### Architecture 
https://miro.com/app/board/o9J_ko_Lm30=/

- connectors
  - postgreSQl connector: get data, put data
    - create user
Will use this https://github.com/go-pg/pg/blob/ee50368e25f8/base.go#L296
    
- router
  - represents endpoints
    - post /create/user
- health check

### TODO:
- Pack app into docker
- Add tests
- Add logger into all
- Integrate health check
- Use swagger for documentation
- notify mechanizm
- add check if it's internal error or not
- add limitations for getUsers enteties count


### Assemptations
- name, second name and email can't be longer 40 symbols
- country code can't be longer than 5 symbols

### Ideas for improvements
- use migrations system f.i. https://github.com/pressly/goose
- add country code set

