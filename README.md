# Gin Error Handling Middleware

Gin Error Handling Middleware is a middleware for the popular [Gin framework](https://github.com/gin-gonic/gin) that enables you to configure error handling centrally as a convention within your Go application, as opposed to explicitly handling exceptions within each handler or controller action.

This gives the following benefits:

- Centralised location for handling errors
- Reduce boilerplate 'error to response' mappings in your request handlers/controller actions
- Helps protect yourself from inadvertently revealing errors to API consumers 

## Quick Start

### Simple error handling configuration

```go
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
```

Returns the following HTTP response:

```
HTTP/1.1 404 Not Found
Date: Wed, 16 Feb 2022 04:08:50 GMT
Content-Length: 0
Connection: close
```

### More control over the response

```go
var (
    NotFoundError = fmt.Errorf("resource could not be found")
)

func main() {
    r := gin.Default()
    r.Use(
        ErrorHandler(
            Map(NotFoundError).ToResponse(func(c *gin.Context, err error) {
                c.Status(http.StatusNotFound)
                c.Writer.Write([]byte(err.Error())),
            }),
        ))

    r.GET("/ping", func(c *gin.Context) {
        c.Error(NotFoundError)
    })

    r.Run()
}
```

Returns the following response:

```
HTTP/1.1 404 Not Found
Date: Wed, 16 Feb 2022 04:21:37 GMT
Content-Length: 27
Content-Type: text/plain; charset=utf-8
Connection: close

resource could not be found
```