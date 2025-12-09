package controllers

import (
	"MechOpss/internal/src/services"

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
	// var Input struct {
	// 	Email    string `json:"email"`
	// 	Password string `json:"password"`
	// }
}

func (sc *StaffController) ProfileUpdate(c *gin.Context) {

}

func (sc *StaffController) Logout(c *gin.Context) {

}
