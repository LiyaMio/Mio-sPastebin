package Middle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)
type Middleware func(http.HandlerFunc) http.HandlerFunc
func Logger() gin.HandlerFunc {
	return func(context *gin.Context) {
		host := context.Request.Host
		url := context.Request.URL
		method := context.Request.Method
		start := time.Now()
		defer func() { log.Println(context.Request.URL.Path, time.Since(start)) }()
		log.Printf("%s::%s \t %s \t %s ", time.Now().Format("2006-01-02 15:04:05"), host, url, method)
		log.Println()
		context.Next()
		fmt.Println(context.Writer.Status())
	}
}