package main

import (
    "book-management/routers"
    "book-management/database"
    "github.com/gin-gonic/gin"
)

func main() {
    database.InitDB()

    router := gin.Default()

    // 提供静态文件
    router.Static("/static", "./static")

    // 初始化路由
    routers.InitRoutes(router)

    // 运行服务器
    router.Run(":8080")
}
