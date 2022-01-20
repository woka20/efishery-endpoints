package user

import (
	"efishery-endpoints/auth-app/config"
	"efishery-endpoints/auth-app/model"
	"encoding/csv"
	"errors"
	"log"
	"os"
	"time"
)

type UserRepoInterface interface {
	InputData(model.User) (err error)
	GetListUser() (users []model.User, err error)
	GetDetailByPasswordPhone(pass, phone string) (user model.User, err error)
}

type UserRepo struct {
}

func NewRepository() UserRepoInterface {
	return &UserRepo{}
}

func (u *UserRepo) InputData(user model.User) (err error) {
	file, err := os.OpenFile(config.FILE_PATH, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()

	timestamp := time.Now()
	csvWriter := csv.NewWriter(file)
	csvWriter.Write([]string{user.Name, user.Phone, user.Role, user.Password, timestamp.Format("01 Jan 01 10:07 MST")})

	log.Println("Register user", user.Name, "at", timestamp.Format("01 Jan 01 10:07 MST"))

	csvWriter.Flush()

	return nil

}

func (u *UserRepo) GetListUser() (users []model.User, err error) {

	file, err := os.Open(config.FILE_PATH)
	if err != nil {
		log.Println(err)
		errors.New("Get Database Problem")

		return users, err

	}

	rows, err := csv.NewReader(file).ReadAll()

	if err != nil {
		log.Println(err)
		errors.New("Read CSV Problem")
		return users, err
	}

	for _, row := range rows {
		user := model.User{
			Name:      row[0],
			Phone:     row[1],
			Role:      row[2],
			Password:  row[3],
			Timestamp: row[4],
		}

		users = append(users, user)
	}

	return users, nil

}

func (u *UserRepo) GetDetailByPasswordPhone(pass, phone string) (user model.User, err error) {
	users, err := u.GetListUser()

	if err != nil {
		log.Println(err)
		err = errors.New("Get list user problem")

		return user, err
	}

	for _, user := range users {
		if user.Password == pass && user.Phone == phone {
			return user, nil
		}
	}

	return user, errors.New("No data match with phone and password ")

}
