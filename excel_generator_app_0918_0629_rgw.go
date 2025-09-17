// 代码生成时间: 2025-09-18 06:29:08
package main

import (
    "buffalo.fi"
    "github.com/markbates/buffalo/render"
    "gorm.io/gorm"
    "log"
    "os"
    "github.com/xuri/excelize/v2"
)

var db *gorm.DB

// ExcelGeneratorApp 应用结构体
type ExcelGeneratorApp struct {
    *buffalo.App
    Renderer render.Renderer
}

// NewExcelGeneratorApp 创建一个新的ExcelGeneratorApp实例
func NewExcelGeneratorApp(db *gorm.DB) *ExcelGeneratorApp {
    a := buffalo.New(buffalo.Options{
        PreWares: []buffalo.PreWare{
            buffalo.WrapHandlerWith(buffaloLogger),
        },
    })
    app := &ExcelGeneratorApp{
        App: a,
        Renderer: render.New(render.Options{}),
    }
    
    // 设置路由
    app.GET("/", app.RootHandler)
    app.POST("/generate", app.GenerateExcel)
    
    return app
}

// RootHandler 根路由处理器
func (a *ExcelGeneratorApp) RootHandler(c buffalo.Context) error {
    return c.Render(200, render.HTML{"index.html"})
}

// GenerateExcel 处理生成Excel表格的请求
func (a *ExcelGeneratorApp) GenerateExcel(c buffalo.Context) error {
    // 获取请求参数
    title := c.Request().FormValue("title")
    columns := c.Request().FormValue("columns")
    data := c.Request().FormValue("data")
    
    if title == "" || columns == "" || data == "" {
        return c.Render(400, render.String{"缺少必要的参数"})
    }
    
    // 解析列和数据
    cols := parseColumns(columns)
    rowData := parseData(data)
    
    // 创建Excel文件
    f := excelize.NewFile()
    sheetName := "Sheet1"
    f.NewSheet(sheetName)
    
    // 设置标题行
    f.SetCellValue(sheetName, "A1", title)
    
    // 设置列头
    for i, col := range cols {
        f.SetCellValue(sheetName, excelize.ToAlphaString(i+2)+"1", col)
    }
    
    // 填写数据行
    for i, row := range rowData {
        for j, col := range row {
            f.SetCellValue(sheetName, excelize.ToAlphaString(j+2)+strconv.Itoa(i+2), col)
        }
    }
    
    // 保存Excel文件
    f.SaveAs("./output.xlsx")
    
    return c.Render(200, render.String{"Excel文件已生成"})
}

// parseColumns 解析列参数，转换为列名数组
func parseColumns(columns string) []string {
    // 假设列参数以逗号分隔
    return strings.Split(columns, ",")
}

// parseData 解析数据参数，转换为数据行数组
func parseData(data string) [][]string {
    // 假设数据参数以换行符分隔，每行以逗号分隔
    lines := strings.Split(data, "
")
    var rowData [][]string
    for _, line := range lines {
        rowData = append(rowData, strings.Split(line, ","))
    }
    return rowData
}

func main() {
    db = setupDB()
    defer db.Close()
    
    app := NewExcelGeneratorApp(db)
    app.Serve()
}

// setupDB 设置数据库连接（示例，实际应根据项目需求配置）
func setupDB() *gorm.DB {
    dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8&parseTime=True&loc=Local"
    db, err := gorm.Open(dsn)
    if err != nil {
        log.Fatal(err)
    }
    return db
}