package errors

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiError interface {
	Message() string
	Code() string
	Status() int
	Cause() []interface{}
	Error() string
}

type apiError struct {
	ErrorMessage string        `json:"message"`
	ErrorCode    string        `json:"error"`
	ErrorStatus  int           `json:"status"`
	ErrorCause   []interface{} `json:"cause"`
}

func NewApiError(status int, code, message string) ApiError {
	return apiError{
		ErrorMessage: message,
		ErrorCode:    code,
		ErrorStatus:  status,
		ErrorCause:   []interface{}{},
	}
}

func (e apiError) Code() string {
	return e.ErrorCode
}

func (e apiError) Error() string {
	return fmt.Sprintf("Message: %s;Error Code: %s;Status: %d;Cause: %v", e.ErrorMessage, e.ErrorCode, e.ErrorStatus, e.ErrorCause)
}

func (e apiError) Cause() []interface{} {
	return e.ErrorCause
}

func (e apiError) Status() int {
	return e.ErrorStatus
}

func (e apiError) Message() string {
	return e.ErrorMessage
}

func SendErrorRespose(c *gin.Context, err error) {
	apiError, ok := err.(ApiError)
	if !ok {
		apiError = NewApiError(http.StatusInternalServerError, UnknownError.Code(), err.Error())
	}
	c.AbortWithStatusJSON(apiError.Status(), err)
}
