package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ServiceConfig struct {
	Name    string `json:"name"`
	Port    int    `json:"port"`
	Prefix  string `json:"prefix"`
	Version string `json:"version"`
}

type Service struct {
	config *ServiceConfig
	core   *gin.Engine
}

func NewService(config *ServiceConfig) *Service {
	return &Service{
		config: config,
		core:   gin.Default(),
	}
}

func (s *Service) SetupMiddleware(middlewares []gin.HandlerFunc) {
	for _, middleware := range middlewares {
		s.core.Use(middleware)
	}
}

func (s *Service) SetupHandlers(endpoints []Endpoint) {

	var (
		path    string
		handler Handler
		process []gin.HandlerFunc
	)

	for _, e := range endpoints {
		handler = e.Handler
		e.Handler.Init()
		process = []gin.HandlerFunc{}
		process = append(process, handler.CoreBeforeMiddleware()...)
		process = append(process, handler.Handle)
		process = append(process, handler.CoreAfterMiddleware()...)
		if e.Type == "" {
			e.Type = "api"
		}
		path = fmt.Sprintf("/%s/%s/%s/%s", s.config.Prefix, s.config.Version, e.Type, e.Path)
		s.core.Handle(e.Method, path, process...)
	}
}

func (s *Service) GetCore() *gin.Engine {
	return s.core
}

func (s *Service) Start() {
	s.core.Run(fmt.Sprintf(":%d", s.config.Port))
}
