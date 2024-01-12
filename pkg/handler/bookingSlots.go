package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/inzarubin80/booking-services")


// @Summary Booking slots
// @Security ApiKeyAuth
// @Tags BookingSlots
// @Description slot reservation
// @IDBookingSlots
// @Accept  json
// @Produce  json
// @Param input body booking.BookingSlots true "Booking Slots"
// @@Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/bookingSlots [post]
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

