package main

import (
	"log"
	"user-management/controllers"
	"user-management/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB() {
	dsn := "root:P110040593abc@tcp(127.0.0.1:3306)/职工管理系统?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	models.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 自动迁移，创建用户表
	err = models.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database connected and migrated successfully!")
}

func main() {
	initDB()

	r := gin.Default()

	// 用户注册路由
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	// 添加到 main.go 中
	r.GET("/users/:id", controllers.GetUserInfo)
	r.PUT("/users/:id", controllers.UpdateUser)

	// 启动服务器
	r.Run(":8080")
}
