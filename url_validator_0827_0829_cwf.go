// 代码生成时间: 2025-08-27 08:29:59
package main

import (
    "net/url"
    "strings"
    "log"
    "github.com/gobuffalo/buffalo"
)

// UrlValidator is a struct used to validate URLs
type UrlValidator struct{}

// Validate checks if the provided URL is valid
// It returns a boolean indicating validity and an error if any
func (uv *UrlValidator) Validate(u string) (bool, error) {
    parsedUrl, err := url.ParseRequestURI(u)
    if err != nil {
        return false, err
    }

    // Check if scheme is present and host is not empty
    return parsedUrl.Scheme != "" && parsedUrl.Host != "", nil
}

// App is the main application struct
type App struct {
    validator *UrlValidator
}

// NewApp creates a new instance of App with a UrlValidator
func NewApp() *App {
    return &App{
        validator: &UrlValidator{},
    }
}

// URLValidationHandler is the handler for URL validation
// It accepts a URL as a query parameter and returns a JSON response with validation result
func (a *App) URLValidationHandler(c buffalo.Context) error {
    // Extract URL from query parameter
    queryURL := c.Param("url")

    // Validate the URL
    valid, err := a.validator.Validate(queryURL)
    if err != nil {
        // Handle error in URL validation
        log.Printf("Error validating URL: %s
", err)
        return c.Render(400, r.JSON(map[string]string{
            "error": "Invalid URL provided",
        }))
    }

    // Return a JSON response with validation result
    result := map[string]bool{
        "isValid": valid,
    }
    return c.Render(200, r.JSON(result))
}

func main() {
    app := buffalo.Automatic(buffalo.Options{
        Env:       "development",
        Logger:    buffalo.NewConsoleLogger(),
        AppName:   "url-validator",
        -assetsBox: buffalo.Box("../buffalo-assets"),
    })

    app.GET("/validate", NewApp().URLValidationHandler)
    app.Serve()
}
