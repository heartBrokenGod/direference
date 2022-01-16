# direference
This repository contains dependency injection reference example in golang, for the golang pune meetup talk.

## To Get all the dependencies
Type the following command in the terminal.
```bash
go mod tidy
```
## To Run  the application
Type the following command in the terminal.
```bash
go run init_server.go app.go
```
## To Test the application
Type the following command in the terminal.
```bash
curl localhost:8080/users
```
 ## To look at the dependency injection using wire
 Switch to the branch : [dependency_injection_using_wire](https://github.com/heartBrokenGod/direference/tree/dependency_injection_using_wire)