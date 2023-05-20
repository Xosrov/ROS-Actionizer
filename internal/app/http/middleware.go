package http

import "github.com/gin-gonic/gin"

type middleware struct {
	authentication gin.HandlerFunc
}

type Middleware interface {
	apply(*middleware)
}

type authMiddleware struct {
	Auth gin.HandlerFunc
}

func (a authMiddleware) apply(mdw *middleware) {
	mdw.authentication = a.Auth
}

// nolint: ireturn
func WithAuth(handler gin.HandlerFunc) Middleware {
	return authMiddleware{Auth: handler}
}
