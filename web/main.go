package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http/httputil"
)

func RequestLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println(ctx.Request.Host, ctx.Request.RemoteAddr, ctx.Request.RequestURI)


		// Save a copy of this request for debugging.
		requestDump, err := httputil.DumpRequest(ctx.Request, true)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(requestDump))


		ctx.Next()
	}
}

func ResponseLogger() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("X-Content-Type-Options", "nosniff")

        c.Next()

        fmt.Printf("%d %s %s\n",
            c.Writer.Status(),
            c.Request.Method,
            c.Request.RequestURI,
        )
    }
}
func main() {
    gin.SetMode(gin.ReleaseMode)
    r := gin.Default()
 
    // Add logging middleware
    r.Use(RequestLogger())
    r.Use(ResponseLogger())

    r.GET("/", func(c *gin.Context) {
    })
	
    r.Run(":80")
}
