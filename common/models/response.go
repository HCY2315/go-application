package models

type Response struct {
	// 状态码
	Code int `json:"code" example:"200"`
	// 数据集
	Data interface{} `json:"data"`
	// 消息
	Msg       string `json:"msg"`
	RequestId string `json:"requestId"`
}

type Page struct {
	List     interface{} `json:"list"`
	Count    int         `json:"count"`
	PageNum  int         `json:"page_num"`
	PageSize int         `json:"page_size"`
}

// ReturnError 返回错误
func (res *Response) ReturnError(code int) *Response {
	res.Code = code
	return res
}

// ReturnOK 正常返回
func (res *Response) ReturnOK() *Response {
	res.Code = 200
	return res
}
