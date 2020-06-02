package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// RequestLogger logs all the incoming requests to file and standard output.
func RequestLogger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[%s]:-[%s] - [%s] \"%s %s [%d] %s \"%s\" %s\"\n",
			param.ClientIP,
			param.Method,
			param.TimeStamp.Format(time.RFC1123),
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	})
}
