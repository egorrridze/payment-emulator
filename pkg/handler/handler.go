package handler

import (
	"github.com/egorrridze/payment-emulator/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		payments := api.Group("/payments")
		{
			payments.POST("/", h.createPayment)
			payments.POST("/status/:id", h.updateStatus)
			payments.GET("/status/:id", h.getStatusById)
			payments.GET("/", h.getAllPayments)
			payments.DELETE("/:id", h.deletePayment)
		}
	}

	return router
}