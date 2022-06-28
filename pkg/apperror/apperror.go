package apperror

import "errors"

// AppError ... an error adaptor to adapting error from each layer into application error
// to make it easier to handle
type AppError struct {
	Code    ErrorCode
	Message ErrorMessage
	Errors  interface{}
}

type SingleErrorMessage struct {
	Message string `json:"message,omitempty"`
}

type ValidateError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e AppError) Error() string {
	return string(e.Message)
}

// IsAppError ... check an error, is that error made by AppError?
func IsAppError(err error) (e AppError, ok bool) {
	ok = errors.As(err, &e)
	return
}

// NewBadRequestError ... Missing mandatory parameter or Invalid data type
func NewBadRequestError(errs interface{}) error {
	return AppError{
		Code:    BadRequestError,
		Message: Message[BadRequestError],
		Errors:  errs,
	}
}

// NewUnAuthorizedError ... Failed from authorization provider
func NewUnAuthorizedError(msg string) error {
	var errMsg interface{}
	if msg != "" {
		errMsg = SingleErrorMessage{Message: msg}
	}

	return AppError{
		Code:    UnAuthorizedError,
		Message: Message[UnAuthorizedError],
		Errors:  errMsg,
	}
}

// NewForbiddenError ... Failed from access to forbidden resources
func NewForbiddenError(msg string) error {
	var errMsg interface{}
	if msg != "" {
		errMsg = SingleErrorMessage{Message: msg}
	}

	return AppError{
		Code:    ForbiddenError,
		Message: Message[ForbiddenError],
		Errors:  errMsg,
	}
}

// NewNotFoundError ... Failed from access to unknown resources
func NewNotFoundError(msg string) error {
	var errMsg interface{}
	if msg != "" {
		errMsg = SingleErrorMessage{Message: msg}
	}
	return AppError{
		Code:    NotFoundError,
		Message: Message[NotFoundError],
		Errors:  errMsg,
	}
}

func NewPageNotFoundError() error {
	return AppError{
		Code:    PageNotFoundError,
		Message: Message[PageNotFoundError],
		Errors:  nil,
	}
}

// NewConflictError ... Failed when interacting with database
func NewConflictError(msg string) error {
	var errMsg interface{}
	if msg != "" {
		errMsg = SingleErrorMessage{Message: msg}
	}

	return AppError{
		Code:    ConflictError,
		Message: Message[ConflictError],
		Errors:  errMsg,
	}
}

// NewGoneError ... use when got request was used
func NewGoneError() error {
	return AppError{
		Code:    GoneError,
		Message: Message[GoneError],
		Errors:  nil,
	}
}

// NewInternalServerError ... use when got an internal server error
func NewInternalServerError(msg string) error {
	var errMsg interface{}
	if msg != "" {
		errMsg = SingleErrorMessage{Message: msg}
	}

	return AppError{
		Code:    InternalServerError,
		Message: Message[InternalServerError],
		Errors:  errMsg,
	}
}

// NewExternalCallError ... use when got a http status code not return 20x or 30x
func NewExternalCallError(msg string) error {
	var errMsg interface{}
	if msg != "" {
		errMsg = SingleErrorMessage{Message: msg}
	}

	return AppError{
		Code:    ExternalCallError,
		Message: Message[ExternalCallError],
		Errors:  errMsg,
	}
}
