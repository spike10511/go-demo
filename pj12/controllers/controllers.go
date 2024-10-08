package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"pj12/database"
	"pj12/models"

	"github.com/gin-gonic/gin"
)

// GetUserByID 根据用户ID获取用户信息
func GetUserByID(c *gin.Context) {
	userID := c.Param("id")
	var user models.User
	err := database.DB.QueryRow("SELECT id, name FROM users WHERE id = ?", userID).Scan(&user.ID, &user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": "用户未找到"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "查询用户时出错"})
		}
		return
	}
	c.JSON(http.StatusOK, user)
}

// GetItemByID 根据物品ID获取物品信息
func GetItemByID(c *gin.Context) {
	itemID := c.Param("id")
	var item models.Item
	err := database.DB.QueryRow("SELECT id, name, user_id FROM items WHERE id = ?", itemID).Scan(&item.ID, &item.Name, &item.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": "物品未找到"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "查询物品时出错"})
		}
		return
	}

	// 在查询物品时，发出对用户查询的接口请求
	var user models.User
	err = database.DB.QueryRow("SELECT id, name FROM users WHERE id = ?", item.UserID).Scan(&user.ID, &user.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "查询物品所属用户时出错"})
		return
	}

	// 打印用户名称到日志
	fmt.Println("物品所属用户名称:", user.Name)

	// 返回物品和用户信息
	c.JSON(http.StatusOK, gin.H{
		"item": item,
		"user": user,
	})
}
