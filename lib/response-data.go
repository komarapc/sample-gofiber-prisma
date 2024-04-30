package lib

import "strings"

type ResponseProps struct {
	Code    int
	Message interface{}
	Data    interface{}
}
type ResponseData struct {
	StatusCode    int         `json:"status_code"`
	StatusMessage string      `json:"status_message"`
	Message       interface{} `json:"message,omitempty"`
	Success       bool        `json:"success"`
	Data          interface{} `json:"data,omitempty"`
}

var HttpStatusCode = map[int]string{
	200: "OK",
	201: "CREATED",
	400: "BAD_REQUEST",
	401: "UNAUTHORIZED",
	402: "NEED_PAYMENT",
	403: "FORBIDDEN",
	404: "NOT_FOUND",
	409: "CONFLICT",
	422: "UNPROCESSABLE_ENTITY",
	429: "TOO_MANY_REQUEST",
	500: "INTERNAL_SERVER_ERROR",
}

func getStatusCode(code int) string {
	return HttpStatusCode[code]
}

func getStatusMessage(code int) string {
	return strings.ReplaceAll(HttpStatusCode[code], "_", " ")
}

func ResponseSuccess(props ResponseProps) ResponseData {
	return ResponseData{
		StatusCode:    props.Code,
		StatusMessage: getStatusCode(props.Code),
		Message:       props.Message,
		Success:       true,
		Data:          props.Data,
	}
}

func ResponseError(props ResponseProps) ResponseData {
	return ResponseData{
		StatusCode:    props.Code,
		StatusMessage: getStatusCode(props.Code),
		Message:       props.Message,
		Success:       false,
		Data:          props.Data,
	}
}
