package handler

import (
	"github.com/gin-gonic/gin"
	mode "github.com/tmazitov/service/mode"
)

type DefaultBehavior struct {
	Mods []mode.IMode
}

func (b *DefaultBehavior) Init() {
	b.Mods = []mode.IMode{}
}

func (b *DefaultBehavior) Handle(ctx *gin.Context) {
}

func (b *DefaultBehavior) BeforeMiddleware() []gin.HandlerFunc {
	var result []gin.HandlerFunc = []gin.HandlerFunc{}

	for _, mod := range b.Mods {
		if mod.Order() == mode.BeforeRequest {
			result = append(result, mod.Setup())
		}
	}
	return result
}

func (b *DefaultBehavior) AfterMiddleware() []gin.HandlerFunc {
	var result []gin.HandlerFunc = []gin.HandlerFunc{}

	for _, mod := range b.Mods {
		if mod.Order() == mode.AfterRequest {
			result = append(result, mod.Setup())
		}
	}
	return result
}
