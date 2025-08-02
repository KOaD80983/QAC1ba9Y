// 代码生成时间: 2025-08-02 17:54:20
package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "os"
    "strings"
    "log"
    "path/filepath"
)

// Analyzer 结构体用于存储文件分析结果
type Analyzer struct {
    FilePath string
    Results  map[string]int
}

// NewAnalyzer 创建一个新的Analyzer实例
func NewAnalyzer(filePath string) *Analyzer {
    return &Analyzer{
        FilePath: filePath,
        Results:  make(map[string]int),
    }
}

// AnalyzeFile 分析指定的文本文件
func (a *Analyzer) AnalyzeFile() error {
    file, err := os.Open(a.FilePath)
    if err != nil {
        return err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        words := strings.Fields(strings.ToLower(line))
        for _, word := range words {
            // 忽略标点符号和数字
            if strings.ContainsAny(word, `.,;:"'?!-()[]{}<>`) || strings.ContainsAny(word, "0123456789") {
                continue
            }
            a.Results[word]++
        }
    }
    if err := scanner.Err(); err != nil {
        return err
    }
    return nil
}

// PrintResults 打印分析结果
func (a *Analyzer) PrintResults() {
    fmt.Println("Word Frequency Analysis Results: ")
    for word, count := range a.Results {
        fmt.Printf("%s: %d
", word, count)
    }
}

// SaveResults 将分析结果保存到JSON文件
func (a *Analyzer) SaveResults(outputFilePath string) error {
    resultsBytes, err := json.MarshalIndent(a.Results, "", "  ")
    if err != nil {
        return err
    }
    return os.WriteFile(outputFilePath, resultsBytes, 0644)
}

// main 函数是程序入口点
func main() {
    if len(os.Args) < 3 {
        log.Fatal("Usage: text_file_analyzer <input file> <output JSON file>")
    }

    inputFilePath := os.Args[1]
    outputFilePath := os.Args[2]

    analyzer := NewAnalyzer(inputFilePath)
    if err := analyzer.AnalyzeFile(); err != nil {
        log.Fatalf("Error analyzing file: %s", err)
    }
    analyzer.PrintResults()
    if err := analyzer.SaveResults(outputFilePath); err != nil {
        log.Fatalf("Error saving results: %s", err)
    }
}
