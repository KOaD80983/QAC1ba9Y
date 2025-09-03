// 代码生成时间: 2025-09-04 03:56:27
package main

import (
    "net"
    "os"
    "log"
    "github.com/gobuffalo/buffalo"
    "context"
)

// NetworkChecker struct contains methods to check network connectivity
type NetworkChecker struct {
    // No fields are required for this basic example
}

// NewNetworkChecker creates a new NetworkChecker instance
func NewNetworkChecker() *NetworkChecker {
    return &NetworkChecker{}
}

// CheckConnectivity checks if the server is reachable
func (nc *NetworkChecker) CheckConnectivity(host string, port int) error {
    conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, strconv.Itoa(port)), 5e9)
    if err != nil {
        return err
    }
    defer conn.Close()
    return nil // No error means the host is reachable
}

// App buffalo application instance
var App buffalo.App

func main() {
    App = buffalo.Automatic(buffalo.Options{
        Env:         buffalo.Env("development"),
        Logger:      log.New(os.Stdout),
        SessionStore: buffalo.SessionStore{},
    })

    // Define route for checking network connectivity
    App.GET("/check", func(c buffalo.Context) error {
        // Extract host and port from query parameters
        host := c.Request().URL.Query().Get("host")
        portStr := c.Request().URL.Query().Get("port")
        port, err := strconv.Atoi(portStr)
        if err != nil {
            return c.Error(400, err)
        }

        // Check network connectivity using the NetworkChecker
        nc := NewNetworkChecker()
        if err := nc.CheckConnectivity(host, port); err != nil {
            return c.Error(503, err)
        }

        return c.Render(200, buffalo.JSON(map[string]string{
            "status": "reachable",
        }))
    })

    if err := App.Serve(); err != nil {
        log.Fatal(err)
    }
}
