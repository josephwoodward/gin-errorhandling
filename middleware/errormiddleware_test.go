package middleware_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/gin-errorhandling/middleware"
	"github.com/gin-gonic/gin"
)

var (
	NotFoundError = fmt.Errorf("this is an error")
)

type ValidationError struct {
	customError string
}

func (e *ValidationError) Error() string {
	return "Invalid request"
}

func TestMapSimpleErrorToStatusCode(t *testing.T) {
	// Arrange
	router := gin.Default()
	router.Use(
		ErrorHandler(
			Map(NotFoundError).ToStatusCode(http.StatusNotFound),
		))

	// Act
	router.GET("/", func(context *gin.Context) {
		_ = context.Error(NotFoundError)
	})

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, httptest.NewRequest("GET", "/", nil))

	// Assert
	assert.Equal(t, recorder.Result().StatusCode, http.StatusNotFound)
}

func TestMapErrorStructToStatusCode(t *testing.T) {
	// Arrange
	router := gin.Default()
	router.Use(
		ErrorHandler(
			Map(&ValidationError{}).ToStatusCode(http.StatusBadRequest),
		))

	// Act
	router.GET("/", func(context *gin.Context) {
		_ = context.Error(&ValidationError{})
	})

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, httptest.NewRequest("GET", "/", nil))

	// Assert
	assert.Equal(t, recorder.Result().StatusCode, http.StatusBadRequest)
}

func TestMapErrorResponseFunc(t *testing.T) {
	// Arrange
	router := gin.Default()
	router.Use(
		ErrorHandler(
			Map(NotFoundError).ToResponse(func(ctx *gin.Context, err error) {
				ctx.Status(http.StatusNotFound)
				ctx.Writer.Write([]byte(err.Error()))
			}),
		))

	// Act
	router.GET("/", func(context *gin.Context) {
		_ = context.Error(NotFoundError)
	})

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, httptest.NewRequest("GET", "/", nil))

	// Assert
	assert.Equal(t, http.StatusNotFound, recorder.Result().StatusCode)
	assert.Equal(t, NotFoundError.Error(), recorder.Body.String())
}
