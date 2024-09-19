package controllers

import (
    "book-management/models"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

// GetBooks 返回所有书籍
func GetBooks(c *gin.Context) {
    c.JSON(http.StatusOK, models.Books)
}

// AddBook 添加一本书
func AddBook(c *gin.Context) {
    var newBook models.Book
    if err := c.ShouldBindJSON(&newBook); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    newBook.ID = uint(len(models.Books) + 1)
    models.Books = append(models.Books, newBook)
    c.JSON(http.StatusOK, newBook)
}

// UpdateBook 更新书籍信息
func UpdateBook(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))

    for i, book := range models.Books {
        if book.ID == uint(id) {
            if err := c.ShouldBindJSON(&models.Books[i]); err != nil {
                c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
                return
            }
            c.JSON(http.StatusOK, models.Books[i])
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

// DeleteBook 删除一本书
func DeleteBook(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))

    for i, book := range models.Books {
        if book.ID == uint(id) {
            models.Books = append(models.Books[:i], models.Books[i+1:]...)
            c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}
