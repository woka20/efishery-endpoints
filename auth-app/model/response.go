package model

type SuccessResp struct {
	Status int         `json:status`
	Data   interface{} `json:data`
}

type BadResp struct {
	Status  int    `json:status`
	Message string `json:message`
}
