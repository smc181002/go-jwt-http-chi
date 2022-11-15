# Golang JWT implementation with `net/http` library

## Introduction

Authentication is one of the common things to be considered
in building an API. Here is a simple code implementing JWT
application in a webserver built with the help go `net/http`
and using [go-chi](https://github.com/go-chi/chi) for 
handling the routing.

## Installation

Go version `1.18`

Start cloning the repository with 

```bash
git clone https://github.com/smc181002/go-jwt-http-chi.git
```

### Package Installation

The following package have been used in the repository

* go-chi's chi/v5 router
* go-chi's cors middleware to handle the cors headers
* go-chi's render to render JSON data in the API handlers
* golang-jwt's jwt/v4 package

```bash
go get github.com/go-chi/chi/v5 v5.0.7
go get github.com/go-chi/cors v1.2.1
go get github.com/golang-jwt/jwt/v4 v4.4.2
go get github.com/go-chi/render v1.0.1

go mod tidy
```

### Starting the server

Use `go run main.go` to run the server

## Testing application

The user name is fixed as `smc` for this demo but in an 
actual application, the username will be fetched from 
database and the username may be replaced or combined 
with other parameters like roles.

### Using cURL

cURL is a tool installed in linux distributions by default 
used to fetch data from an endpoint.

**Get JWT token**

we can get the JWT token by sending `POST` request to 
`/get-token` endpoint.

```bash
curl \
--data '{"uname": "smc", "password": "passwd@123"}' \
-H 'Content-Type: application/json' \
http://localhost:8080/get-token
```

save the output value in environmental variables from the 
response JSON in the above request which will be something 
like below.

```json
{"token":"eyJh...XVCJ9.eyJ1...4ODQ0NH0.BZZltUp...2-Urx0HUfb-I","error":null}
```


To request public endpoint, send `GET` request to `/posts`

```bash
curl http://localhost:8080/posts/
To get the protected routes, we can send the token through 
Authorization Header as Bearer token.
```

```bash
curl \
-H "Authorization: Bearer $TOKEN" \
http://localhost:8080/secret-posts
```
