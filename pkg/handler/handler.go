package handler

import "github.com/gin-gonic/gin"

type Handler struct {

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
			payments.GET("/status/:id", h.getStatus)
			payments.PUT("/", h.getAllPayments)
			//api.GET("/payment", h.getAllPaymentsByEmail)
			payments.DELETE("/:id", h.deletePayment)
		}
	}

	return router
}