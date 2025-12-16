package controllers

import (
	"MechOpss/internal/src/models"
	"MechOpss/internal/src/utils"
	"MechOpss/internal/src/utils/constants"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// booking the service by user
func (s *UserController) UserBooking(c *gin.Context) {

	UserID := c.MustGet("id").(uint)
	var Input struct {
		CarModel  string `json:"carmodel" binding:"required"`
		CarNumber string `json:"carnumber" binding:"required"`
		FuelType  string `json:"fueltype" binding:"required"`
		Problem   string `json:"problem" binding:"required"`
		Time      string `json:"time" binding:"required"`
		Date      string `json:"date" binding:"required"`
		Address   string `json:"address" binding:"required"`
		LandMark  string `json:"landmark" binding:"required"`
	}

	if err := c.ShouldBindJSON(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": utils.ErrorMessage(constants.BADREQUEST, err)})
		return
	}

	Booking := models.Booking{
		UserID:    UserID,
		CarModel:  Input.CarModel,
		CarNumber: Input.CarNumber,
		FuelType:  Input.FuelType,
		Problem:   Input.Problem,
		Time:      Input.Time,
		Date:      Input.Date,
		Address:   Input.Address,
		LandMark:  Input.LandMark,
	}
	var BookingId string
	var err error
	BookingId, err = s.Service.ServiceBookingUser(Booking)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.ErrorMessage(constants.BADREQUEST, err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Sucess": "Our Team will contact soon", "ID": utils.SuccessResponse(BookingId)})
}

func (uc *UserController) GetPayments(c *gin.Context) {
	id, _ := c.Get("id")
	UserID := fmt.Sprintf("%v", id)

	data, err := uc.Service.ServiceGetPayments(UserID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.ErrorMessage(constants.BADREQUEST, err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Sucess": "your payment details", "payments": utils.SuccessResponse(data)})
}


// pay the service
func (uc *UserController) PayTheService(c *gin.Context) {
	BID := c.Param("id")
	id, _ := c.Get("id")
	UserID := fmt.Sprintf("%v", id)

	var input struct {
		Amount float64 `json:"amount"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": utils.ErrorMessage(constants.BADREQUEST, err)})
	}

	data, err := uc.Service.ServicePayTheService(BID, UserID, input.Amount)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.ErrorMessage(constants.BADREQUEST, err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Sucess": utils.SuccessResponseMsg(data, "amount recived successfully")})
}
