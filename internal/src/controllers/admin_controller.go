package controllers

import (
	"MechOpss/internal/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)


//indirect conncetion with repository using dependency injection
//strcut to access the service
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

	admin, accessToken, _, err := ac.Service.AdminLogin(email, password)
	if err != nil {
		c.HTML(http.StatusUnauthorized, "Login.html", gin.H{
			"error": err.Error(),
		})
		return
	}

	c.SetCookie("admin_id", accessToken, 3600*24*7, "/", "localhost", false, true)
	c.Set("admin_id", admin.ID)

	// redirect to dashboard
	c.Redirect(http.StatusSeeOther, "/admin/dashboard")
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
