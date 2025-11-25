# Back End Test Golang BookingToGo



## Tech Stack

- Golang : https://github.com/golang/go
- MySQL (Database) : https://github.com/mysql/mysql-server

## Framework & Library

- Gorilla Mux (HTTP Framework) : https://github.com/gorilla/mux
- GORM (ORM) : https://github.com/go-gorm/gorm
- Viper (Configuration) : https://github.com/spf13/viper
- Golang Migrate (Database Migration) : https://github.com/golang-migrate/migrate
- Go Playground Validator (Validation) : https://github.com/go-playground/validator
- Logrus (Logger) : https://github.com/sirupsen/logrus

## Configuration

All configuration is in `config.json` file.

## API Spec

All API Spec is in `api` folder.

## Database Migration

All database migration is in `db/migrations` folder.

### Create Migration

```Shell
migrate create -ext sql -dir db/migrations table_status
```

### Run Migration

```shell
migrate -database "mysql://root:@tcp(localhost:3306)/article?charset=utf8mb4&parseTime=True&loc=Local" -path db/migrations up
```

### Run Appication


### Run unit test


```bash
go test -v ./test/
```


### Run Web Server
``` bash
go run cmd/web/main.go
```























