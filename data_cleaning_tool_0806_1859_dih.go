// 代码生成时间: 2025-08-06 18:59:18
package main

import (
    "buffalo"
    "buffalo/worker"
    "encoding/csv"
    "errors"
    "fmt"
    "os"
    "strings"
)

// DataCleaner 定义数据清洗工具的结构
type DataCleaner struct{}

// NewDataCleaner 创建一个新的数据清洗工具实例
func NewDataCleaner() *DataCleaner {
    return &DataCleaner{}
}

// CleanData 清洗CSV文件中的数据
// @receiver cleaner
// @param filename CSV文件路径
// @param outputFilename 输出文件路径
// @param filterColumns 需要保留的列名列表
// @return error 错误信息
func (cleaner *DataCleaner) CleanData(filename, outputFilename string, filterColumns []string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        return err
    }

    // 检查是否提供了需要保留的列名
    if len(filterColumns) == 0 {
        return errors.New("no columns to filter")
    }

    // 创建新的CSV文件用于存储清洗后的数据
    outputFile, err := os.Create(outputFilename)
    if err != nil {
        return err
    }
    defer outputFile.Close()

    writer := csv.NewWriter(outputFile)
    defer writer.Flush()

    // 写入表头
    for _, column := range filterColumns {
        if err := writer.Write([]string{column}); err != nil {
            return err
        }
    }

    // 遍历记录，清洗数据并写入新文件
    for _, record := range records {
        cleanRecord := []string{}
        for _, column := range filterColumns {
            for i, header := range records[0] {
                if strings.EqualFold(header, column) {
                    cleanRecord = append(cleanRecord, record[i])
                    break
                }
            }
        }
        if err := writer.Write(cleanRecord); err != nil {
            return err
        }
    }

    return nil
}

// main 函数，程序入口点
func main() {
    app := buffalo.App()
    app.Use(buffalo.Worker(func(c buffalo.Context) error {
        // 创建数据清洗工具实例
        cleaner := NewDataCleaner()

        // 调用数据清洗函数
        if err := cleaner.CleanData("input.csv", "output.csv", []string{"column1", "column2"}); err != nil {
            return c.Error(err, 500)
        }

        // 返回成功消息
        return c.Render(200, buffalo.Text("You have successfully cleaned the data!"))
    }))

    app.Start()
}
