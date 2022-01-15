package repository

import "github.com/heartBrokenGod/direference/models"

type UserRepository interface {
	GetUsers() ([]models.User, error)
}
