package main

import (
	"MechOpss/infra/db"
	"MechOpss/internal/src/controllers"
	"MechOpss/internal/src/repository"
	"MechOpss/internal/src/routes"
	"MechOpss/internal/src/services"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//1
	DB := db.Connection()
	//2
	UserRepo := repository.Newrepo(DB)
	//3
	userService := services.NewUserservices(UserRepo)
	//4
	userController := controllers.NewController(userService)
	//5
	routes.UserRoutes(r, userController)
	r.LoadHTMLGlob("../templates/*.html")
	routes.StaffRoutes(r)
	routes.AdminRoute(r)

	r.Run(":8008")
}
