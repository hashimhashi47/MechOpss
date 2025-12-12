package controllers

import (
	"MechOpss/internal/src/models"
	"MechOpss/internal/src/utils"
	"MechOpss/internal/src/utils/constants"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// add staffs to database
func (ac *AdminController) AddStaff(c *gin.Context) {
	var Input struct {
		FirstName  string `json:"firstname" binding:"required"`
		LastName   string `json:"lastname"  binding:"required"`
		Email      string `json:"email" binding:"required"`
		Password   string `json:"password" binding:"required"`
		Department string `json:"department" binding:"required"`
	}

	if err := c.ShouldBindJSON(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Binding erorr": utils.ErrorMessage(constants.BADREQUEST, err)})
		return
	}

	AddStaff := models.Staff{
		FirstName:  Input.FirstName,
		LastName:   Input.LastName,
		Email:      Input.Email,
		Password:   Input.Password,
		Department: Input.Department,
		Role:       constants.Staff,
	}

	var erorr error
	data, erorr := ac.Service.ServiceStaffRegister(&AddStaff)

	if erorr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.ErrorMessage(constants.BADREQUEST, erorr)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Success": utils.SuccessResponseMsg(data, "staff signup succesfully")})
}

func (ac *AdminController) GetStaff(c *gin.Context) {
	var staffs []models.Staff
	staffs, err := ac.Service.ServiceGetStaff()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.ErrorMessage(constants.BADREQUEST, err)})
	}

	c.JSON(http.StatusOK, gin.H{"Success": utils.SuccessResponseMsg(staffs, "Succesfully find users")})
}


//upadte staff on admin and staff side
func (ac *AdminController) UpdateStaff(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		staffID, _ := c.Get("id")
		id = fmt.Sprintf("%v", staffID)
	}

	var Input models.Staff

	if err := c.ShouldBindJSON(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	data, err := ac.Service.ServiceUpdateStaff(Input, id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.ErrorMessage(constants.BADREQUEST, err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Success": utils.SuccessResponseMsg(data, "succesfully updated staff details")})
}


//block staff
func (ac *AdminController) BlockStaff(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Block *bool `json:"block"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if body.Block == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.ErrorMessage(constants.BADREQUEST, errors.New("block value is required"))})
		return
	}

	data, err := ac.Service.ServiceBlockStaff(id, *body.Block)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": utils.ErrorMessage(constants.BADREQUEST, err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": utils.SuccessResponseMsg(data, "Blocked succesfully")})
}

func (as *AdminController) DeleteStaff(c *gin.Context) {
	id := c.Param("id")

	if err := as.Service.ServiceDeleteuser(models.Staff{}, id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": utils.ErrorMessage(constants.BADREQUEST, err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": utils.SuccessResponseMsg("", "deleted succesfully"),
	})
}
