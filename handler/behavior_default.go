package handler

import (
	"github.com/gin-gonic/gin"
)

type DefaultBehavior struct {
	Mods []Modificator
}

func (b *DefaultBehavior) Init() {
	b.Mods = []Modificator{}
}

func (b *DefaultBehavior) Handle(ctx *gin.Context) {
}

func (b *DefaultBehavior) BeforeMiddleware() []gin.HandlerFunc {
	var result []gin.HandlerFunc = []gin.HandlerFunc{}

	for _, mod := range b.Mods {
		if mod.Order() == BeforeRequest {
			result = append(result, mod.Setup())
		}
	}
	return result
}

func (b *DefaultBehavior) AfterMiddleware() []gin.HandlerFunc {
	var result []gin.HandlerFunc = []gin.HandlerFunc{}

	for _, mod := range b.Mods {
		if mod.Order() == AfterRequest {
			result = append(result, mod.Setup())
		}
	}
	return result
}
