package rest_err

import "net/http"

type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int    `json:"code"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Algumas libs retornam/recebem erros por parametros. Criei uma interface para quando precisar mandar um erro por parametro, mandar um RestErr
func (r *RestErr) Error() string {
	return r.Message
}

func NewRestErr(message, err string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}

func NewForbidden(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}
