package routes

import (
	"MechOpss/internal/src/controllers"
	"MechOpss/internal/src/middleware"
	"MechOpss/internal/src/utils/constants"

	"github.com/gin-gonic/gin"
)

func UserRoutes(e *gin.Engine, userController *controllers.UserController) {
	
	User := e.Group("/user")

	{
		User.POST("/signup", userController.RegisterAuth)
		User.POST("/login", userController.LoginAuth)
		User.POST("/logout", middleware.Middleware(constants.User), userController.UserLogout)
		User.POST("/bookingservice", middleware.Middleware(constants.User), userController.UserBooking)
	}

}
