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
	//staff controller
	repo := repository.Newrepo(db.DB)
	service := services.NewStaffService(repo)
	staffController := controllers.NewStaffController(service)

	// admin
	adminRepo := repository.Newrepo(db.DB)
	adminService := services.NewAdminService(adminRepo)
	adminController := controllers.NewAdminController(adminService)

	//user
	userRepo := repository.Newrepo(db.DB)
	userService := services.NewUserservices(userRepo)
	UserController := controllers.NewController(userService)

	r.POST("/staff/login", staffController.StaffLogin)

	staff := r.Group("/staff")
	staff.Use(middleware.Middleware(constants.Staff))
	{
		staff.POST("/logout", UserController.UserLogout)
		staff.POST("/updateprofile", adminController.UpdateStaff)
		staff.GET("/getallbookeds", staffController.StaffCheckBookeds)
		staff.POST("/updatestatus/:id", staffController.UpdateStatus)
		staff.GET("/getslots", staffController.GetSlots)
		staff.POST("/editslots/:id" , adminController.EditSlot)
	}
}
