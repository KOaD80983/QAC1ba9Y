// 代码生成时间: 2025-09-15 14:44:34
package main

import (
    "github.com/gobuffalo/buffalo"
    "github.com/gobuffalo/buffalo/generators"
    "github.com/gobuffalo/buffalo/meta/inflect"
    "github.com/gobuffalo/packd"
    "github.com/gobuffalo/packr/v2"
    "github.com/markbates/inflect"
    "log"
)

// UiLibraryApp 定义了BUFFALO应用的结构
type UiLibraryApp struct {
    *buffalo.App
}

// NewUiLibraryApp 创建一个新的UiLibraryApp实例
func NewUiLibraryApp() *UiLibraryApp {
    app := buffalo.NewApp(
        buffalo.Options{
            Environment: "development",
        },
    )

    // 设置静态文件路径
    app.ServeFiles("/assets", packr.New("assets", "./assets"))

    // 定义路由
    app.GET("/", HomeHandler)
    app.GET("/components/{componentName}", ComponentHandler)

    return &UiLibraryApp{App: app}
}

// HomeHandler 处理首页请求
func HomeHandler(c buffalo.Context) error {
    // 返回页面
    return c.Render(200, r.HTML("home/index.html"))
}

// ComponentHandler 处理组件请求
func ComponentHandler(c buffalo.Context) error {
    componentName := c.Param("componentName")
    // 验证组件名称是否有效
    if componentName == "" {
        return c.Error(404, "Component not found")
    }
    // 返回组件页面
    return c.Render(200, r.HTML("components/"+componentName+"/index.html"))
}

func main() {
    // 创建UiLibraryApp实例
    app := NewUiLibraryApp()

    // 启动应用
    if err := app.Serve(); err != nil {
        log.Fatal(err)
    }
}
