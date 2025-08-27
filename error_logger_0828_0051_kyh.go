// 代码生成时间: 2025-08-28 00:51:00
package main

import (
    "buffalo"
    "buffalo/buffalo"
    "buffalo/middleware"
    "log"
    "net/http"
)

// ErrorLogger is a middleware that logs errors
func ErrorLogger(next buffalo.Handler) buffalo.Handler {
    return func(c buffalo.Context) error {
        // Call the next handler
        err := next(c)
        // Check if there is an error to log
        if err != nil {
            log.Printf("Error occured: %s
", err.Error())
        }
        return err
    }
}

// App is the main application struct
type App struct {
    *buffalo.App
    errorLog *log.Logger
}

// NewApp creates a new instance of the App
func NewApp() *App {
    a := buffalo.New(buffalo.Options{})
    a.Use(middleware.Logger)
    a.Use(ErrorLogger)
    a.GET("/error", func(c buffalo.Context) error {
        return buffalo.NewError("test error")
    })
    return a
}

// Run the application
func main() {
    app := NewApp()
    app.ErrorLog = log.New(log.Writer(), "ERROR: ", log.Ldate|log.Ltime)
    if err := app.Serve(); err != nil {
        log.Fatal(err)
    }
}
