package webservice

import (
	"bytes"
	"strconv"
	"net/http"
)

//var InvalidParam = FlyError{101, "Invalid Params"}
//var ExceptionError = FlyError{100, "Exception found"}

type WsError struct {
	Code    int    `json:"errorCode"`
	Message string `json:"errorMessage"`
}

func (we *WsError) Error() string {
	buf := bytes.Buffer{}
	buf.WriteString(strconv.Itoa(we.Code))
	buf.WriteString(" : ")
	buf.WriteString(we.Message)
	return buf.String()
}

type ResError struct {
	WsError
	StatusCode int
}

func (re *ResError) Error() string {
	buf := bytes.Buffer{}
	buf.WriteString("Code : ")
	buf.WriteString(strconv.Itoa(re.StatusCode))
	buf.WriteByte('(')
	buf.WriteString(strconv.Itoa(re.Code))
	buf.WriteString(") Message : ")
	buf.WriteString(re.Message)
	return buf.String()
}

func (re ResError) NewMessage(message string) ResError {
	re.Message = message
	return re
}

func CreateResponseError(code int, message string, statusCode int) ResError {
	return ResError{WsError{code, message}, statusCode}
}

type wsError struct {
	InvalidParams       ResError
	UnAuthorized        ResError
	SqlError            ResError
	InternalServerError ResError
}

var WsErrors *wsError

func init() {
	WsErrors = &wsError{
		InvalidParams:       CreateResponseError(10000, "Invalid Params", http.StatusNotAcceptable),         //{10001, "Invalid Params"},
		UnAuthorized:        CreateResponseError(10001, "Unauthoried User", http.StatusUnauthorized),        //flyError.FlyError{10000, "Unauthoried User"},
		SqlError:            CreateResponseError(10002, "Sql Excute Error", http.StatusInternalServerError), //flyError.FlyError{10002, "Sql excute error"},
		InternalServerError: CreateResponseError(10003, "Internal Server Error", http.StatusInternalServerError),
	}
}
