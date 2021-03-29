package controllers

import (
	"net/http"

	"github.com/MilsonCodes/gin-template/models"
	"github.com/gin-gonic/gin"
)

// CreateUserInput - the types accpeted for user creation
type CreateUserInput struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

// FindUsers - search the database for all user entries
// GET /users
func FindUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// CreateUser - create a new user and add to the database
// POST /users
func CreateUser(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user := models.User{Name: input.Name, Email: &input.Email}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}
