package handler

import (
	"github.com/gin-gonic/gin"
)

type CoreBehavior[I, O, Q any] struct {
	InputModificator[I]
	QueryModificator[Q]
	OutputModificator[O]
	Mods []Modificator
}

func (h *CoreBehavior[I, O, Q]) Init() {
	h.Mods = []Modificator{
		&h.InputModificator,
		&h.QueryModificator,
		&h.OutputModificator,
	}
}

func (h *CoreBehavior[I, O, Q]) CoreAfterMiddleware() []gin.HandlerFunc {
	var result []gin.HandlerFunc = []gin.HandlerFunc{}

	for _, mod := range h.Mods {
		if mod.Order() == AfterRequest {
			result = append(result, mod.Setup())
		}
	}

	return result
}

func (h *CoreBehavior[I, O, Q]) CoreBeforeMiddleware() []gin.HandlerFunc {
	var result []gin.HandlerFunc = []gin.HandlerFunc{}

	for _, mod := range h.Mods {
		if mod.Order() == BeforeRequest {
			result = append(result, mod.Setup())
		}
	}

	return result
}
