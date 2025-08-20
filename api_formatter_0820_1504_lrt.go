// 代码生成时间: 2025-08-20 15:04:15
package main

import (
    "buffalo"
    "github.com/gobuffalo/buffalo/(buffalo)"
    "github.com/gobuffalo/buffalo/middleware"
    "net/http"
)

// APIResponseFormatter middleware formats the response from Buffalo to be more API-friendly
func APIResponseFormatter(next buffalo.Handler) buffalo.Handler {
    return func(ctx *buffalo.Context) error {
        // Call the next middleware in the chain
        err := next(ctx)
        if err != nil {
            return err
        }

        // Format the response
        formatResponse(ctx)
        return nil
    }
}

// formatResponse takes care of formatting the standard Buffalo response
func formatResponse(ctx *buffalo.Context) {
    if ctx.Response().Status == http.StatusOK {
        // Convert the response to a JSON response with the data from the context
        ctx.JSON(http.StatusOK, ctx.Data())
    } else if ctx.Response().Status == http.StatusNotFound {
        // Handle not found responses
        ctx.JSON(http.StatusNotFound, map[string]string{"error": "Resource not found"})
    } else {
        // Handle other responses
        ctx.JSON(ctx.Response().Status, ctx.Data())
    }
}

func main() {
    app := buffalo.Automatic()
    app.Use(middleware.Logger)
    app.Use(middleware._recovery.Recovery)
    app.Use(APIResponseFormatter)

    // Define your routes here, for example:
    // app.GET("/", HomeHandler)
    // app.POST("/items", ItemsCreateHandler)

    app.Serve()
}
