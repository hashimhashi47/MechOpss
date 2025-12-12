package controllers

import (
	"MechOpss/internal/src/models"
	"MechOpss/internal/src/utils"
	"MechOpss/internal/src/utils/constants"
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

	c.JSON(http.StatusOK, gin.H{"success": utils.SuccessResponseMsg(Slots, "get slots succesfully")})

}

// get the count of slot
func (ac *AdminController) CountOfSlot(c *gin.Context) {

	count, err := ac.Service.ServiceCountOfSlot()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.ErrorMessage(constants.BADREQUEST, err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Success": utils.SuccessResponse(count)})
}

// upadte slot
func (s *AdminController) EditSlot(c *gin.Context) {
	id := c.Param("id")

	
	var input models.Slot

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.ErrorMessage(constants.BADREQUEST, err)})
		return
	}
	data, err := s.Service.ServiceUpdateSlot(input, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.ErrorMessage(constants.BADREQUEST, err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"succes": utils.SuccessResponseMsg(data, "slot edited succesfully")})
}

func (s *AdminController) RemoveSlot(c *gin.Context) {
	id := c.Param("id")
	if err := s.Service.ServiceRemoveSlot(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.ErrorMessage(constants.BADREQUEST, err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"succes": "deleted succesfully"})
}

func (ac *AdminController) EmptySlot(c *gin.Context) {
	id := c.Param("id")
	data, err := ac.Service.ServiceEmptySlot(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.ErrorMessage(constants.BADREQUEST, err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": utils.SuccessResponse(data)})
}
