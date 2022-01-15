package service

import "github.com/heartBrokenGod/direference/models"

type UserService interface {
	GetUsers() ([]models.User, error)
}
