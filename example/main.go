package main

import (
	"fmt"
	. "github.com/gin-errorhandling/middleware"
	"github.com/gin-gonic/gin"

	"net/http"
)

var (
	NotFoundError = fmt.Errorf("resource could not be found")
)

func main() {
	r := gin.Default()
	r.Use(
		ErrorHandler(
			Map(NotFoundError).ToStatusCode(http.StatusNotFound),
		))

	r.GET("/ping", func(c *gin.Context) {
		c.Error(NotFoundError)
	})

	r.Run()
}
