package controllers

import (
	"goticle/config"
	"goticle/models"
	"goticle/requests"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePermission(c *gin.Context) {
	var input requests.PermissionInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	permission := models.Permission{
		Name:        input.Name,
		Description: input.Description,
		Active:      input.Active,
	}

	config.DB.Create(&permission)

	role_permission := models.RolePermission{
		RoleID:       input.RoleID,
		PermissionID: permission.ID,
	}

	config.DB.Create(&role_permission)

	c.JSON(http.StatusOK, gin.H{"data": permission})
}

func GetPermissions(c *gin.Context) {
	var permission []models.Permission

	config.DB.Find(&permission)

	c.JSON(http.StatusOK, gin.H{"data": permission})
}

func FindPermission(c *gin.Context) {
	var permission models.Permission

	if err := config.DB.Where("id = ?", c.Param("id")).First(&permission).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Permission not found"})

		return
	}

	c.JSON(http.StatusOK, gin.H{"data": permission})
}

func UpdatePermission(c *gin.Context) {
	var permission models.Permission

	if err := config.DB.Where("id = ?", c.Param("id")).First(&permission).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Permission not found"})

		return
	}

	var input requests.PermissionInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	config.DB.Model(permission).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": permission})
}

func DeletePermission(c *gin.Context) {
	var permission models.Permission

	if err := config.DB.Where("id = ?", c.Param("id")).First(&permission).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Permission not found"})

		return
	}

	config.DB.Delete(&permission)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func ForceDeletePermission(c *gin.Context) {
	var permission models.Permission

	if err := config.DB.Unscoped().Where("id = ?", c.Param("id")).First(&permission).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Permission not found"})

		return
	}

	config.DB.Unscoped().Delete(&permission)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
