package controllers

import (
	"MechOpss/internal/src/constants"
	"MechOpss/internal/src/models"
	"MechOpss/internal/src/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (ac *AdminController) GetBooked(c *gin.Context) {
	Booked, err := ac.Service.ServiceGetBooked()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.ErrorMessage(constants.BADREQUEST, err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Success": utils.SuccessResponse(Booked)})
}

func (ac *AdminController) AssignStaff(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		StaffID string `json:"staff_id"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	data, err := ac.Service.AssignStaffService(id, body.StaffID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Success": data})
}

func (ac *AdminController) UpdateBooked(c *gin.Context) {
	id := c.Param("id")

	var Input models.Booked
	if err := c.ShouldBindJSON(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	Booked, err := ac.Service.ServiceUpadteBooked(id, Input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.ErrorMessage(constants.BADREQUEST, err)})
	}

	c.JSON(http.StatusOK, gin.H{"success": utils.SuccessResponseMsg(Booked, "updated succesfully")})
}
