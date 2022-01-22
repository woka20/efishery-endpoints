package user

import (
	"efishery-endpoints/auth-app/model"
	"efishery-endpoints/auth-app/repository/user"
	"errors"
	"log"

	"github.com/sethvargo/go-password/password"
)

type UserLogicInterface interface {
	AddNewUser(user model.User) (pass string, err error)
	ProducePassword() (pass string, errr error)
	CheckUser(user model.User) (bool, error)
}

type UserLogic struct {
	UserRepo user.UserRepoInterface
}

func NewLogicUser() UserLogicInterface {
	return &UserLogic{
		UserRepo: user.NewRepository(),
	}
}

func (u *UserLogic) AddNewUser(user model.User) (pass string, err error) {
	exist, err := u.CheckUser(user)
	if err != nil {
		log.Println(err)
		errors.New("Cannot check user exist")

		return "", err
	}

	if exist {
		return "", errors.New("User already exists")

	} else {
		pass, err := u.ProducePassword()
		if err != nil {
			log.Println(err)
			errors.New("Cannot produce password")

			return "", err
		}

		user.Password = pass
		err = u.UserRepo.InputData(user)

		if err != nil {
			return "", err
		}

	}

	return user.Password, nil

}

func (u *UserLogic) ProducePassword() (pass string, err error) {
	pass, err = password.Generate(4, 1, 1, true, false)
	if err != nil {
		log.Println(err)
		errors.New("Create Password Problem")

		return "", err
	}

	return pass, nil

}

func (u *UserLogic) CheckUser(user model.User) (bool, error) {
	users, err := u.UserRepo.GetListUser()

	if err != nil {
		log.Println(err)

		errors.New("Get List Problem")

		return true, err
	}

	for _, detail := range users {
		if user.Phone == detail.Phone {
			return true, nil
		}
	}

	return false, nil

}
