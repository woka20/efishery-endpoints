package cache

import (
	"efishery-endpoints/fetch-app/repository/rates"
	"errors"
	"log"
	"time"
)

var UpdatedRate float64

type CacheInterface interface {
	StartCache(tick *time.Ticker)
}

type CacheRepo struct {
	RateRepo rates.RatesRepoInterface
}

func NewCache() CacheInterface {
	return &CacheRepo{
		RateRepo: rates.NewRatesRepo(),
	}
}

func (c *CacheRepo) StartCache(tick *time.Ticker) {
	c.UpdatedCache()

	for {
		select {
		case <-tick.C:
			c.UpdatedCache()
		}
	}

}

func (c *CacheRepo) UpdatedCache() {
	rates, err := c.RateRepo.GetRates("IDR_USD")

	if err != nil {
		log.Println(err)
		err = errors.New("Failed To Get rates")
		return
	}

	UpdatedRate = rates
}
