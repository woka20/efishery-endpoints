package model

type SuccessResp struct {
	status int         `json:status`
	Data   interface{} `json:data`
}

type BadResp struct {
	status  int    `json:status`
	message string `json:message`
}
