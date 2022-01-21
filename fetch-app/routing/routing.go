package routing

import (
	authmod "efishery-endpoints/auth-app/model"
	"efishery-endpoints/fetch-app/config"
	"efishery-endpoints/fetch-app/model"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/kataras/iris/v12"
)

func JWTIsExist(ctx iris.Context) {
	headerToken := ctx.GetHeader("Authorization")

	if headerToken == " " {
		ctx.StatusCode(401)
		err := errors.New("You Don't Have Authorization To Access This Page")
		ctx.JSON(model.BadResp{
			Status:  401,
			Message: err.Error(),
		})
		ctx.EndRequest()
		return
	}
	tmpToken := strings.Split(headerToken, " ")
	// log.Println("AUTO")
	// log.Println(tmpToken)

	tokenString := tmpToken[0]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SECRET), nil
	})

	if err != nil {
		log.Println(err)
		err = errors.New("Failed to parse JWT Token")
		ctx.StatusCode(500)
		ctx.JSON(model.BadResp{
			Status:  500,
			Message: err.Error(),
		})
		return
	}

	if !token.Valid {

		ctx.StatusCode(401)
		err = errors.New("JWT Token Not Valid")
		ctx.JSON(model.BadResp{
			Status:  401,
			Message: err.Error(),
		})
		return

	}
	ctx.Next()
}

func IsAdmin(ctx iris.Context) {
	var claim authmod.Token

	headerToken := ctx.GetHeader("Authorization")

	if headerToken == " " {
		ctx.StatusCode(401)
		err := errors.New("You Don't Have Authorization To Access This Page")
		ctx.JSON(model.BadResp{
			Status:  401,
			Message: err.Error(),
		})
		ctx.EndRequest()
		return
	}
	tmpToken := strings.Split(headerToken, " ")

	tokenString := tmpToken[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SECRET), nil
	})

	if err != nil {
		log.Println(err)
		err = errors.New("Failed to parse JWT Token")
		ctx.StatusCode(500)
		ctx.JSON(model.BadResp{
			Status:  500,
			Message: err.Error(),
		})
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claim.Role = fmt.Sprintf("", claims["role"])
	} else {
		log.Println(err)
		err = errors.New("Failed to parse private claims")
		ctx.JSON(model.BadResp{
			Status:  500,
			Message: err.Error(),
		})

		return
	}

	if claim.Role != "admin" || claim.Role != "super Admin" {
		log.Println(err)
		err = errors.New("Only administrator can access this page")
		ctx.JSON(model.BadResp{
			Status:  401,
			Message: err.Error(),
		})

	}
	ctx.Next()

}
