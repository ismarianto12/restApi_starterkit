package utils

type Response struct {
	Data       interface{} `json:"data"`
	Statuscode int         `json:"code"`
	Messages   string      `json:"message"`
}

type Appresponse interface {
	successResponse()
}

func SuccessResponse(Data interface{}, Statuscode int, Messages string) *Response {

	return &Response{
		Data:       Data,
		Statuscode: Statuscode,
		Messages:   Messages,
	}
}

func ErrorResponse(Data interface{}, Statuscode int, Messages string) *Response {

	return &Response{
		Data:       nil,
		Statuscode: Statuscode,
		Messages:   Messages,
	}
}
