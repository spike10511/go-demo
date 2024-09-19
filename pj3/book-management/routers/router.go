package routers

import (
    "book-management/controllers"
    "github.com/gin-gonic/gin"
)

// InitRoutes 初始化路由
func InitRoutes(router *gin.Engine) {
    // 定义书籍管理相关的路由
    router.GET("/books", controllers.GetBooks)
    router.POST("/books", controllers.AddBook)
    router.PUT("/books/:id", controllers.UpdateBook)
    router.DELETE("/books/:id", controllers.DeleteBook)
}
