package ecode

var OK = &Response{
	Code:    200,
	Message: "成功",
	Data:    nil,
}

var ErrInvalidParam = &Response{
	Code:    400,
	Message: "参数不合法",
	Data:    nil,
}

var ErrSystemError = &Response{
	Code:    500,
	Message: "系统错误",
	Data:    nil,
}

type PageData struct {
	Total     int
	TotalPage int
	Data      interface{}
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (r *Response) WithData(data interface{}) *Response {
	tmp := &Response{
		Code:    r.Code,
		Message: r.Message,
		Data:    data,
	}

	return tmp
}

func (r *Response) WithPageData(data interface{}, count int, pageSize int) *Response {
	if pageSize <= 0 {
		pageSize = 1
	}

	tmp := &Response{
		Code:    r.Code,
		Message: r.Message,
		Data: &PageData{
			Total:     count,
			TotalPage: count/pageSize + 1,
			Data:      data,
		},
	}

	return tmp
}

func (r *Response) WithError(err error) *Response {
	tmp := &Response{
		Code:    r.Code,
		Message: r.Message + ":" + err.Error(),
		Data:    nil,
	}

	return tmp
}
