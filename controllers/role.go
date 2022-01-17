package controllers

import (
	"goticle/config"
	"goticle/models"
	"goticle/requests"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRole(c *gin.Context) {
	var input requests.RoleInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	role := models.Role{
		Name:        input.Name,
		Description: input.Description,
		Active:      input.Active,
	}

	config.DB.Create(&role)

	c.JSON(http.StatusOK, gin.H{"data": role})
}

func GetRoles(c *gin.Context) {
	var roles []models.Role

	config.DB.Preload("Permissions").Find(&roles)

	c.JSON(http.StatusOK, gin.H{"data": roles})
}

func FindRole(c *gin.Context) {
	var role models.Role

	if err := config.DB.Where("id = ?", c.Param("id")).Preload("Permissions").First(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Role not found"})

		return
	}

	c.JSON(http.StatusOK, gin.H{"data": role})
}

func UpdateRole(c *gin.Context) {
	var role models.Role

	if err := config.DB.Where("id = ?", c.Param("id")).First(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Role not found"})

		return
	}

	var input requests.RoleInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	config.DB.Model(role).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": role})
}

func DeleteRole(c *gin.Context) {
	var role models.Role

	if err := config.DB.Where("id = ?", c.Param("id")).First(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Role not found"})

		return
	}

	config.DB.Delete(&role)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func ForceDeleteRole(c *gin.Context) {
	var role models.Role

	if err := config.DB.Unscoped().Where("id = ?", c.Param("id")).First(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Role not found"})

		return
	}

	config.DB.Unscoped().Delete(&role)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
