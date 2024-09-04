package controllers

import (
	"github.com/Trend20/go-shopper-api/config"
	"github.com/Trend20/go-shopper-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// creating a user
func CreateUser(c *gin.Context) {
	var userInput models.CreateUserInput
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.User{Name: userInput.Name, Email: userInput.Email, Pass: userInput.Pass, Role: userInput.Role}
	config.DB.Create(&user)
	c.JSON(http.StatusCreated, gin.H{"user": user})
}

// getting all users
func GetAllUsers(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}

// getting a single user by ID
func GetUser(c *gin.Context) {
	var user models.User
	err := config.DB.Where("id = ?", c.Param("id")).First(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// updating user details
func UpdateUser(c *gin.Context) {
	//get the model if it exists
	var user models.User
	if err := config.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	//validate the input
	var updateInput models.UpdateUserInput
	if err := c.ShouldBindJSON(&updateInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Model(&user).Where("id = ?", c.Param("id")).Updates(&updateInput)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// deleting a user
func DeleteUser(c *gin.Context) {
	var user models.User
	if err := config.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	config.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
