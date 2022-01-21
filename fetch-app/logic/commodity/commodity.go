package commodity

import (
	"efishery-endpoints/fetch-app/cache"
	"efishery-endpoints/fetch-app/model"
	"efishery-endpoints/fetch-app/repository/commodity"
	"efishery-endpoints/fetch-app/repository/rates"
	"efishery-endpoints/fetch-app/utils"
	"errors"
	"fmt"
	"log"
	"sort"
	"strconv"
)

type temporaryData struct {
	Provinsi string
	Tahun    string
	Minggu   string
	Amount   int
}

type CommodityLogicInterface interface {
	GetList() (listCommodity []model.Commodity, err error)
}

type CommodityLogic struct {
	CommodityRepo commodity.CommodityRepoInterface
	RateRepo      rates.RatesRepoInterface
}

func NewCommodityLogic() CommodityLogicInterface {
	return &CommodityLogic{
		CommodityRepo: commodity.NewCommodityRepo(),
		RateRepo:      rates.NewRatesRepo(),
	}
}

func (c *CommodityLogic) GetList() (listCommodity []model.Commodity, err error) {
	commod, err := c.CommodityRepo.GetCommodityList()

	if err != nil {
		log.Println(err)
		err := errors.New("Failed to get List")
		return listCommodity, err
	}

	listCommodity, err = c.USPriceAdd(commod)
	if err != nil {
		log.Println(err)
		err := errors.New("Failed to Add US Price")
		return listCommodity, err
	}

	return listCommodity, nil

}

func (c *CommodityLogic) USPriceAdd(listcommod []model.Commodity) (listresult []model.Commodity, err error) {

	var NewRate float64

	if cache.UpdatedRate != 0 {
		NewRate = cache.UpdatedRate

	} else {
		convRate, err := c.RateRepo.GetRates("IDR_USD")
		if err != nil {
			log.Println(err)
			err := errors.New("Failed to get rates idr to usd")
			return listresult, err

		}

		NewRate = convRate
	}
	for _, value := range listcommod {
		if value.Price == "" {
			continue
		}
		pricefloat64, err := strconv.ParseFloat(value.Price, 64)
		if err != nil {
			log.Println(err, "Failed to convert string price to float")
		}

		dollarField := pricefloat64 * NewRate

		value.USDPrice = fmt.Sprint("", dollarField)

		listresult = append(listresult, value)

	}
	return listresult, nil
}

func (c *CommodityLogic) AggregateData() (listCommodity []model.CommodityDetailResult, err error) {
	commod, err := c.CommodityRepo.GetCommodityList()

	if err != nil {
		log.Println(err)
		err := errors.New("Failed to get List")
		return listCommodity, err
	}

	var datas []temporaryData

	for _, value := range commod {
		if value.Date == "" || value.Price == "" || value.Size == "" || value.Province == "" {
			continue
		}

		dateTimes := utils.StringToTime(value.Date)
		year, week := dateTimes.ISOWeek()

		size, _ := strconv.Atoi(value.Size)
		price, _ := strconv.Atoi(value.Price)

		amount := size * price

		temp := temporaryData{
			Provinsi: value.Province,
			Amount:   amount,
			Tahun:    fmt.Sprintf("", year),
			Minggu:   fmt.Sprintf("", week),
		}

		datas = append(datas, temp)

	}

	temporaryMap := make(map[string]map[string]map[string]int)

	for _, value := range datas {
		if province, ok := temporaryMap[value.Provinsi]; !ok {
			minggu := map[string]int{value.Minggu: value.Amount}
			tahun := map[string]map[string]int{value.Tahun: minggu}
			temporaryMap[value.Provinsi] = tahun
		} else {

			if year, ok := province[value.Tahun]; !ok {
				minggu := map[string]int{value.Minggu: value.Amount}
				province[value.Tahun] = minggu

			} else {
				if amount, ok := year[value.Minggu]; !ok {
					year[value.Minggu] = value.Amount

				} else {
					year[value.Minggu] = year[value.Minggu] + amount
				}
			}
		}
	}
	var compiledData []model.CommodityDetailResult

	for key, value := range temporaryMap {
		newData := model.CommodityDetailResult{
			Province: key,
			Profit:   value,
			Min:      c.FindMin(value),
			Max:      c.FindMax(value),
			Average:  c.FindAvg(value),
			Median:   c.FindMedian(value),
		}
		compiledData = append(compiledData, newData)
	}

	return compiledData, nil

}

func (c *CommodityLogic) FindMax(profits map[string]map[string]int) float64 {
	var max int

	for _, value := range profits {
		for _, number := range value {
			if number >= max {
				max = number
			}
		}

	}

	return float64(max)

}

func (c *CommodityLogic) FindMin(profits map[string]map[string]int) float64 {
	min := 0
	for _, value := range profits {
		for _, number := range value {
			if number <= min {
				min = number
			}
		}

	}

	return float64(min)
}

func (c *CommodityLogic) FindAvg(profits map[string]map[string]int) float64 {
	tot := 0
	count := 0

	for _, value := range profits {
		for _, number := range value {

			tot = tot + number
			count = count + 1
		}

	}

	return float64(tot / count)

}

func (c *CommodityLogic) FindMedian(profits map[string]map[string]int) float64 {
	var medArray []int
	for _, value := range profits {
		for _, number := range value {
			medArray = append(medArray, number)

		}

	}
	sort.Ints(medArray)
	counter := len(medArray)

	if counter%2 == 0 {
		x := medArray[(counter/2)-1]
		y := medArray[(counter / 2)]
		return float64((x + y) / 2)

	} else {

		return float64(medArray[((counter+1)/2)-1])
	}

}
