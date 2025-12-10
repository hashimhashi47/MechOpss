package controllers

import (
	"MechOpss/internal/src/models"
	"MechOpss/internal/src/utils"
	"MechOpss/internal/src/utils/constants"
	"net/http"
	"github.com/gin-gonic/gin"
)

// get all users
func (a *AdminController) GetAllUsers(c *gin.Context) {
	var users []models.User
	var err error
	users, err = a.Service.GetUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.ErrorMessage(constants.BADREQUEST, err)})
	}
	c.JSON(http.StatusOK, gin.H{"Sucess": utils.SuccessResponseMsg(users, "Succesfully find users")})
}

// update user
func (a *AdminController) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := a.Service.ServiceUpdateUser(input, id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": utils.ErrorMessage(constants.BADREQUEST, err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": utils.SuccessResponseMsg(user, "User updated successfully"),
	})
}

//  delete the user from database
func (a *AdminController) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if err := a.Service.ServiceDeleteuser(models.User{}, id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": utils.ErrorMessage(constants.BADREQUEST, err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": utils.SuccessResponseMsg("", "deleted succesfully"),
	})
}

// Block / UnBlock user
func (a *AdminController) Blockuser(c *gin.Context) {
	id := c.Param("id")

	var Body struct {
		Block *bool `json:"block"`
	}

	if err := c.ShouldBindJSON(&Body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"invalid json": utils.ErrorMessage(constants.BADREQUEST, err)})
		return
	}

	var user models.User
	var err error
	user, err = a.Service.ServiceBlockUser(id,  *Body.Block)
	if err != nil {

	}
	c.JSON(http.StatusOK, gin.H{"successully blocked user": utils.SuccessResponseMsg(user, "Blocked succesfully")})
}

