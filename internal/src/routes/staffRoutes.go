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

func StaffRoutes(r *gin.Engine) {
	repo := repository.Newrepo(db.DB)
	service := services.NewStaffService(repo)
	staffController := controllers.NewStaffController(service)

	//user
	userRepo := repository.Newrepo(db.DB)
	userService := services.NewUserservices(userRepo)
	UserController := controllers.NewController(userService)

	staff := r.Group("/staff")
	{
		staff.POST("/login", staffController.StaffLogin)
		staff.POST("/logout", middleware.Middleware(constants.Staff), UserController.UserLogout)
	}
}
