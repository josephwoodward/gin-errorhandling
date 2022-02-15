package errormiddleware

import (
	"github.com/gin-gonic/gin"
	"reflect"
)

func ErrorHandler(errMap ...*errorMap) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()

		lastErr := context.Errors.Last()
		if lastErr == nil {
			return
		}

		for _, e := range errMap {
			for _, e2 := range e.fromErrors {
				if lastErr.Err == e2 {
					context.Status(e.toStatusCode)
				} else if isType(lastErr.Err, e2) {
					context.Status(e.toStatusCode)
				}
			}
		}
	}
}

func isType(a, b interface{}) bool {
	return reflect.TypeOf(a) == reflect.TypeOf(b)
}

type mapperConfig struct {
}

type errorMap struct {
	fromErrors   []error
	toStatusCode int
}

func (r *errorMap) ToStatusCode(statusCode int) *errorMap {
	r.toStatusCode = statusCode
	return r
}

func Map(err ...error) *errorMap {
	return &errorMap{
		fromErrors: err,
	}
}

func MapFunc(errFunc func() error) *errorMap {
	return &errorMap{}
}
