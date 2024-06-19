package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaodeluchi/ms-users/application/services"
	"github.com/joaodeluchi/ms-users/domain"
)

type UserHandler struct {
	userService services.UserService
}

func (uh UserHandler) CreateUserHandler(c *gin.Context) {
	var newUser domain.User
	err := c.BindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err = uh.userService.CreateUser(newUser)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully!"})
}

func (uh UserHandler) GetUserHandler(c *gin.Context) {
	userID := c.Param("id")

	user, err := uh.userService.GetUser(userID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uh UserHandler) UpdateRoles(c *gin.Context) {
	userID := c.Param("id")

	var user domain.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err = uh.userService.UpdateUserRoles(userID, user.Roles)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User roles updated successfully!"})
}

func (uh UserHandler) DeleteUser(c *gin.Context) {
	userID := c.Param("id")

	err := uh.userService.DeleteUser(userID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "User deleted")
}

func NewUserHandler(service services.UserService) UserHandler {
	return UserHandler{userService: service}
}
