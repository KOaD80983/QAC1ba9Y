// 代码生成时间: 2025-08-05 21:42:51
// restful_api.go
package main

import (
    "github.com/gobuffalo/buffalo"
    "github.com/gobuffalo/buffalo/middleware"
    "net/http"
)

// 定义一个简单的用户模型
type User struct {
    ID    uint   `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

// 用户数据存储
var users = []User{
    {ID: 1, Name: "John", Email: "john@example.com"},
    {ID: 2, Name: "Jane", Email: "jane@example.com"},
}

// newUserHandler 创建一个新的用户
func newUserHandler(c buffalo.Context) error {
    user := User{}
    if err := c.Bind(&user); err != nil {
        return err
    }
    users = append(users, user)
    return c.Render(200, r.JSON(user))
}

// listUsersHandler 列出所有用户
func listUsersHandler(c buffalo.Context) error {
    return c.Render(200, r.JSON(users))
}

// showUserHandler 展示单个用户信息
func showUserHandler(c buffalo.Context) error {
    id, _ := c.Param("id")
    idInt, err := strconv.Atoi(id)
    if err != nil {
        return c.Error(404, "User not found")
    }
    for _, user := range users {
        if user.ID == uint(idInt) {
            return c.Render(200, r.JSON(user))
        }
    }
    return c.Error(404, "User not found")
}

// main 函数设置路由并启动服务器
func main() {
    app := buffalo.Automatic()

    // 注册中间件
    app.Use(middleware.Logger)

    // 注册路由
    app.GET("/users", listUsersHandler)
    app.POST("/users", newUserHandler)
    app.GET("/users/{id}", showUserHandler)

    // 启动服务器
    app.Serve()
}
