package middleware

import (
	"bytes"
	"go-clean-template/internal/entity"
	"go-clean-template/pkg/jwt"
	"io"
	"net/http"

	"go.uber.org/zap"

	"github.com/niudaii/util"

	"github.com/gin-gonic/gin"
)

func Operation() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body []byte
		if c.Request.Method != http.MethodGet {
			var err error
			body, err = io.ReadAll(c.Request.Body)
			if err == nil {
				c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
			}
		}
		operation := entity.Operation{
			Operator: jwt.GetUsername(c),
			IP:       c.ClientIP(),
			Agent:    c.Request.UserAgent(),
			Method:   c.Request.Method,
			Path:     c.Request.URL.Path,
			Query:    c.Request.URL.RawQuery,
			Body:     string(body),
		}

		writer := responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer
		c.Next()

		operation.Status = c.Writer.Status()
		operation.Resp = writer.body.String()
		operation.Resp = operation.Resp[:util.Min(1024, len(operation.Resp))]
		// logger or db
		zap.L().Sugar().Named("request").Infof("%v %v %v\n%v", operation.Method, operation.Path, operation.Query, operation.Body)
		zap.L().Sugar().Named("response").Infof("%v\n%v", operation.Status, operation.Resp)
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
