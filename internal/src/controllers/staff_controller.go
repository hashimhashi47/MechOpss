package controllers

import (
	"MechOpss/internal/src/services"
	"MechOpss/internal/src/utils"
	"MechOpss/internal/src/utils/constants"
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

func (sc *StaffController) ProfileUpdate(c *gin.Context) {

}


