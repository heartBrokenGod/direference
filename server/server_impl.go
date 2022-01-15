package server

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/heartBrokenGod/direference/config"
	"github.com/heartBrokenGod/direference/handler"
)

type ServerImpl struct {
	handler handler.Handler
	config  *config.Config
	server  *http.Server
}

func (s *ServerImpl) Run() error {
	err := s.server.ListenAndServe() // start the server
	if err != nil {
		return err
	}
	return nil
}

func NewServer(h handler.Handler, c *config.Config) (*ServerImpl, error) {
	if h == nil {
		return nil, errors.New("handler.Handler dependency is nil")
	} else if c == nil {
		return nil, errors.New("*config.Config dependency is nil")
	}

	// create the http server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", c.Port),
		Handler: h,
	}

	return &ServerImpl{
		handler: h,
		config:  c,
		server:  server,
	}, nil
}
