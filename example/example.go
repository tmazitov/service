package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tmazitov/service"
	"github.com/tmazitov/service/handler"
)

type PingMessage struct {
	Message string `json:"message"`
}

type PingHandler struct {
	handler.CoreBehavior[PingMessage, PingMessage, PingMessage]
}

func (h *PingHandler) Handle(ctx *gin.Context) {
	fmt.Println("Handling request", h.Input.Message)
	h.Output.Message = h.Input.Message + h.Query.Message
}

func main() {
	config := &service.ServiceConfig{
		Name:    "example",
		Port:    8080,
		Prefix:  "exm",
		Version: "v1",
	}

	s := service.NewService(config)

	s.SetupHandlers([]service.Endpoint{
		{
			Method:  "POST",
			Path:    "ping",
			Handler: &PingHandler{},
		},
	})

	s.Start()
}
