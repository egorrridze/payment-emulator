package handler

import (
	"github.com/egorrridze/payment-emulator"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)


func (h *Handler) createPayment (c *gin.Context) {
	var input emulator.Payment
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, status, err := h.services.Payment.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
		"payment_status": status,
	})
}

func (h *Handler) updateStatus (c *gin.Context) {
}

func (h *Handler) getStatus (c *gin.Context) {
}

type getAllPaymentsResponse struct {
	Data []emulator.Payment `json:"data"`
}

func (h *Handler) getAllPayments (c *gin.Context) {
	userId := c.Query("user_id")
	userEmail := c.Query("user_email")
	var payments []emulator.Payment

	if userId != "" {
		userIdInt, err := strconv.Atoi(userId)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		payments, err = h.services.GetAllById(userIdInt)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	} else if userEmail != "" {
		var err error
		payments, err = h.services.GetAllByEmail(userEmail)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, getAllPaymentsResponse{
		Data: payments,
	})
}


func (h *Handler) deletePayment (c *gin.Context) {
}

