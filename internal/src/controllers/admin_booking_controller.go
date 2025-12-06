package controllers

import (
	"MechOpss/internal/src/constants"
	"MechOpss/internal/src/models"
	"MechOpss/internal/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

//get booking
func (ac *AdminController) GetBooking(c *gin.Context) {
	var bookings []models.Booking
	var err error
	bookings, err = ac.Service.ServiceGetAllBooking()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.ErrorMessage(constants.BADREQUEST, err)})
		return
	}
	c.JSON(http.StatusOK,gin.H{ "Success": utils.SuccessResponse(bookings)})
}

//approve booking
func (ac *AdminController) ApproveBooking(c *gin.Context) {
	id := c.Param("id")
	var booked models.Booked
	var err error
	booked, err = ac.Service.ServiceApproveBooking(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": utils.ErrorMessage(constants.BADREQUEST, err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Succes": utils.SuccessResponseMsg(booked, "Booked succcesfully")})

}
