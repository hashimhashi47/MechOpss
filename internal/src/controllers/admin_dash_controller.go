package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


// user count
func (ac *AdminController) GetUsersCount(c *gin.Context) {
	count, err := ac.Service.UsersCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"count": count})
}

// staff count
func (ac *AdminController) GetStaffCount(c *gin.Context) {
	count, err := ac.Service.StaffCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"count": count})
}

// booking count
func (ac *AdminController) GetBookingCount(c *gin.Context) {
	count, err := ac.Service.BookingCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"count": count})
}

// next service
func (ac *AdminController) GetNextService(c *gin.Context) {
	booking, err := ac.Service.NextBooking()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, booking)
}

// recent bookings
func (ac *AdminController) GetRecentBookings(c *gin.Context) {
	bookings, err := ac.Service.RecentBookings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bookings)
}
