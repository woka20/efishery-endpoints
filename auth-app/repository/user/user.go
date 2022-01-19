package user

import (
	"efishery-endpoints/auth-app/config"
	"efishery-endpoints/auth-app/model"
	"log"
	"os"
)

type UserRepoInterface interface {
}

type UserRepo struct {
}

func NewRepository() UserRepoInterface {
	return &UserRepo{}
}

func (u *UserRepo) InputData(user model.User) error {
	file, err := os.OpenFile(config.FILE_PATH, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()
	return nil

}
