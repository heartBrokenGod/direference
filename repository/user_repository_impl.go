package repository

import "github.com/heartBrokenGod/direference/models"

type UserRepositoryImpl struct {
	// in real case the repository may depend on db, cache,etc
	// this is just for the demo purpose
	users []models.User
}

func (ur *UserRepositoryImpl) GetUsers() ([]models.User, error) {
	return ur.users, nil
}

func NewUserRepoImpl() (*UserRepositoryImpl, error) {
	users := []models.User{
		models.NewUser("Neel", 25, "neelwanky@gmail.com"),
		models.NewUser("Nikhil", 27, "nikhilkumar.nv80@gmail.com"),
		models.NewUser("Naruto", 26, "naruto@uzumaki.com"),
	}

	return &UserRepositoryImpl{
		users: users,
	}, nil
}
