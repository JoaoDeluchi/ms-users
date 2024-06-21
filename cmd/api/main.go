package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaodeluchi/ms-users/application/handlers"
	"github.com/joaodeluchi/ms-users/application/repositories"
	"github.com/joaodeluchi/ms-users/application/services"
)

func main() {
	router := gin.Default()

	registerUserRoutes(router)
	router.GET("/", redirect)
	router.Run(":8080")
}

func redirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/users")
}

func registerUserRoutes(router *gin.Engine) {
	userRouter := router.Group("/users")

	userRepo := repositories.NewUserRepository()
	userServices := services.NewUserService(userRepo)
	userHandlers := handlers.NewUserHandler(userServices)
	userRouter.POST("", userHandlers.CreateUserHandler)
	userRouter.GET("/", userHandlers.GetUserHandler)
	userRouter.GET("/:id", userHandlers.GetUserHandler)
	userRouter.PUT("/:id", userHandlers.UpdateRoles)
	userRouter.DELETE("/:id", userHandlers.DeleteUser)
}
