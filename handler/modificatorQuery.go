package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type QueryModificator[Q any] struct {
	Query Q
}

func (q *QueryModificator[Q]) Setup() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := ctx.ShouldBindQuery(&q.Query); err != nil {
			ctx.AbortWithStatusJSON(400, gin.H{"error": http.StatusText(400)})
			return
		}
		ctx.Next()
	}
}

func (q *QueryModificator[Q]) Order() ModificatorOrder {
	return BeforeRequest
}
