package tokenization

import (
	"efishery-endpoints/auth-app/config"
	"efishery-endpoints/auth-app/model"
	"efishery-endpoints/auth-app/repository/user"
	"errors"
	"fmt"
	"log"

	"github.com/golang-jwt/jwt"
)

type TokenLogiCInterface interface {
	ProduceToken(user model.User) (tokenz string, err error)
	ParseToken(tokenstr string) (tokenization model.Token, err error)
}

type TokenLogic struct {
	UserRepo user.UserRepoInterface
}

func NewTokenLogic() TokenLogiCInterface {
	return &TokenLogic{
		UserRepo: user.NewRepository(),
	}
}

func (t *TokenLogic) ProduceToken(user model.User) (tokenz string, err error) {
	userData, err := t.UserRepo.GetDetailByPasswordPhone(user.Phone, user.Password)
	if (err != nil || userData == model.User{}) {
		log.Println(err)
		err = errors.New("User not found")
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":  userData.Name,
		"phone": userData.Phone,
		"role":  userData.Role,
		// "password":  userData.Password,
		"timestamp": userData.Timestamp,
	})

	tokenz, err = token.SignedString([]byte(config.SECRET))

	if err != nil {
		log.Println(err)
		errors.New("Error in Produce String")
		return "", err
	}

	return tokenz, nil

}

func (t *TokenLogic) ParseToken(tokenstr string) (tokenization model.Token, err error) {
	token, err := jwt.Parse(tokenstr, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SECRET), nil
	})

	if err != nil {
		log.Println(err)
		errors.New("Error in Parse Token String")
		return tokenization, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		tokenization.Name = fmt.Sprint("", claims["name"])
		tokenization.Phone = fmt.Sprint("", claims["phone"])
		tokenization.Role = fmt.Sprint("", claims["role"])
		// tokenization.Password = fmt.Sprint("%v", claims["Password"])
		tokenization.Timestamp = fmt.Sprint("", claims["timestamp"])

	} else {

		log.Println(err)
		errors.New("Error in Parse Token String")
		return tokenization, err

	}
	return tokenization, nil

}
