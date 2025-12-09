package controllers

import (
	"MechOpss/internal/src/constants"
	"MechOpss/internal/src/models"
	"MechOpss/internal/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// get all slots
func (ac *AdminController) GetSlots(c *gin.Context) {
	var Slots []models.Slot
	Slots, err := ac.Service.ServiceGetAllSlots()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.ErrorMessage(constants.BADREQUEST, err)})
		return
	}

	c.JSON(http.StatusOK,gin.H{"success": utils.SuccessResponseMsg(Slots,"get slots succesfully")})

}



