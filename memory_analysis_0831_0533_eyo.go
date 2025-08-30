// 代码生成时间: 2025-08-31 05:33:26
package main

import (
    "bytes"
    "fmt"
    "os"
    "runtime"
    "runtime/pprof"
    "strings"
)

// MemoryAnalysis 结构体用于存储内存分析相关的信息
type MemoryAnalysis struct {
    // 无额外字段
}

// NewMemoryAnalysis 创建一个新的内存分析实例
func NewMemoryAnalysis() *MemoryAnalysis {
    return &MemoryAnalysis{}
}

// StartCPUProfile 开始CPU分析并保存到指定文件
func (m *MemoryAnalysis) StartCPUProfile(filename string) error {
    f, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer f.Close()
    if err := pprof.StartCPUProfile(f); err != nil {
        return err
    }
    return nil
}

// StopCPUProfile 停止CPU分析
func (m *MemoryAnalysis) StopCPUProfile() {
    pprof.StopCPUProfile()
}

// WriteHeapProfile 写入堆内存分析到指定文件
func (m *MemoryAnalysis) WriteHeapProfile(filename string) error {
    f, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer f.Close()
    if err := pprof.WriteHeapProfile(f); err != nil {
        return err
    }
    return nil
}

// PrintMemoryUsage 打印当前内存使用情况
func (m *MemoryAnalysis) PrintMemoryUsage() {
    var b bytes.Buffer
    runtime.GC() // 触发GC以获取准确的内存使用情况
    memStats := &runtime.MemStats{}
    runtime.ReadMemStats(memStats)
    fmt.Fprintf(&b, "Alloc = %v
", memStats.Alloc)
    fmt.Fprintf(&b, "TotalAlloc = %v
", memStats.TotalAlloc)
    fmt.Fprintf(&b, "Sys = %v
", memStats.Sys)
    fmt.Fprintf(&b, "Mallocs = %v
", memStats.Mallocs)
    fmt.Fprintf(&b, "Frees = %v
", memStats.Frees)
    fmt.Fprintf(&b, "LiveObjects = %v
", memStats.Mallocs-memStats.Frees)
    fmt.Fprintf(&b, "Pauses = %v
", formatPauses(memStats.PauseTotalNs))
    fmt.Println(b.String())
}

// formatPauses 格式化暂停时间
func formatPauses(pauseTotalNs uint64) string {
    pauseTotal := float64(pauseTotalNs) / 1e6 // 将纳秒转换为毫秒
    return fmt.Sprintf("%v ms", pauseTotal)
}

func main() {
    analysis := NewMemoryAnalysis()
    
    // 开始CPU分析
    if err := analysis.StartCPUProfile("cpu.prof"); err != nil {
        fmt.Printf("Failed to start CPU profile: %v
", err)
        return
    }
    defer analysis.StopCPUProfile()
    
    // 执行一些操作以产生内存使用
    // 这里为了示例，我们只是简单地创建一个大的字符串
    largeString := strings.Repeat("Hello, World!", 10000)
    
    // 打印内存使用情况
    analysis.PrintMemoryUsage()
    
    // 写入堆内存分析
    if err := analysis.WriteHeapProfile("heap.prof"); err != nil {
        fmt.Printf("Failed to write heap profile: %v
", err)
        return
    }
}
