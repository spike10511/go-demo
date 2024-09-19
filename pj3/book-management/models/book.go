package models

// Book 代表一本书的结构体
type Book struct {
    ID     uint   `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
    ISBN   string `json:"isbn"`
}

// 用于存储书籍数据的全局变量
var Books []Book
