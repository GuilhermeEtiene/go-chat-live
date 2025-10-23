package user

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser handles POST requests to create a new user.
// Validates input data, generates password hash and persists to database.
func CreateUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password is required"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating password hash"})
		return
	}
	user.Password = string(hash)

	if err := Create(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Password = ""
	c.JSON(http.StatusCreated, user)
}

// ListUsers handles GET requests to return all registered users.
func ListUsers(c *gin.Context) {
	users, err := List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// GetUserById handles GET requests to find a specific user by ID.
func GetUserById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	user, err := FindById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// UpdateUser handles PUT requests to update an existing user's data.
func UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedData User
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	updatedUser, err := Update(id, &updatedData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

// DeleteUser handles DELETE requests to remove a user from the system.
func DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = Delete(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.Status(http.StatusNoContent)
}

// LoginUser handles POST requests for user authentication.
// Validates credentials and returns JWT token on success.
func LoginUser(c *gin.Context) {
	var loginReq LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	response, err := Login(loginReq.Email, loginReq.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
