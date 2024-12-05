package dto

type JsonStruct struct {
	Code    int         `json:"code"`
	Message interface{} `json:"msg"`
	Data    interface{} `json:"data"`
	Count   int         `json:"count"`
}

type JsonErrorStruct struct {
	Code    int         `json:"code"`
	Message interface{} `json:"msg"`
}
