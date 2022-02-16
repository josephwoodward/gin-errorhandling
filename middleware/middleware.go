package middleware

import (
	"github.com/gin-gonic/gin"
	"reflect"
)

func ErrorHandler(errMap ...*errorMapping) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()

		lastErr := context.Errors.Last()
		if lastErr == nil {
			return
		}

		for _, e := range errMap {
			for _, e2 := range e.fromErrors {
				if lastErr.Err == e2 {
					e.toResponse(context, lastErr.Err)
					//context.Status(e.toStatusCode)
				} else if isType(lastErr.Err, e2) {
					//context.Status(e.toStatusCode)
					e.toResponse(context, lastErr.Err)
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

// errorMapping maps a single set of errors to a single response
type errorMapping struct {
	fromErrors   []error
	toStatusCode int

	toResponse func(ctx *gin.Context, err error)
}

func (r *errorMapping) ToStatusCode(statusCode int) *errorMapping {
	r.toStatusCode = statusCode
	r.toResponse = func(ctx *gin.Context, err error) {
		ctx.Status(statusCode)
	}
	return r
}

func (r *errorMapping) ToResponse(response func(ctx *gin.Context, err error)) *errorMapping {
	r.toResponse = response
	return r
}

func Map(err ...error) *errorMapping {
	return &errorMapping{
		fromErrors: err,
	}
}

func MapFunc(errFunc func() error) *errorMapping {
	return &errorMapping{}
}
