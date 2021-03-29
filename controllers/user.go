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

// UpdateUserInput - the types accpeted for user creation
type UpdateUserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// FindUsers - search the database for all user entries
// GET /users
func FindUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// FindUser - find a single user from the database based on ID
// GET /users/:id
func FindUser(c *gin.Context) {
	var user models.User

	err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// CreateUser - create a new user and add to the database
// POST /users
func CreateUser(c *gin.Context) {
	var input CreateUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user := models.User{Name: input.Name, Email: &input.Email}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// UpdateUser - update a current user in the database
// PATH /users/:id
func UpdateUser(c *gin.Context) {
	var user models.User
	err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateUserInput
	err = c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&user).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// DeleteUser - delete the user from the database
// DELETE /users/:id
func DeleteUser(c *gin.Context) {
	var user models.User
	err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
