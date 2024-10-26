package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuccessBody struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
}

type FailureBody struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

// helper functions to get the basic success and failure responses skeleton
func getFailureResponse(c *gin.Context, status int, error string, message string) {
	c.JSON(status, FailureBody{
		Success: false,
		Error:   error,
		Message: message,
	})
	return
}

func getSuccessResponse(c *gin.Context, status int, data gin.H) {
	c.JSON(status, SuccessBody{
		Success: true,
		Data:    data,
	})
}

// functions returning responses based on various HTTP status codes
// success
func Ok(c *gin.Context, data gin.H) {
	getSuccessResponse(c, http.StatusOK, data)
	return
}

func Created(c *gin.Context, data gin.H) {
	getSuccessResponse(c, http.StatusCreated, data)
	return
}

func NoContent(c *gin.Context, data gin.H) {
	getSuccessResponse(c, http.StatusNoContent, data)
	return
}

func Processing(c *gin.Context, data gin.H) {
	getSuccessResponse(c, http.StatusProcessing, data)
	return
}

// redirection
func MovedPermanently(c *gin.Context, data gin.H) {
	getSuccessResponse(c, http.StatusMovedPermanently, data)
	return
}

func Found(c *gin.Context, data gin.H) {
	getSuccessResponse(c, http.StatusFound, data)
	return
}

// client errors
func BadRequest(c *gin.Context, error string, message string) {
	getFailureResponse(c, http.StatusBadRequest, error, message)
	return
}

func NotFound(c *gin.Context, error string, message string) {
	getFailureResponse(c, http.StatusNotFound, error, message)
	return
}

func RequestTimeout(c *gin.Context, error string, message string) {
	getFailureResponse(c, http.StatusRequestTimeout, error, message)
	return
}

func TooManyRequests(c *gin.Context, error string, message string) {
	getFailureResponse(c, http.StatusTooManyRequests, error, message)
	return
}

func UnavailableForLegalReasons(c *gin.Context, error string, message string) {
	getFailureResponse(c, http.StatusUnavailableForLegalReasons, error, message)
	return
}

func MethodNotAllowed(c *gin.Context, error string, message string) {
	getFailureResponse(c, http.StatusMethodNotAllowed, error, message)
	return
}

func Unauthorized(c *gin.Context, error string, message string) {
	getFailureResponse(c, http.StatusUnauthorized, error, message)
	return
}

func Forbidden(c *gin.Context, error string, message string) {
	getFailureResponse(c, http.StatusForbidden, error, message)
	return
}

// server errors
func InternalServerError(c *gin.Context, error string, message string) {
	getFailureResponse(c, http.StatusInternalServerError, error, message)
	return
}

func NotImplemented(c *gin.Context, error string, message string) {
	getFailureResponse(c, http.StatusNotImplemented, error, message)
	return
}

func BadGateway(c *gin.Context, error string, message string) {
	getFailureResponse(c, http.StatusBadGateway, error, message)
	return
}
