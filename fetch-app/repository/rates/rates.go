package rates

import (
	"efishery-endpoints/fetch-app/config"
	"efishery-endpoints/fetch-app/model"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type RatesRepoInterface interface {
	GetRates(rates string) (float64, error)
}

type RatesRepo struct {
}

func NewRatesRepo() RatesRepoInterface {
	return &RatesRepo{}

}

func (r *RatesRepo) GetRates(rates string) (float64, error) {

	var NewRate model.Rates

	log.Println("Start Converting Currency Rates")
	url := fmt.Sprintf("https://free.currconv.com/api/v7/convert?q=%s&compact=ultra&apiKey=%s", rates, config.KEY_API)

	response, err := http.Get(url)

	if err != nil {
		log.Println(err)
		errors.New("Failed To Get Rates From Third Parties")
		return 0.0, err
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	sb := string(content)

	err = json.Unmarshal([]byte(sb), &NewRate)

	// err = json.NewDecoder(response.Body).Decode(&NewRate)

	// log.Println(NewRate)
	if err != nil {
		log.Println(err)
		err = errors.New("Failed to unmarshal conversion rate")
		return 0, err
	} else if NewRate.IDR_USD == 0 {
		log.Println(err)
		err = errors.New("Conversion Rate is 0")
		return 0, err
	}

	return NewRate.IDR_USD, nil

}
