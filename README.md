App architecture

### Commands
- `make db-run`- run database
- `make db-migrate` run migration
- `make db-migrate-test` run migration for testing

### Architecture 
https://miro.com/app/board/o9J_ko_Lm30=/

- connectors
  - postgreSQl connector: get data, put data
    - interface with dto
    - create user
    - get user
    - ger user by property
Will use this https://github.com/go-pg/pg/blob/ee50368e25f8/base.go#L296
    
- router
  - represents endpoints
    - post /create/user
    - get /user/{user_id}
    - get /users
- app.go
   - create router
   - create connector
   - create service
   - connect a connecter to service
   - connect a router with service
- logger
- health check

### TODO:
- Create database
- Add migrations
- Pack app into docker
- Add tests
- Add logger into all
- Integrate health check
- Use swagger for documentation
- turn of gingonic log output
- pretyfy zap output


### Assemptations
- name, second name and email can't be longer 40 symbols
- country code can't be longer than 5 symbols

### Ideas for improvements
- use migrations system f.i. https://github.com/pressly/goose
- add country code set

