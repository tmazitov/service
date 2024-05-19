package service

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	Handle(ctx *gin.Context)
	Init()
	CoreAfterMiddleware() []gin.HandlerFunc
	CoreBeforeMiddleware() []gin.HandlerFunc
}
