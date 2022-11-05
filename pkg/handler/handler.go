package handler

import (
	"github.com/SyberiaEmperor/avito_task/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service	
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{service: s}
}

func(h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group(":id/account")
	{
		api.GET("/",h.getAccountInfo)
		api.POST("/deposit",h.deposit)
		api.POST("/debit",h.debit)
		api.POST("/transfer",h.transfer)
	}

	return router
}