package routes

import (
	"MechOpss/internal/src/constants"
	"MechOpss/internal/src/controllers"
	"MechOpss/internal/src/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(e *gin.Engine, userController *controllers.UserController) {
	User := e.Group("/user")

	{
		User.POST("/signup", userController.RegisterAuth)
		User.POST("/login", userController.LoginAuth)
		User.POST("/logout", userController.UserLogout)
		User.POST("/bookingservice", middleware.Middleware(constants.User), userController.UserBooking)

	}

}
