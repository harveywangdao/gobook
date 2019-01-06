package models

type ApiResult struct {
	Status bool        `json:"status" description:"status"`
	Msg    string      `json:"msg" description:"error msg"`
	Data   interface{} `json:"data" description:"data"`
}
