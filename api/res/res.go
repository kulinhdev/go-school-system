package res

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

const (
	StatusSuccess = 1
	StatusError   = 0
)

type ApiError struct {
	Field string
	Msg   string
}

// ErrorResponse represents the structure of an error response.
type ErrorResponse struct {
	Status  int         `json:"status"`
	Code    string      `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data"`
}

// SuccessResponse represents the structure of a success response.
type SuccessResponse struct {
	Status int         `json:"status"`
	Code   string      `json:"code"`
	Data   interface{} `json:"data"`
}

// ResponseError sends a generic error response.
func ResponseError(c *gin.Context, code int, message string, payload interface{}) {
	c.JSON(code, ErrorResponse{
		Status:  StatusError,
		Code:    strconv.Itoa(code),
		Message: message,
		Data:    payload,
	})
}

// ResponseSuccess sends a success response with a specific code and payload.
func ResponseSuccess(c *gin.Context, code int, payload interface{}) {
	c.JSON(code, SuccessResponse{
		Status: StatusSuccess,
		Code:   strconv.Itoa(code),
		Data:   payload,
	})
}

// ResponseFake sends a fake response indicating a temporarily suspended feature.
func ResponseFake(c *gin.Context) {
	c.JSON(200, ErrorResponse{
		Status: StatusError,
		Code:   "Test fake response api...",
	})
}

// GetIntQuery retrieves an integer query parameter from the request.
func GetIntQuery(c *gin.Context, key string) int {
	s := c.Query(key)
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return v
}

// GetInt64Query retrieves an int64 query parameter from the request.
func GetInt64Query(c *gin.Context, key string) int64 {
	s := c.Query(key)
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return v
}

// GetBoolQuery retrieves a boolean query parameter from the request.
func GetBoolQuery(c *gin.Context, key string) bool {
	s := c.Query(key)
	return s == "1" || s == "true"
}
