package tokenization

import (
	"efishery-endpoints/fetch-app/logic/tokenization"
	"efishery-endpoints/fetch-app/model"
	"strings"

	"github.com/kataras/iris/v12"
)

type TokenHandlerInterface interface {
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

func (h *TokenHandler) GetClaims(ctx iris.Context) {
	bearerToken := ctx.GetHeader("Authorization")
	tempString := strings.Split(bearerToken, " ")

	tokenString := tempString[1]
	claim, err := h.TokenLogic.ParseToken(tokenString)
	if err != nil {
		ctx.JSON(model.BadResp{
			Status:  500,
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(model.SuccessResp{
		Status: 200,
		Data:   claim,
	})

}
