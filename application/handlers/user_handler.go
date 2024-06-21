package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaodeluchi/ms-users/application/services"
	"github.com/joaodeluchi/ms-users/domain"
)

type UserHandler struct {
	userService services.UserService // dependency on UserService
}

func (uh UserHandler) CreateUserHandler(c *gin.Context) {
	var newUser domain.User
	err := c.BindJSON(&newUser)
	if err != nil {
		// Return status 400 for invalid request body
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err = uh.userService.CreateUser(newUser)
	if err != nil {
		// Return status 409 for conflict (e.g., user already exists)
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	// Return status 201 for successful creation
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully!"})
}

func (uh UserHandler) GetUserHandler(c *gin.Context) {
	userID := c.Param("id")

	user, err := uh.userService.GetUser(userID)
	if err != nil {
		// Return status 404 for not found (e.g., user not found)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	// Return status 200 and the user data
	c.JSON(http.StatusOK, user)
}

func (uh UserHandler) UpdateRoles(c *gin.Context) {
	userID := c.Param("id")

	var user domain.User
	err := c.BindJSON(&user)
	if err != nil {
		// Return status 400 for invalid request body
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err = uh.userService.UpdateUserRoles(userID, user.Roles)
	if err != nil {
		// Return status 404 if user not found
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	// Return status 200 for successful role update
	c.JSON(http.StatusOK, gin.H{"message": "User roles updated successfully!"})
}

func (uh UserHandler) DeleteUser(c *gin.Context) {
	userID := c.Param("id")

	err := uh.userService.DeleteUser(userID)
	if err != nil {
		// Return status 404 if user not found
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	// Return status 200 for successful deletion
	c.JSON(http.StatusOK, "User deleted")
}

func NewUserHandler(service services.UserService) UserHandler {
	return UserHandler{userService: service}
}
