package ginmiddleware

import (
	"context"

	"github.com/gin-gonic/gin"
)

type contextKey string

const (
	ginContextKey contextKey = "GinContext"
	userDataKey   contextKey = "UserData"
)

func getRequestContext(
	c *gin.Context,
	options *Options,
) context.Context {
	requestContext := context.WithValue(context.Background(), ginContextKey, c)
	if options != nil {
		requestContext = context.WithValue(requestContext, userDataKey, options.UserData)
	}

	return requestContext
}
