// 代码生成时间: 2025-08-19 03:45:16
package main

import (
    "github.com/gobuffalo/buffalo"
    "github.com/gobuffalo/buffalo/generators"
    "github.com/gobuffalo/buffalo/generators/assets/webpack"
    "github.com/gobuffalo/buffalo/meta/inflect"
    "github.com/gobuffalo/envy"
    "github.com/markbates/inflect"
    "net/http"
)

// App is the struct that represents the application.
// It contains the entire app state. This is where you would also
// add any middleware, models, templates, etc.
type App struct {
    *buffalo.App
}

// NewApp creates a new App
// It sets up any necessary middleware, configures the app,
// and sets up the routes for the application.
func NewApp() *App {
    e := envy.Env()
    app := buffalo.New(buffalo.Options{
        Env: e,
    })

    // Automatically set the layout for each page
    app.Use(layouts.Provide())

    // Wrap actions with actions that should be run before
    // and after each request.
    app.Middleware.Before("github.com/gobuffalo/x/middleware/logger")
    app.Middleware.Before("github.com/gobuffalo/x/middleware/recover")
    app.Middleware.After("github.com/gobuffalo/x/middleware/buffalo")

    // app.Middleware.Skip("github.com/gobuffalo/buffalo/render")
    // app.Middleware.Skip("github.com/gobuffalo/buffalo/generators/assets/webpack")

    // Setup a layout for the application.
    // The layout is the HTML that is shared across all pages.
    // In this case, it's a simple HTML5 doc with a div that the application
    // will mount into.
    app.Use(layout())

    // app.GET("/", HomeHandler)
    // app.POST("/form", FormSubmitHandler)
    // app.GET("/form", FormHandler)

    // Add routes to the application. These routes use the same function
    // names as the ones defined in the actions.
    app.GET("/", HomeHandler)

    return &App{App: app}
}

// HomeHandler is a function that renders the Index page.
func HomeHandler(c buffalo.Context) error {
    // Add any data to the context that you want to be available in the template.
    // c.Set("someKey", "someValue")

    // Return a simple JSON response with a welcome message.
    return c.Render(200, r.JSON(map[string]string{"message": "Welcome to the responsive layout!"}))
}

// layout is a function that sets up the layout for the application.
// It is used by the layout middleware.
func layout() buffalo.MiddlewareFunc {
    return func(h buffalo.Handler) buffalo.Handler {
        return func(c buffalo.Context) error {
            // Run the handler that the middleware is wrapping.
            err := h(c)
            if err != nil {
                return err
            }

            // Get the current render path.
            rp := c.Request().URL.Path
            // If the path is for a file (not a buffalo route)
            // then we do not want to wrap the layout around it.
            if strings.HasPrefix(rp, "/assets/") || strings.HasPrefix(rp, "/buffalo/") {
                return nil
            }

            // Get the current box to be rendered.
            b, err := c.RenderNamed(rp)
            if err != nil {
                return err
            }

            // Wrap the current box with the layout template.
            c.Set("layout", "base")
            c.Set("content", b)
            return c.Render(200, r.HTMLBytes("layouts/base.plush.html"))
        }
    }
}

// main is the entry point for the application.
func main() {
    if err := NewApp().Serve(); err != nil {
        log.Fatal(err)
    }
}
