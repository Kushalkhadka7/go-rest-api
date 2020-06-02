package error

import (
	"net/http"
)

// Response for custom errors in the app.
type Response struct {
	Message        string `json:"message"`
	HTTPStatusCode int    `json:"statusCode"`
	StatusText     string `json:"status"`
}

func (r *Response) Error() string { return r.Message }

// NotFound error is implemented when desired result is not found.
func NotFound(err error) *Response {
	return &Response{
		Message:        err.Error(),
		HTTPStatusCode: http.StatusNotFound,
		StatusText:     http.StatusText(http.StatusNotFound),
	}
}

// CustomError is generic error format.
func CustomError(err string, code int) *Response {
	return &Response{
		Message:        err,
		HTTPStatusCode: http.StatusNotFound,
		StatusText:     http.StatusText(http.StatusNotFound),
	}
}
