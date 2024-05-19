package handler

import "github.com/gin-gonic/gin"

type ModificatorOrder int

const (
	BeforeRequest ModificatorOrder = iota
	AfterRequest
)

type Modificator interface {
	Setup() gin.HandlerFunc
	Order() ModificatorOrder
}
