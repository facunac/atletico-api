package controllers

import (
	"net/http"
	"time"

	"atletico-api/models"
	"github.com/gin-gonic/gin"
)

type CreateUserInput struct {
	Email       string    `json:"email" binding:"required"`
	Rut         string    `json:"rut" binding:"required"`
	First_name  string    `json:"first_name" binding:"required"`
	Last_name   string    `json:"last_name" binding:"required"`
	Password    string    `json:"password" binding:"required"`
	Phone       string    `json:"phone" binding:"required"`
	Birthdate   time.Time `json:"birthdate" binding:"required"`
	Gender      string    `json:"gender" binding:"required"`
	Nationality string    `json:"nationality" binding:"required"`
	Graduate    uint      `json:"graduate" binding:"required"`
	Type        string    `json:"type" binding:"required"`
}

type UpdateUserInput struct {
	Email       string    `json:"email" binding:"required"`
	Rut         string    `json:"rut" binding:"required"`
	First_name  string    `json:"first_name" binding:"required"`
	Last_name   string    `json:"last_name" binding:"required"`
	Password    string    `json:"password" binding:"required"`
	Phone       string    `json:"phone" binding:"required"`
	Birthdate   time.Time `json:"birthdate" binding:"required"`
	Gender      string    `json:"gender" binding:"required"`
	Nationality string    `json:"nationality" binding:"required"`
	Graduate    uint      `json:"graduate" binding:"required"`
	Type        string    `json:"type" binding:"required"`
}

type LoginUserInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// GET /users
// Find all users
func FindUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// GET /users/:id
// Find a user
func FindUser(c *gin.Context) {
	// Get model if exist
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// POST /users
// Create new user
func CreateUser(c *gin.Context) {
	// Validate input
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create user
	user := models.User{
		Email:       input.Email,
		Rut:         input.Rut,
		First_name:  input.First_name,
		Last_name:   input.Last_name,
		Password:    input.Password,
		Phone:       input.Phone,
		Birthdate:   input.Birthdate,
		Gender:      input.Gender,
		Nationality: input.Nationality,
		Graduate:    input.Graduate,
		Typeuser:    input.Type,
		Created_on:  time.Now(),
		Last_login:  time.Now()}

	err := models.DB.Create(&user)
	if err.Error == nil {
		c.JSON(http.StatusOK, gin.H{"data": user})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": &err.Error})
	}
}

// PATCH /users/:id
// Update a user
func UpdateUser(c *gin.Context) {

	// Validate input
	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	//var user models.User
	user := models.User{
		Email:       input.Email,
		Rut:         input.Rut,
		First_name:  input.First_name,
		Last_name:   input.Last_name,
		Password:    input.Password,
		Phone:       input.Phone,
		Birthdate:   input.Birthdate,
		Gender:      input.Gender,
		Nationality: input.Nationality,
		Graduate:    input.Graduate,
		Typeuser:    input.Type,
		Created_on:  time.Now(),
		Last_login:  time.Now()}

	if err := models.DB.Where("ID = ?", c.Param("id")).Save(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": user})
	}
}

// DELETE /users/:id
// Delete a user
func DeleteUser(c *gin.Context) {
	// Get model if exist
	var user models.User

	err := models.DB.Where("ID = ?", c.Param("id")).Delete(&user)
	if err.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": true})
	}

}

// Login /users/login
// Login a user
func LoginUser(c *gin.Context) {

	// Validate input
	var input LoginUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	if err := models.DB.Where("email = ? and password = ?", input.Email, input.Password).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "usuario o clave incorrecto"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})

}
