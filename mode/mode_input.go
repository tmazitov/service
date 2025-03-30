package mode

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type InputMode[I any] struct {
	Input *I
}

func (i *InputMode[I]) Setup() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		v := reflect.ValueOf(i.Input).Elem()
		v.Set(reflect.Zero(v.Type()))

		// fmt.Printf("Input before: %v\n", i.Input)
		if err := ctx.ShouldBindJSON(&i.Input); err != nil {
			ctx.AbortWithStatusJSON(400, gin.H{"error": http.StatusText(400)})
			return
		}
		// fmt.Printf("Input after: %v\n", i.Input)
		ctx.Next()
	}
}

func (i *InputMode[I]) Order() ModeOrder {
	return BeforeRequest
}
