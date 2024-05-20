package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type InputModificator[I any] struct {
	Input *I
}

func (i *InputModificator[I]) Setup() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := ctx.ShouldBindJSON(&i.Input); err != nil {
			ctx.AbortWithStatusJSON(400, gin.H{"error": http.StatusText(400)})
			return
		}
		ctx.Next()
	}
}

func (i *InputModificator[I]) Order() ModificatorOrder {
	return BeforeRequest
}
