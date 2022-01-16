package main

import (
	"github.com/heartBrokenGod/direference/config"
	"github.com/heartBrokenGod/direference/handler"
	"github.com/heartBrokenGod/direference/repository"
	"github.com/heartBrokenGod/direference/server"
	"github.com/heartBrokenGod/direference/service"
	"go.uber.org/dig"
)

func InitServer() (server.Server, error) {

	var serverIns server.Server

	// create the container for the providers
	container := dig.New()

	// add the providers necessary for the initialization of the server
	container.Provide(config.NewConfig)
	container.Provide(repository.NewUserRepoImpl, dig.As(new(repository.UserRepository)))
	container.Provide(service.NewUserServiceImpl, dig.As(new(service.UserService)))
	container.Provide(handler.NewHandlerImpl, dig.As(new(handler.Handler)))
	container.Provide(server.NewServer, dig.As(new(server.Server)))

	err := container.Invoke(func(server server.Server) {
		serverIns = server
	})

	if err != nil {
		return nil, err
	}

	// return the initialized server
	return serverIns, nil

}
