package controllers

import (
	"goticle/config"
	"goticle/models"
	"goticle/requests"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// type LoginInput struct {
// 	Username string `json:"username" binding:"required"`
// 	Password string `json:"password" binding:"required"`
// }

func Register(c *gin.Context) {
	var input requests.RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user := models.User{}
	user.Username = input.Username
	user.Name = input.Name
	user.Email = input.Email
	user.Password = input.Email

	hashing, hashErr := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if hashErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": hashErr.Error()})
	}

	user.Password = string(hashing)

	config.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func Login(c *gin.Context) {
	var input requests.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	user := models.User{}

	if err := config.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func GetUsers(c *gin.Context) {
	var users []models.User
	config.DB.Preload("Role.Permissions").Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func Finduser(c *gin.Context) {
	var user models.User

	if err := config.DB.Where("id = ?", c.Param("id")).Preload("Role.Permissions").First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})

		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func UpdateUser(c *gin.Context) {
	var user models.User

	if err := config.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})

		return
	}

	var input requests.UserUpdateInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	config.DB.Model(user).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": input})
}

func DeleteUser(c *gin.Context) {
	var user models.User

	if err := config.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})

		return
	}

	config.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func ForceDeleteUser(c *gin.Context) {
	var user models.User

	if err := config.DB.Unscoped().Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})

		return
	}

	config.DB.Unscoped().Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
