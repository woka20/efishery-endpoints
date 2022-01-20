package model

type Commodity struct {
	ID        string `json:id`
	Comodity  string `json:commodity`
	Province  string `json:province`
	City      string `json:city`
	Size      string `json:size`
	IDRPrice  string `json:idr_price`
	Date      string `json:date"`
	Timestamp string `json:timestamp`
	USDPrice  string `json:price_usd`
}

type CommodityDetailResult struct {
	Province float64                   `json:province`
	Profit   map[string]map[string]int `json:profit`
	Max      float64                   `json:max`
	Min      float64                   `json:min`
	Average  float64                   `json:avg`
	Median   float64                   `json:median`
}
