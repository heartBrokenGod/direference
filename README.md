# direference
This repository contains dependency injection reference example in golang, for the golang pune meetup talk.

## To install wire
if you are on go version 1.17 or above, type the following command in the terminal.
```bash
go install github.com/google/wire/cmd/wire@latest
```

## To Get all the dependencies
Type the following command in the terminal.
```bash
go mod tidy
```
## To Run  the application
Type the following command in the terminal.
```bash
wire . && go run wire_gen.go app.go
```
## To Test the application
Type the following command in the terminal.
```bash
curl localhost:8080/users
```