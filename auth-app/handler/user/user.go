package user

import (
	"efishery-endpoints/auth-app/logic/user"
	"efishery-endpoints/auth-app/model"
	"encoding/json"
	"log"

	"github.com/kataras/iris/v12"
)

type UserHandlerInterface interface {
	Hello(ctx iris.Context)
	AddNewUser(ctx iris.Context)
}

type UserHandler struct {
	UserLogic user.UserLogicInterface
}

func NewUserHandler() UserHandlerInterface {
	return &UserHandler{
		UserLogic: user.NewLogicUser(),
	}
}

func (h *UserHandler) Hello(ctx iris.Context) {
	ctx.JSON("Hello World!")
}

func (h *UserHandler) AddNewUser(ctx iris.Context) {
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

	password, err := h.UserLogic.ProducePassword()

	if err != nil {
		ctx.StatusCode(500)
		ctx.JSON(model.BadResp{
			Status:  500,
			Message: err.Error(),
		})
		return
	}
	user.Password = password

	ctx.StatusCode(200)
	ctx.JSON(model.SuccessResp{
		Status: 200,
		Data:   user.Password,
	})

}
