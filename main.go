package main

import (
	"goticle/config"
	"goticle/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	v1 := r.Group("/api/v1")

	config.ConnectionDatabase()

	v1.POST("/register", controllers.Register)
	v1.POST("/login", controllers.Login)
	v1.GET("/users", controllers.GetUsers)
	v1.GET("/users/:id", controllers.Finduser)
	v1.PUT("/users/:id", controllers.UpdateUser)
	v1.DELETE("/users/:id", controllers.DeleteUser)
	v1.DELETE("/users/:id/force", controllers.ForceDeleteUser)

	v1.POST("/menus", controllers.CreateMenu)
	v1.GET("/menus", controllers.GetMenus)
	v1.GET("/menus/:id", controllers.FindMenu)
	v1.PUT("/menus/:id", controllers.UpdateMenu)
	v1.DELETE("/menus/:id", controllers.DeleteMenu)
	v1.DELETE("/menus/:id/force", controllers.ForceDeleteMenu)

	v1.POST("/roles", controllers.CreateRole)
	v1.GET("/roles", controllers.GetRoles)
	v1.GET("/roles/:id", controllers.FindRole)
	v1.PUT("/roles/:id", controllers.UpdateRole)
	v1.DELETE("/roles/:id", controllers.DeleteRole)
	v1.DELETE("/roles/:id/force", controllers.ForceDeleteRole)

	v1.POST("/permissions", controllers.CreatePermission)
	v1.GET("/permissions", controllers.GetPermissions)
	v1.GET("/permissions/:id", controllers.FindPermission)
	v1.PUT("/permissions/:id", controllers.UpdatePermission)
	v1.DELETE("/permissions/:id", controllers.DeletePermission)
	v1.DELETE("/permissions/:id/force", controllers.ForceDeletePermission)

	r.Run(":1000")
}
