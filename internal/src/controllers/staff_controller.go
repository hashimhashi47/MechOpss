package controllers

import (
	"MechOpss/internal/src/services"
	"MechOpss/internal/src/utils"
	"MechOpss/internal/src/utils/constants"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// indirect conncetion with service using dependency injection
type StaffController struct {
	Service *services.StaffService
}

// constrctor
func NewStaffController(s *services.StaffService) *StaffController {
	return &StaffController{Service: s}
}

// staff login
func (sc *StaffController) StaffLogin(c *gin.Context) {
	var Input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.ErrorMessage(constants.BADREQUEST, err)})
		return
	}

	staff_id, data, accesstoken, refershtoken, err := sc.Service.ServiceStaffLogin(Input.Email, Input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.ErrorMessage(constants.BADREQUEST, err)})
		return
	}

	c.SetCookie("Token", accesstoken, 7*24*3600, "/", "localhost", false, true)
	c.Set("id", staff_id)
	c.JSON(http.StatusOK, gin.H{"Success": data, "accesstoken": accesstoken, "refershtoken": refershtoken})
}

//staff check assiganed bookeds
func (sc *StaffController) StaffCheckBookeds(c *gin.Context) {
	id, _ := c.Get("id")
	staffID := fmt.Sprintf("%v", id)

	data, err := sc.Service.ServiceCheckStaffBookeds(staffID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.ErrorMessage(constants.BADREQUEST, err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Succes": utils.SuccessResponseMsg(data, "bookeds data found succesfully")})
}


//staff can update the booked status
func (sc *StaffController) UpdateStatus(c *gin.Context) {
	id := c.Param("id")

	var input struct {
		Status      string `json:"status"`
		Description string `json:"description"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.ErrorMessage(constants.BADREQUEST, err)})
		return
	}

	data, err := sc.Service.ServiceUpdateStatus(input.Status, input.Description, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.ErrorMessage(constants.BADREQUEST, err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Succes": utils.SuccessResponseMsg(data, "bookeds data found succesfully")})
}


//staff can get assiganed slots
func (sc *StaffController) GetSlots(c *gin.Context) {
	id, _ := c.Get("id")
	staffID := fmt.Sprintf("%v", id)

	data, err := sc.Service.ServiceGetSlots(staffID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.ErrorMessage(constants.BADREQUEST, err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Succes": utils.SuccessResponseMsg(data, "slots found succesfully")})
}

