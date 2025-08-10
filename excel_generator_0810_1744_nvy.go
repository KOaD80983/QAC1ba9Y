// 代码生成时间: 2025-08-10 17:44:07
package main

import (
    "os"
    "log"
    "encoding/csv"
    "github.com/tealeg/xlsx/v3"
    "github.com/gobuffalo/buffalo"
)

// ExcelGenerator represents the struct that will handle Excel file generation
type ExcelGenerator struct {
    SheetName string
    FileName  string
}

// NewExcelGenerator creates a new instance of ExcelGenerator
func NewExcelGenerator(sheetName, fileName string) *ExcelGenerator {
    return &ExcelGenerator{
        SheetName: sheetName,
        FileName:  fileName,
    }
}

// Generate generates an Excel file with the given data
func (e *ExcelGenerator) Generate(data [][]string) error {
    // Create a new Excel file
    file := xlsx.NewFile()
    sheet, err := file.AddSheet(e.SheetName)
    if err != nil {
        return err
    }

    // Write data to the sheet
    for _, row := range data {
        if err := sheet.AddRow(row); err != nil {
            return err
        }
    }

    // Save the Excel file
    f, err := os.Create(e.FileName)
    if err != nil {
        return err
    }
    defer f.Close()
    if err := file.Write(f); err != nil {
        return err
    }
    return nil
}

// App is the Buffalo application
type App struct{}

// Actions returns the Buffalo actions
func (a *App) Actions() buffalo.Actions {
    return buffalo.Actions{
        buffalo.Get("/generate", a.generateExcel),
    }
}

// generateExcel handles the HTTP request to generate an Excel file
func (a *App) generateExcel(c buffalo.Context) error {
    // Sample data to be written to the Excel file
    data := [][]string{
        {"Header1", "Header2", "Header3"},
        {"Data1", "Data2", "Data3"},
        {"Data4", "Data5", "Data6"},
    }

    // Create a new ExcelGenerator instance
    generator := NewExcelGenerator("Sheet1", "example.xlsx")

    // Generate the Excel file
    if err := generator.Generate(data); err != nil {
        log.Printf("Error generating Excel file: %v", err)
        return c.Error(500, err)
    }

    // Write the Excel file to the HTTP response
    f, err := os.Open("example.xlsx")
    if err != nil {
        log.Printf("Error opening Excel file: %v", err)
        return c.Error(500, err)
    }
    defer f.Close()

    c.Response().SetContentType("application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
    c.Response().WriteHeader(200)
    _, err = c.Response().Write(f.Bytes())
    if err != nil {
        log.Printf("Error writing Excel file to response: %v", err)
        return c.Error(500, err)
    }

    return nil
}

// main is the entry point of the application
func main() {
    // Create a new instance of the App
    app := &App{}

    // Start the Buffalo application
    if err := buffalo.buffalo(buffalo.Options{
        AppName: "ExcelGenerator",
        Action: app,
    }); err != nil {
        log.Fatal(err)
    }
}