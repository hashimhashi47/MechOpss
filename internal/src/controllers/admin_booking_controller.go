package controllers

import (
	"MechOpss/internal/src/models"
	"MechOpss/internal/src/utils"
	"MechOpss/internal/src/utils/constants"
	"net/http"
	"github.com/gin-gonic/gin"
)

// get booking
func (ac *AdminController) GetBooking(c *gin.Context) {
	var bookings []models.Booking
	var err error
	bookings, err = ac.Service.ServiceGetAllBooking()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.ErrorMessage(constants.BADREQUEST, err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Success": utils.SuccessResponse(bookings)})
}

// approve booking
func (ac *AdminController) ApproveBooking(c *gin.Context) {
	id := c.Param("id")
	var booked models.Bookeds
	var err error
	booked, err = ac.Service.ServiceApproveBooking(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": utils.ErrorMessage(constants.BADREQUEST, err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Succes": utils.SuccessResponseMsg(booked, "Booked succcesfully")})

}

// assign booking
func (ac *AdminController) AsignBooking(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		StaffID string `json:"staff_id"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.ErrorMessage(constants.BADREQUEST, err)})
		return
	}
	var err error
	Data, err := ac.Service.ServiceAssignBooking(id, body.StaffID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.ErrorMessage(constants.BADREQUEST, err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"succes": Data})
}

func (ac *AdminController) RejectBooking(c *gin.Context) {
	id := c.Param("id")
	var err error
	data, err := ac.Service.ServiceRejectBooking(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.ErrorMessage(constants.BADREQUEST, err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": utils.SuccessResponse(data)})
}
