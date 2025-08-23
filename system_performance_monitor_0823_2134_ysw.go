// 代码生成时间: 2025-08-23 21:34:37
package main

import (
    "github.com/gobuffalo/buffalo"
    "github.com/gobuffalo/buffalo/middleware"
    "github.com/gobuffalo/buffalo/worker"
    "github.com/shirou/gopsutil/cpu"
    "github.com/shirou/gopsutil/disk"
    "github.com/shirou/gopsutil/mem"
    "github.com/shirou/gopsutil/net"
    "github.com/markbates/pkger"
)

// SystemPerformanceMonitor defines the structure for our application.
type SystemPerformanceMonitor struct{
}

// NewSystemPerformanceMonitor creates a new instance of SystemPerformanceMonitor.
func NewSystemPerformanceMonitor() *SystemPerformanceMonitor {
    return &SystemPerformanceMonitor{}
}

// SystemMetrics contains the metrics for system performance.
type SystemMetrics struct {
    CPUUsage float64 `json:"cpu_usage"`
    RAMUsage float64 `json:"ram_usage"`
    DiskUsage float64 `json:"disk_usage"`
    NetworkTraffic struct {
        Sent float64 `json:"sent"`
        Received float64 `json:"received"`
    } `json:"network_traffic"`
}

// GetSystemMetrics fetches system metrics and returns them as SystemMetrics.
func (spm *SystemPerformanceMonitor) GetSystemMetrics(c buffalo.Context) error {
    cpuPercentages, err := cpu.Percent(0, false)
    if err != nil {
        return err
    }
    totalRAM, err := mem.Total()
    if err != nil {
        return err
    }
    usedRAM, err := mem.Used()
    if err != nil {
        return err
    }
    diskUsage, err := disk.Usage="/"
    if err != nil {
        return err
    }
    netStats, err := net.IOCounters(true)
    if err != nil {
        return err
    }

    return c.Render(200, buffalo.JSON(&SystemMetrics{
        CPUUsage: cpuPercentages[0],
        RAMUsage: float64(usedRAM) / float64(totalRAM),
        DiskUsage: diskUsage.UsedPercent,
        NetworkTraffic: struct {
            Sent float64
            Received float64
        }{
            Sent: netStats.BytesSent,
            Received: netStats.BytesRecv,
        },
    }))
}

// main is the entry point for the Buffalo application.
func main() {
    app := buffalo.Automatic(buffalo.Options{
        AppName: "SystemPerformanceMonitor",
    })

    // Automatically generate routers that match the URL to the appropriate handler.
    app.GET("/system-metrics", func(c buffalo.Context) error {
        return NewSystemPerformanceMonitor().GetSystemMetrics(c)
    })

    // Add middleware to handle CORS, security, and other cross-cutting concerns.
    app.Use(middleware.CSRF)
    app.Use(middleware.Security)
    app.Use(middleware.Static("assets"))

    // Run the application.
    if err := app.Serve(); err != nil {
        app.Stop(err)
    }
}
