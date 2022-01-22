package model

type Commodity struct {
	UUID          string `json:uuid`
	Komoditas     string `json:komoditas`
	Area_Provinsi string `json:area_provinsi`
	Area_Kota     string `json:area_kota`
	Size          string `json:size`
	Price         string `json:price`
	Tgl_parsed    string `json:tgl_parsed`
	Timestamp     string `json:timestamp`
	USD_Price     string `json:usd_price`
}

type CommodityDetailResult struct {
	Province string                    `json:area_provinsi`
	Profit   map[string]map[string]int `json:profit`
	Max      float64                   `json:max`
	Min      float64                   `json:min`
	Average  float64                   `json:avg`
	Median   float64                   `json:median`
}
