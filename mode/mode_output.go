package mode

import (
	"reflect"

	"github.com/gin-gonic/gin"
)

type OutputMode[O any] struct {
	Output *O
}

func (o *OutputMode[O]) Setup() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var status int = 200

		if ctx.Writer.Status() != 0 {
			status = ctx.Writer.Status()
		}

		ctx.JSON(status, o.Output)
		ctx.Next()
		v := reflect.ValueOf(o.Output).Elem()
		v.Set(reflect.Zero(v.Type()))
	}
}

func (o *OutputMode[O]) Order() ModeOrder {
	return AfterRequest
}
