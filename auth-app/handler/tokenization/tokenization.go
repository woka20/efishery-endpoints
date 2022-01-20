package tokenization

import (
	"efishery-endpoints/auth-app/logic/tokenization"
	"efishery-endpoints/auth-app/model"
	"encoding/json"
	"log"
	"strings"

	"github.com/kataras/iris/v12"
)

type TokenHandlerInterface interface {
	GetToken(ctx iris.Context)
	GetClaims(ctx iris.Context)
}

type TokenHandler struct {
	TokenLogic tokenization.TokenLogiCInterface
}

func NewTokenHandler() TokenHandlerInterface {
	return &TokenHandler{
		TokenLogic: tokenization.NewTokenLogic(),
	}
}

func (t *TokenHandler) GetToken(ctx iris.Context) {
	body, err := ctx.GetBody()

	if err != nil {
		log.Println(err)
		ctx.StatusCode(500)
		ctx.JSON(model.BadResp{
			Status:  500,
			Message: "Error reading request body information",
		})
		return
	}

	var user model.User

	if err := json.Unmarshal(body, &user); err != nil {
		log.Println(err)
		ctx.StatusCode(500)
		ctx.JSON(model.BadResp{
			Status:  500,
			Message: "Error Umarshal request body information",
		})
		return
	}

	tkn, err := t.TokenLogic.ProduceToken(user)

	if err != nil {
		ctx.StatusCode(500)
		ctx.JSON(model.BadResp{
			Status:  500,
			Message: err.Error(),
		})
		return
	}

	ctx.StatusCode(200)
	ctx.JSON(model.SuccessResp{
		Status: 200,
		Data:   tkn,
	})

}

func (t *TokenHandler) GetClaims(ctx iris.Context) {
	headerToken := ctx.GetHeader("Authorization")
	tmpToken := strings.Split(headerToken, " ")

	tokenString := tmpToken[1]

	tokenClaims, err := t.TokenLogic.ParseToken(tokenString)

	if err != nil {
		log.Println(err)
		ctx.StatusCode(500)
		ctx.JSON(model.BadResp{
			Status:  500,
			Message: err.Error(),
		})
		return
	}

	ctx.StatusCode(200)
	ctx.JSON(model.SuccessResp{
		Status: 200,
		Data:   tokenClaims,
	})

}
