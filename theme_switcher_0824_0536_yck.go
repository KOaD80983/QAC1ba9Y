// 代码生成时间: 2025-08-24 05:36:03
package main

import (
    "github.com/gobuffalo/buffalo"
    "github.com/gobuffalo/buffalo/middleware"
    "github.com/markbates/pkger"
    "github.com/markbates/pkger/middleware"
)

// ThemeSwitcher is a buffalo application that allows theme switching.
func main() {
    app := buffalo.Automatic()

    // Set up middleware
    app.Use(middleware.CSRF)
    app.Use(middleware.SetTrustedProxies)
    app.Use(pkger.New(&pkger.Options{
        Path: "/assets",
        CachePath: "/pkged",
    }))

    // Define routes
    app.GET("/", HomeHandler)
    app.POST("/", SetThemeHandler)

    // Start the application
    app.Serve()
}

// HomeHandler handles the home page view.
func HomeHandler(c buffalo.Context) error {
    return c.Render(200, r.HTML("home.html"))
}

// SetThemeHandler handles setting the theme for the application.
func SetThemeHandler(c buffalo.Context) error {
    theme := c.Request().URL.Query().Get("theme")
    if theme == "" {
        return c.Error(400, "Theme parameter is required")
    }
    c.Session().Set("theme", theme)
    return c.Redirect(302, "/")
}
