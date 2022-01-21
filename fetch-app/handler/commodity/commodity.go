package commodity

import (
	"efishery-endpoints/auth-app/model"
	"efishery-endpoints/fetch-app/logic/commodity"

	"github.com/kataras/iris/v12"
)

type CommodityHandlerInterface interface {
	GetList(ctx iris.Context)
	AggregateData(ctx iris.Context)
}

type CommodityHandler struct {
	CommodityLogic commodity.CommodityLogicInterface
}

func NewComodityHandler() CommodityHandlerInterface {
	return &CommodityHandler{
		CommodityLogic: commodity.NewCommodityLogic(),
	}
}

func (h *CommodityHandler) GetList(ctx iris.Context) {
	listCommodity, err := h.CommodityLogic.GetList()
	if err != nil {
		ctx.JSON(model.BadResp{
			Status:  500,
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(model.SuccessResp{
		Status: 200,
		Data:   listCommodity,
	})
}

func (h *CommodityHandler) AggregateData(ctx iris.Context) {
	listCommodity, err := h.CommodityLogic.AggregateData()
	if err != nil {
		ctx.JSON(model.BadResp{
			Status:  500,
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(model.SuccessResp{
		Status: 200,
		Data:   listCommodity,
	})
}
