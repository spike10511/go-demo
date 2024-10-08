# 使用 Go 和 Gin 框架构建简单的用户和物品管理 Web 服务

在本项目中，我们使用 Go 语言和 Gin 框架构建了一个简单的 Web 服务，能够管理用户和物品的信息。该服务实现了两个主要接口：根据用户 ID 获取用户名称，以及根据物品 ID 获取物品名称。本文将介绍项目的整体结构、数据库设计、接口实现以及如何测试这些接口。

## 项目结构

项目的目录结构如下：
![图片](https://github.com/user-attachments/assets/8a2e8635-5eab-4486-850c-6d3adc94701b)

## 数据库设计

本项目使用 MySQL 数据库存储用户和物品信息。我们创建了两张表：

![图片](https://github.com/user-attachments/assets/cc9b456a-2887-4615-8336-2d2917c02daf)

### 示例数据

我们在数据库中插入了一些示例用户和物品数据：

```sql
-- 插入用户数据
INSERT INTO users (name) VALUES ('张三'); -- 用户ID 1
INSERT INTO users (name) VALUES ('李四'); -- 用户ID 2

-- 插入物品数据
INSERT INTO items (name, user_id) VALUES ('书籍', 1);  -- 张三的书籍
INSERT INTO items (name, user_id) VALUES ('电脑', 2);  -- 李四的电脑
INSERT INTO items (name, user_id) VALUES ('手机', 1);   -- 张三的手机
```

## 接口实现
1.获取用户名称
// GetUserByID 根据用户ID获取用户名称
```go
func GetUserByID(c *gin.Context) {
    id := c.Param("id") // 从请求中获取用户ID
    var user models.User

    err := database.DB.QueryRow("SELECT id, name FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name)
    
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
```
2.获取物品名称
```go
// GetItemByID 根据物品ID获取物品信息
func GetItemByID(c *gin.Context) {
    id := c.Param("id") // 从请求中获取物品ID
    var item models.Item

    err := database.DB.QueryRow("SELECT id, name, user_id FROM items WHERE id = ?", id).Scan(&item.ID, &item.Name, &item.UserID)
    
    if err != nil {
        if err == sql.ErrNoRows {
            c.JSON(http.StatusNotFound, gin.H{"message": "物品未找到"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"message": "查询物品时出错"})
        }
        return
    }

    // 查询物品所属用户的信息
    var user models.User
    err = database.DB.QueryRow("SELECT id, name FROM users WHERE id = ?", item.UserID).Scan(&user.ID, &user.Name)
    if err == nil {
        fmt.Println("用户名称:", user.Name) // 模拟打印日志
    }

    c.JSON(http.StatusOK, item)
}
```
3.路由设置
我们将上述接口绑定到路由中：
```go
func SetupRouter() *gin.Engine {
    router := gin.Default()

    // 用户相关接口
    router.GET("/user/:id", controllers.GetUserByID) // 根据用户ID获取用户名称

    // 物品相关接口
    router.GET("/item/:id", controllers.GetItemByID) // 根据物品ID获取物品名称

    return router
}
```

## 测试接口
使用 Postman 来测试接口:
获取用户名称：GET http://localhost:8080/user/1  # 查询用户ID为1（张三）
查询结果：![图片](https://github.com/user-attachments/assets/08bff73b-ba9e-49fb-a6b9-80a34dce6fb4)
获取物品名称：GET http://localhost:8080/item/2  # 查询物品ID为2（电脑）
查询结果：![图片](https://github.com/user-attachments/assets/e0ca2688-fdce-488c-ab23-b1f9289f0d69)
