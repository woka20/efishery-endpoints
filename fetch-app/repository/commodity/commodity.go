package commodity

import (
	"efishery-endpoints/fetch-app/model"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type CommodityRepoInterface interface {
	GetCommodityList() (commod []model.Commodity, err error)
}

type CommodityRepo struct {
}

func NewCommodityRepo() CommodityRepoInterface {
	return &CommodityRepo{}

}

func (c *CommodityRepo) GetCommodityList() (commod []model.Commodity, err error) {

	response, err := http.Get("https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list")
	if err != nil {
		log.Println(err)
		errors.New("Failed To Get List Commodities")
		return commod, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&commod)

	if err != nil {
		log.Println(err)
		errors.New("Failed To Decode Get List Commodities")
		return commod, err

	}

	return commod, nil

}
