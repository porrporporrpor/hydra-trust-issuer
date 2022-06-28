package appresponse

import (
	"fmt"
	"reflect"

	"github.com/porrporporrpor/hydra-trust-issuer/config"
	"github.com/porrporporrpor/hydra-trust-issuer/pkg/apperror"
)

var (
	codeResponsePattern = "%s-%d"
)

type ResponseSuccess struct {
	Code    string      `json:"code" example:"TRUST-2001"`
	Message string      `json:"message" example:"Success"`
	Data    interface{} `json:"data"`
}

type ResponseError struct {
	Code    string      `json:"code" example:"TRUST-5001"`
	Message string      `json:"message" example:"Internal Server Error"`
	Errors  interface{} `json:"errors"`
}

func Success(data interface{}) ResponseSuccess {
	if data == nil {
		type Empty struct{}
		data = Empty{}
	}

	return ResponseSuccess{
		Code:    fmt.Sprintf(codeResponsePattern, config.GetConfig().App.Prefix, OKSuccessCode),
		Message: SuccessMessage,
		Data:    data,
	}
}

func Error(code apperror.ErrorCode, message string, errs interface{}) ResponseError {
	var (
		response = ResponseError{
			Code:    fmt.Sprintf(codeResponsePattern, config.GetConfig().App.Prefix, code),
			Message: message,
			Errors:  []interface{}{},
		}
	)

	if errs != nil {
		response.Errors = []interface{}{errs}
		if reflect.TypeOf(errs) == reflect.TypeOf([]apperror.ValidateError{}) {
			response.Errors = errs
		}
	}

	return response

}
