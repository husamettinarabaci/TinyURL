[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)

# TinyURL

API Based Url Shorter Application


### Test

```bash

make test

```

## Run Locally

```bash
  sudo docker run --name tinyurl -p 8080:8080 -d devhsmtek/tinyurl:latest
  or
  make server
```

## API Reference

#### Create a new user

```http
  POST http://0.0.0.0:8080/users
```

//Body raw:
```json
{
    "username" : "xxx",
    "full_name": "xxx",
    "email": "xxx",
    "password" : "xxx",
    "user_type":"FREE | PREMIUM"
}
```

#### Login

```http
  POST http://0.0.0.0:8080/users/login
```

//Body raw:
```json
{
    "username" : "xxx",
    "password" : "xxx"
}
```

#### Create a new transform

```http
  POST http://0.0.0.0:8080/transforms
```

//Body raw:
```json
{
    "longurl" : "xxx"
}
```
//Headers:
```json
{
    "Authorization": "Bearer <token>"
}
```

#### Get a transform

```http
  GET http://0.0.0.0:8080/transforms/:id
```

//Body raw:
```json
{
    "id" : "xxx"
}
```
//Headers:
```json
{
    "Authorization": "Bearer <token>"
}
```

#### Get all transforms

```http
  GET http://0.0.0.0:8080/transforms
```

//Query Params:
```json
page_id = x
page_size = x
```
//Headers:
```json
{
    "Authorization": "Bearer <token>"
}
```

## Directory Structure

─── TinyURL

    ├── api
    │   ├── main_test.go
    │   ├── middleware.go
    │   ├── middleware_test.go
    │   ├── server.go
    │   ├── token.go
    │   ├── transform.go
    │   ├── transform_test.go
    │   ├── user.go
    │   ├── user_test.go
    │   └── validator.go
    ├── app.env
    ├── db
    │   ├── migration
    │   │   ├── 000001_init_schema.down.sql
    │   │   ├── 000001_init_schema.up.sql
    │   │   └── docs
    │   ├── mock
    │   │   └── store.go
    │   ├── query
    │   │   ├── transforms.sql
    │   │   └── user.sql
    │   └── sqlc
    │       ├── db.go
    │       ├── main_test.go
    │       ├── models.go
    │       ├── querier.go
    │       ├── store.go
    │       ├── store_test.go
    │       ├── transforms.sql.go
    │       ├── transforms_test.go
    │       ├── user.sql.go
    │       └── user_test.go
    ├── Dockerfile
    ├── go.mod
    ├── go.sum
    ├── main.go
    ├── Makefile
    ├── README.md
    ├── sqlc.yaml
    ├── token
    │   ├── jwt_maker.go
    │   ├── jwt_maker_test.go
    │   ├── maker.go
    │   ├── paseto_maker.go
    │   ├── paseto_maker_test.go
    │   └── payload.go
    └── util
        ├── config.go
        ├── password.go
        ├── password_test.go
        ├── random.go
        └── usertype.go

## Tech Stack

![Go](https://img.shields.io/badge/Go-v1.19-blue)
![Go](https://img.shields.io/badge/Go-Migrate-blue)
![Go](https://img.shields.io/badge/Go-Viper-blue)
![Go](https://img.shields.io/badge/Go-Validator-blue)
![Go](https://img.shields.io/badge/Go-Mock-blue)
![Go](https://img.shields.io/badge/Go-Gin-blue)
![Go](https://img.shields.io/badge/Go-Testify-blue)
![Docker](https://img.shields.io/badge/Docker-passing-green)
![Aws RDS](https://img.shields.io/badge/Aws-Rds-blue)
![github](https://img.shields.io/badge/Github-Actions-green)
![API](https://img.shields.io/badge/API-http-blue)
![Test](https://img.shields.io/badge/Test-unit-green)
![CI/CD](https://img.shields.io/badge/CI%20CD-automation-green)
![Go](https://img.shields.io/badge/Go-Paseto-green)
![Go](https://img.shields.io/badge/Go-Jwt-green)
![Postgresql](https://img.shields.io/badge/Go-Pq-green)


## License

[MIT](https://choosealicense.com/licenses/mit/)


## Developers:

<img src="https://github.com/husamettinarabaci/husamettinarabaci/blob/main/hsmtek-logo.png?raw=true" width="200"/>

[@husamettinarabaci](https://www.github.com/husamettinarabaci)