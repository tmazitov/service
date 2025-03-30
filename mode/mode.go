package mode

import "github.com/gin-gonic/gin"

type ModeOrder int

const (
	BeforeRequest ModeOrder = iota
	AfterRequest
)

type IMode interface {
	Setup() gin.HandlerFunc
	Order() ModeOrder
}
