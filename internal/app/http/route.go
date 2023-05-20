package http

import (
	"goros-server/internal/app/http/handler"
	"goros-server/internal/app/http/middlewares"

	"github.com/gin-gonic/gin"
)

func Register(rg *gin.RouterGroup, mdw ...Middleware) error {
	middleware := middleware{
		authentication: middlewares.DefaultAuthMiddleware,
	}
	for _, o := range mdw {
		o.apply(&middleware)
	}

	rg.POST("/audio",
		middleware.authentication,
		handler.AudioHandler)
	rg.POST("/video",
		middleware.authentication,
		handler.VideoHandler)
	rg.POST("/webview",
		middleware.authentication,
		handler.WebviewHandler)
	rg.POST("/motor",
		middleware.authentication,
		handler.MotorHandler)

	return nil
}
