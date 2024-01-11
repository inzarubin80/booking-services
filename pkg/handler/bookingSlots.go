package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/inzarubin80/booking-services"
)


func (h *Handler) createBookingSlots(c *gin.Context) {

	var input booking.BookingSlots
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.BookingSlots.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

