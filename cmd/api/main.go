package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joaodeluchi/ms-users/application/handlers"
	"github.com/joaodeluchi/ms-users/application/repositories"
	"github.com/joaodeluchi/ms-users/application/services"
	"github.com/joaodeluchi/ms-users/framework/database"
)

func main() {
	router := gin.Default()

	registerUserRoutes(router)

	router.Run(":8080")
}

func registerUserRoutes(router *gin.Engine) {
	userRouter := router.Group("/users")
	db := database.NewDb()

	userRepo := repositories.NewUserRepository(db)
	userServices := services.NewUserService(userRepo)
	userHandlers := handlers.NewUserHandler(userServices)
	userRouter.POST("", userHandlers.CreateUserHandler)
	userRouter.GET("/:id", userHandlers.GetUserHandler)
	userRouter.PUT("/:id", userHandlers.UpdateRoles)
	userRouter.DELETE("/:id", userHandlers.GetUserHandler)
}
