// 代码生成时间: 2025-09-10 18:56:54
// excel_generator_app.go
// 这是一个使用GOLANG和BUFFALO框架创建的Excel表格自动生成器程序。

package main

import (
    "os"
    "github.com/gobuffalo/buffalo"
    "github.com/gobuffalo/buffalo-cli/v2/cli"
    "github.com/gobuffalo/packr/v2"
    "github.com/go-excel-tools/excelize/v2"
    "net/http"
)

// App 是BUFFALO应用的入口点
var App = buffalo.Automatic(buffalo.Options{
    PreWarmed:  make(map[string]buffalo.PreWarmedRenderer),
})

// Box 是packr.Box的实例，用于管理静态文件和模板
var Box = packr.New("app", ".")

// HomeHandler 是一个处理GET请求的handler，用于生成Excel文件
func HomeHandler(c buffalo.Context) error {
    var err error
    // 创建一个新的Excel文件
    f := excelize.NewFile()
    
    // 设置Excel文件的标题
    title := "My Excel Sheet"
    f.SetSheetName(0, title)
    
    // 创建一个数据表，这里只是示例数据
    data := [][]string{
        {"Name", "Age", "Email"},
        {"John Doe", "30", "john@example.com"},
        {"Jane Doe", "25", "jane@example.com"},
    }
    
    // 将数据写入Excel文件
    for i, row := range data {
        for j, value := range row {
            f.SetCellValue(title, excelize.CoordinatesForCell(i+1, j+1), value)
        }
    }
    
    // 生成Excel文件
    file, err := f.WriteToBytes()
    if err != nil {
        return handleError(c, err)
    }
    
    // 设置响应头，使浏览器下载Excel文件
    c.Response().Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
    c.Response().Header().Set("Content-Disposition", "attachment; filename="my_excel_file.xlsx"
)
    _, err = c.Response().Write(file)
    if err != nil {
        return handleError(c, err)
    }
    
    return nil
}

// handleError 是一个辅助函数，用于处理错误并返回HTTP错误响应
func handleError(c buffalo.Context, err error) error {
    return c.Error(http.StatusInternalServerError, "An error occurred: %s", err)
}

// main 是程序的入口点，设置路由并启动BUFFALO应用
func main() {
    // 设置路由
    App.GET("/", HomeHandler)
    
    // 启动应用，监听8080端口
    if err := cli.Start(App, 8080); err != nil {
        os.Exit(1)
    }
}
