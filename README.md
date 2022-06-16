# REST API Server in Golang

- This API server provides endpoints to create,read,update & delete users and their repositories (like Github).
  
[![Go Report Card](https://github.com/gojp/goreportcard)](https://github.com/gojp/goreportcard)

## To Start API Server
```$ git clone https://github.com/pkbhowmick/go-rest-api.git```

```$ cd go-rest-api```

```$ go install```

```$ go-rest-api version``` [print the version of the api server]

```$ go-rest-api start```  [run the api server]

## Command to run unit test for API endpoints
```$ cd api```

```$ go test```

## Commands to run API server in docker container
```shell
$ docker build -t <image_name> .
$ docker run -p 8080:8080 <given_image_name> # to start the server with default config
$ docker run -p 8080:8080 <given_image_name> start -a=false # to start the server without authentication
```

## Commands to pull image from docker hub and run locally
```shell
$ docker pull pkbhowmick/go-rest-api
$ docker run -p 8080:8080 pkbhowmick/go-rest-api # to start the server with default config
$ docker run -p 8080:8080 pkbhowmick/go-rest-api start -a=false # to start the server without authentication
```


## Data Model

- User Model
``````
type User struct {
	ID           string       `json:"id"`
	FirstName    string       `json:"firstName"`
	LastName     string       `json:"lastName"`
	Repositories []Repository `json:"repositories"`
	CreatedAt    time.Time    `json:"createdAt"`
}
``````
- Repository Model
``````
type Repository struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Visibility string    `json:"visibility"`
	Star       int       `json:"star"`
}
``````

## Available API Endpoints

|  Method | API Endpoint  | Authentication Type | Description |
|---|---|---|---|
|POST| /api/login | Basic | Return jwt token in response for successful authentication
|GET| /api/users | Basic or Bearer token | Return a list of all users in response| 
|GET| /api/users/{id} | Basic or Bearer token| Return the data of given user id in response| 
|POST| /api/users | Basic or Bearer token |Add an user in the database and return the added user data in response | 
|PUT| /api/users/{id} | Basic or Bearer token |Update the user and return the updated user info in response| 
|DELETE| /api/users/{id} | Basic or Bearer token |Delete the user and return the deleted user data in response| 

## Available Flags

| Flag | Shorthand | Default value | Example | Description
|---|---|---|---|---|
|port|p|8080| go-rest-api start --port=8090 | Start API server in the given port otherwise in default port
|auth|a|true| go-rest-api start --auth=false | If true impose authentication on API server otherwise bypass it

## Environment Variables

| Name | Default(in docker image) | Description
|---|---|---|
ADMIN_USER | admin | API server admin username
ADMIN_PASS | demo | API server admin password
SIGNING_KEY | veryverysecretkey | API server jwt token signing key


## Sample Curl commands without authentication

Run API server without authentication

```shell
$ go-rest-api start --port=8080 --auth=false
``` 

Get all users information

```shell
$ curl -X GET http://localhost:8080/api/users
``` 

Get user information with id 1

```shell
$ curl -X GET http://localhost:8080/api/users/1
```

Create user with given id

```shell
$ curl -X POST  -H "Content-Type:application/json" -d '{"id":"6","firstName":"testfirst","lastName":"testlast"}' http://localhost:8080/api/users
``` 

Modify user data with gigen id

```shell
$ curl -X PUT  -H "Content-Type:application/json" -d '{"firstName":"test","lastName":"test"}' http://localhost:8080/api/users/1
``` 

Delete user with given id

```shell
$ curl -X DELETE http://localhost:8080/api/users/1
``` 

## Sample Curl commands with Basic authentication

```shell
$ export ADMIN_USER=admin
```

```shell
$ export ADMIN_PASS=demo
```

Run API server with authentication

```shell
$ go-rest-api start --port=8080 --auth=true
``` 

Get all users information

```shell
$ curl -X GET --user admin:demo http://localhost:8080/api/users
``` 

Get user information with id 1

```shell
$ curl -X GET --user admin:demo http://localhost:8080/api/users/1
``` 

Create user with given id

```shell
$ curl -X POST  --user admin:demo -H "Content-Type:application/json" -d '{"id":"6","firstName":"testfirst","lastName":"testlast"}' http://localhost:8080/api/users
``` 

Modify user data with given id

```shell
$ curl -X PUT  --user admin:demo -H "Content-Type:application/json" -d '{"firstName":"test","lastName":"test"}' http://localhost:8080/api/users/1
``` 

Delete user with given id

```shell
$ curl -X DELETE --user admin:demo http://localhost:8080/api/users/1
``` 

## Sample Curl commands with Bearer token(JWT token) authentication

```shell
$ export ADMIN_USER=admin
```

```shell
$ export ADMIN_PASS=demo
```

```shell
$ export SIGNING_KEY=veryverysecretkey
```

Run API server with authentication

```shell
$ go-rest-api start --port=8080 --auth=true
``` 

Get jwt token via login with basic authentication

```shell
$ curl -X POST --user admin:demo  http://localhost:8080/api/login
``` 

Get all users information

```shell
$ curl -X GET -H "Authorization: Bearer <jwt_token>"  http://localhost:8080/api/users
``` 

Get user information with id 1

```shell
$ curl -X GET -H "Authorization: Bearer <jwt_token>" http://localhost:8080/api/users/1
``` 

Create user with given id

```shell
$ curl -X POST -H "Authorization: Bearer <jwt_token>" -H "Content-Type:application/json" -d '{"id":"6","firstName":"testfirst","lastName":"testlast"}' http://localhost:8080/api/users
``` 

Modify user data with given id

```shell
$ curl -X PUT -H "Authorization: Bearer <jwt_token>" -H "Content-Type:application/json" -d '{"firstName":"test","lastName":"test"}' http://localhost:8080/api/users/1 
``` 

Delete user with given id

```shell
$ curl -X DELETE -H "Authorization: Bearer <jwt_token>" http://localhost:8080/api/users/1
``` 
