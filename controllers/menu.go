package controllers

import (
	"goticle/config"
	"goticle/models"
	"goticle/requests"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateMenu(c *gin.Context) {
	var input requests.MenuInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	menu := models.Menu{
		Name:     input.Name,
		Url:      input.Url,
		Icon:     input.Icon,
		Order:    input.Order,
		ParentID: input.ParentID,
	}

	config.DB.Create(&menu)

	c.JSON(http.StatusOK, gin.H{"data": menu})
}

func GetMenus(c *gin.Context) {
	var menus []models.Menu

	config.DB.Preload("Childs").Where("parent_id", nil).Find(&menus)

	c.JSON(http.StatusOK, gin.H{"data": menus})
}

func FindMenu(c *gin.Context) {
	var menu []models.Menu

	if err := config.DB.Preload("Childs").Where("parent_id", nil).Where("id = ?", c.Param("id")).First(&menu).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Menu not found"})

		return
	}

	c.JSON(http.StatusOK, gin.H{"data": menu})
}

func UpdateMenu(c *gin.Context) {
	var menu models.Menu

	if err := config.DB.Where("id = ?", c.Param("id")).First(&menu).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Menu not found"})

		return
	}

	var input requests.MenuInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	config.DB.Model(menu).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": menu})
}

func DeleteMenu(c *gin.Context) {
	var menu models.Menu

	if err := config.DB.Where("id = ?", c.Param("id")).First(&menu).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Menu not found"})
	}

	config.DB.Delete(&menu)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func ForceDeleteMenu(c *gin.Context) {
	var menu models.Menu

	if err := config.DB.Unscoped().Where("id = ?", c.Param("id")).First(&menu).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Menu not found"})
	}

	config.DB.Delete(&menu)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
