package model

type Commodity struct {
	ID        string `json:uuid`
	Comodity  string `json:komoditas`
	Province  string `json:area_provinsi`
	City      string `json:area_kota`
	Size      string `json:size`
	Price     string `json:price`
	Date      string `json:tgl_parsed`
	Timestamp string `json:timestamp`
	USDPrice  string `json:usd_price`
}

type CommodityDetailResult struct {
	Province string                    `json:area_provinsi`
	Profit   map[string]map[string]int `json:profit`
	Max      float64                   `json:max`
	Min      float64                   `json:min`
	Average  float64                   `json:avg`
	Median   float64                   `json:median`
}
