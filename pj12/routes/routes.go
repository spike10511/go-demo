package routes

import (
	"pj12/controllers"
	"pj12/database"

	"github.com/gin-gonic/gin"
)

func InitRoutes() error {
	if err := database.InitDB(); err != nil {
		return err
	}

	r := gin.Default()

	r.GET("/user/:id", controllers.GetUserByID)
	r.GET("/item/:id", controllers.GetItemByID)

	return r.Run(":8080")
}
