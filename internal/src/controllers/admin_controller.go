package controllers

import (
	"MechOpss/internal/src/services"
	"MechOpss/internal/src/utils"
	"MechOpss/internal/src/utils/constants"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

// indirect conncetion with repository using dependency injection
// strcut to access the service
type AdminController struct {
	Service *services.AdminService
}

// constructor
func NewAdminController(s *services.AdminService) *AdminController {
	return &AdminController{Service: s}
}

// ---------------- Login ----------------
func (ac *AdminController) AdminLoginHandler(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	fmt.Println("✅", email, password)
	admin, NewAccessToken, _, err := ac.Service.AdminLogin(email, password)
	if err != nil {
		c.HTML(http.StatusUnauthorized, "Login.html", gin.H{"error": utils.ErrorMessage(constants.BADREQUEST, err)})
		return
	}

	c.SetCookie("admin_id", NewAccessToken, 7*24*3600, "/", "localhost", false, true)
	c.Set("admin_id", admin.ID)

	c.Redirect(http.StatusSeeOther, "/admin/dashboard")
}

func (ac *AdminController) AdminLogout(c *gin.Context) {
	id, _ := c.Get("admin_id")
	adminID := fmt.Sprintf("%v", id)

	if err := ac.Service.ServiceAdminLogout(adminID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.ErrorMessage(constants.BADREQUEST, err)})
		return
	}
	c.SetCookie("admin_id", "", -1, "/", "", false, true)
	// Redirect
	c.Redirect(http.StatusSeeOther, "/admin/login")
}

// ---------------- Pages ----------------
func (ac *AdminController) AdminDashboardPage(c *gin.Context) {
	c.HTML(http.StatusOK, "AdminDashboard.html", gin.H{})
}

func (ac *AdminController) Booked(c *gin.Context) {
	c.HTML(http.StatusOK, "Booked.html", nil)
}

func (ac *AdminController) ManageBookings(c *gin.Context) {
	c.HTML(http.StatusOK, "BookingManagement.html", nil)
}

func (ac *AdminController) ManageSlots(c *gin.Context) {
	c.HTML(http.StatusOK, "SlotOnGarage.html", nil)
}

func (ac *AdminController) ManageStaff(c *gin.Context) {
	c.HTML(http.StatusOK, "StaffPage.html", nil)
}

func (ac *AdminController) ManageUsers(c *gin.Context) {
	c.HTML(http.StatusOK, "UsersPage.html", nil)
}

func (ac *AdminController) LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "Login.html", nil)
}
