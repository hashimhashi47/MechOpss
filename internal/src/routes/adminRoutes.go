package routes

import (
	"MechOpss/infra/db"
	"MechOpss/internal/src/controllers"
	"MechOpss/internal/src/middleware"
	"MechOpss/internal/src/repository"
	"MechOpss/internal/src/services"

	"github.com/gin-gonic/gin"
)

func AdminRoute(r *gin.Engine) {

	// admin
	adminRepo := repository.Newrepo(db.DB)
	adminService := services.NewAdminService(adminRepo)
	adminController := controllers.NewAdminController(adminService)

	//user
	userRepo := repository.Newrepo(db.DB)
	userService := services.NewUserservices(userRepo)
	UserController := controllers.NewController(userService)

	r.GET("admin/login", adminController.LoginPage)
	r.POST("admin/login", adminController.AdminLoginHandler)

	Admin := r.Group("/admin")
	Admin.Use(middleware.AdminAuth())
	{
		// Pages
		Admin.GET("/logout", adminController.AdminLogout)
		Admin.GET("/dashboard", adminController.AdminDashboardPage)
		Admin.GET("/users", adminController.ManageUsers)
		Admin.GET("/staff", adminController.ManageStaff)
		Admin.GET("/bookings", adminController.ManageBookings)
		Admin.GET("/booked", adminController.Booked)
		Admin.GET("/slots", adminController.ManageSlots)

		// Dashboard Details
		Admin.GET("/users/count", adminController.GetUsersCount)
		Admin.GET("/staff/count", adminController.GetStaffCount)
		Admin.GET("/bookings/count", adminController.GetBookingCount)
		Admin.GET("/bookings/next", adminController.GetNextService)
		Admin.GET("/bookings/recent", adminController.GetRecentBookings)
		Admin.GET("/bookings/manage", adminController.ManageBookings)

		// Users
		Admin.GET("/getallusers", adminController.GetAllUsers)
		Admin.PUT("/getallusers/:id", adminController.UpdateUser)
		Admin.DELETE("/delete/:id", adminController.DeleteUser)
		Admin.PUT("/block/user/:id", adminController.Blockuser)
		Admin.POST("/adduser", UserController.RegisterAuth)

		// Staff
		Admin.POST("/addstaff", adminController.AddStaff)
		Admin.GET("/getstaff", adminController.GetStaff)
		Admin.PUT("/updateStaff/:id", adminController.UpdateStaff)
		Admin.PUT("/blockstaff/:id", adminController.BlockStaff)
		Admin.DELETE("/deletestaff/:id", adminController.DeleteStaff)

		//Booking
		Admin.GET("/getbookings", adminController.GetBooking)
		Admin.POST("/approve/:id", adminController.ApproveBooking)

		//boooked
		Admin.GET("/getbooked", adminController.GetBooked)
		Admin.POST("/assignstaff/:id", adminController.AssignStaff)
		Admin.PUT("/update/booked/:id", adminController.UpdateBooked)
		Admin.POST("/assign/:id", adminController.AsignBooking)
		Admin.GET("/rejectbooking/:id", adminController.RejectBooking)

		//slots
		Admin.GET("/getslots", adminController.GetSlots)
		Admin.POST("/addslot/:id", adminController.AddSlot)
		Admin.GET("/getslotcount", adminController.CountOfSlot)
		Admin.PUT("/slot/update/:id", adminController.EditSlot)
		Admin.GET("/slot/empty/:id", adminController.EmptySlot)
	}

}
