package handler

import "github.com/gin-gonic/gin"

type OutputModificator[O any] struct {
	Output *O
}

func (o *OutputModificator[O]) Setup() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var status int = 200

		if ctx.Writer.Status() != 0 {
			status = ctx.Writer.Status()
		}

		ctx.JSON(status, o.Output)
		ctx.Next()
	}
}

func (o *OutputModificator[O]) Order() ModificatorOrder {
	return AfterRequest
}
