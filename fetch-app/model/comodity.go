package model

type Comodity struct {
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

type ComodityDetailResult struct {
	Province string `json:province`
	Profit   string `json:profit`
	Max      string `json:max`
	Min      string `json:min`
	Average  string `json:avg`
	Median   string `json:median`
}
