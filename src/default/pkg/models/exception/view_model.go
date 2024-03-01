package exception

import (
	"errors"
	"net/http"
)

type ErrorView struct {
	ResponseCode int
	ResponseType string
	ErrorMessage string
}

func NewErrorView(err error) *ErrorView {
	var notFoundErr *NotFoundErr
	var badRequestErr *BadRequestErr
	if errors.As(err, &notFoundErr) {
		return &ErrorView{
			ResponseCode: http.StatusNotFound,
			ResponseType: "Not Found",
			ErrorMessage: notFoundErr.Error(),
		}
	} else if errors.As(err, &badRequestErr) {
		return &ErrorView{
			ResponseCode: http.StatusBadRequest,
			ResponseType: "Bad Request",
			ErrorMessage: badRequestErr.Error(),
		}
	}

	return &ErrorView{
		ResponseCode: http.StatusInternalServerError,
		ResponseType: "Internal Server Error",
		ErrorMessage: "An internal server error has occurred. This should not happen in a normal process flow. Please send us a support case with any sanitized inputs and the date and time of the failed API call.",
	}
}
