package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/heartBrokenGod/direference/service"
)

type HandlerImpl struct {
	*mux.Router
	userService service.UserService
}

func (handler *HandlerImpl) getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := handler.userService.GetUsers()
	if err != nil {
		// or any specific response
		w.WriteHeader(http.StatusInternalServerError) // returning for demo purpose only
		fmt.Fprint(w, err.Error())
	}
	// reurn successful response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	usersJson, _ := json.Marshal(users)
	fmt.Fprint(w, string(usersJson))
}

func NewHandlerImpl(userService service.UserService) (*HandlerImpl, error) {

	if userService == nil {
		return nil, errors.New("userservice dependency is nil")
	}

	router := mux.NewRouter()

	handler := &HandlerImpl{
		Router:      router,
		userService: userService,
	}

	// register the routes
	handler.HandleFunc("/users", handler.getUsers).Methods(http.MethodGet)

	return handler, nil
}
