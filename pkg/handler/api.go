package handler

import (
	"fmt"
	"github.com/egorrridze/payment-emulator/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)


func (h *Handler) createPayment (c *gin.Context) {
	var input models.Payment
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
		"status": status,
	})
}

func (h *Handler) updateStatus (c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	rowsCounter, newStatus, err := h.services.UpdateStatus(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if rowsCounter != 0 {
		c.JSON(http.StatusOK, map[string]interface{}{
			"new status": newStatus,
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"update status": fmt.Sprintf("updated %b rows", rowsCounter),
		})
	}


}

func (h *Handler) getStatusById (c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	status, err := h.services.GetStatusById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, status)
}

type getAllPaymentsResponse struct {
	Data []models.Payment `json:"data"`
}

func (h *Handler) getAllPayments (c *gin.Context) {
	userId := c.Query("user_id")
	userEmail := c.Query("user_email")
	var payments []models.Payment

	if userId != "" {
		userIdInt, err := strconv.Atoi(userId)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, "invalid id param")
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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	rowsCounter, err := h.services.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if rowsCounter == 0 {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("deleted %b rows", rowsCounter))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"deletion status": fmt.Sprintf("deleted %b rows", rowsCounter),
	})
}

