//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/heartBrokenGod/direference/config"
	"github.com/heartBrokenGod/direference/handler"
	"github.com/heartBrokenGod/direference/repository"
	"github.com/heartBrokenGod/direference/server"
	"github.com/heartBrokenGod/direference/service"
)

var serverProviderSet = wire.NewSet(
	config.NewConfig,

	repository.NewUserRepoImpl,
	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),

	service.NewUserServiceImpl,
	wire.Bind(new(service.UserService), new(*service.UserServiceImpl)),

	handler.NewHandlerImpl,
	wire.Bind(new(handler.Handler), new(*handler.HandlerImpl)),

	server.NewServer,
	wire.Bind(new(server.Server), new(*server.ServerImpl)),
)

func InitServer() (server.Server, error) {

	wire.Build(serverProviderSet)

	return nil, nil
}
