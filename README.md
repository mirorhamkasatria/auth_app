# **Instructions**
Auth App repository user management

### **Service Feature**

- migration database file
- linter
- unit test
- api test (postman)


## **A. Migrate and Import Database**

1. run command `make migrate.up` to migrate database

## **B. How To Run**

1. load golang module with command `go mod tidy`
2. copy environtment with command `cp env.example .env`
3. run code with command `make run.serve`
4. hit running service. ex : `http://localhost:5005`

## **C. Unit Test**
- user library mockery
1. generate mock file with command `mockery -inpkg -all -recursive -case snake -testonly`
2. create unit test form mockery
3. run test golang with command `go test ./...` or run from test menu in editor

## **D. Linter**
- linter with `github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2`
1. Linting code with command `golangci-lint run`
## **E. API Documentation**
API documentation
via doc api `https://api.postman.com/collections/16010179-ad0998c2-4ee8-4534-9894-f4bf70cfd50a?access_key=PMAT-01HNPC2SYP5GYHRMEN73NWCQGA`
1. Login API
```
curl --location 'localhost:8080/api/v1/users/login' \
--header 'Content-Type: application/json' \
--data '{
    "email": "email",
    "password": "password"
}'
```
2. Register API
```
curl --location 'localhost:8080/api/v1/users/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "first_name": "first name ",
    "last_name": "last name",
    "email": "dummyemail@dummy.com",
    "password": "password1"
}'
```

### **Work with Migration**

-Golang migrate with `https://github.com/golang-migrate/migrate`

1. download goalng migrate with command `go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`
2. create golang file with command `migrate create -ext sql -dir database/migrations '{table_name}'`

### **About Repository Root**
```
.
├── .env            [environtment file]
├── main.go         [main file]
├── README.md       [readme file]
├── Makerfile       [maker file]
├── database
|   └── migrations
|       └── [all migraion files]
├── configs
|   └── [all configuration]
├── route
|   ├── new_server.go   [server initialization]
|   └── route.go        [all route/path url]
├── app
|   ├── controllers
|   |   └── [handling all controllers]
|   ├── services
|   |   └── [handling all service]
|   ├── models
|   |   └── [handling all models]
|   ├── repositories
|   |   └── [handling all repositories]
|   └── transport 
|       ├── mappers
|       |   └── [handling all mappers from/to models] 
|       ├── request
|       |   └── [handling all request body] 
|       └── response
|           └── [handling all response body] 
|       
└── pkg
    └──utils
       └── [handling general reusable function] 

```                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     