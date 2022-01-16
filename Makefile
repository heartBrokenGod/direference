wire :
	wire .

run : wire
	go run wire_gen.go app.go

test :
	curl localhost:8080/users