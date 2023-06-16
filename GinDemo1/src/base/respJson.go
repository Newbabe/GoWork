package base

import "encoding/json"

type resp struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
	Count int         `json:"count"`
}

func RespJson(code int, msg string, data interface{}, count int) string {
	j := &resp{
		Code:  code,
		Msg:   msg,
		Data:  data,
		Count: count,
	}
	marshal, err := json.Marshal(j)
	if err != nil {
		return ""
	}
	return string(marshal)
}
