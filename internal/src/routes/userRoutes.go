package routes

import (
	"MechOpss/infra/db"
	"MechOpss/internal/src/controllers"
	"MechOpss/internal/src/middleware"
	"MechOpss/internal/src/repository"
	"MechOpss/internal/src/services"
	"MechOpss/internal/src/utils/constants"

	"github.com/gin-gonic/gin"
)

func UserRoutes(e *gin.Engine, userController *controllers.UserController) {

	// admin
	adminRepo := repository.Newrepo(db.DB)
	adminService := services.NewAdminService(adminRepo)
	adminController := controllers.NewAdminController(adminService)

	User := e.Group("/user")
	User.Use(middleware.Middleware(constants.User))

	e.POST("/user/signup", userController.RegisterAuth)
	e.POST("/user/login", userController.LoginAuth)

	{
		User.POST("/updateuser", adminController.UpdateUser)
		User.POST("/logout", userController.UserLogout)
		User.POST("/bookingservice", userController.UserBooking)
		User.GET("/getbookingstatus", userController.GetBookingStatus)
		User.GET("/bookingstatus/:id", userController.GetBookingstatusID)
		User.GET("/getbookedstatus", userController.GetBookedsStatus)
		User.GET("/bookedstatus/:id", userController.GetBookedstatusID)
	}

}
