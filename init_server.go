package main

import (
	"github.com/heartBrokenGod/direference/config"
	"github.com/heartBrokenGod/direference/handler"
	"github.com/heartBrokenGod/direference/repository"
	"github.com/heartBrokenGod/direference/server"
	"github.com/heartBrokenGod/direference/service"
)

func InitServer() (server.Server, error) {
	config := config.NewConfig()

	repo, err := repository.NewUserRepoImpl()
	if err != nil {
		return nil, err
	}
	service, err := service.NewUserServiceImpl(repo)
	if err != nil {
		return nil, err
	}
	handler, err := handler.NewHandlerImpl(service)
	if err != nil {
		return nil, err
	}

	server, err := server.NewServer(handler, config)
	if err != nil {
		return nil, err
	}

	return server, nil

}
