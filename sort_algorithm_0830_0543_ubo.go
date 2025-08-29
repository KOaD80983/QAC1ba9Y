// 代码生成时间: 2025-08-30 05:43:55
package main

import (
    "buffalo"
    "buffalo/buffaloplug"
    "buffalo/middleware"
    "github.com/gobuffalo/packd"
    "github.com/gobuffalo/packr/v2"
    "log"
)

// main is the main entry point of the application.
func main() {
    // Create a new Buffalo app instance
    app := buffalo.Automatic(buffaloplug.Default)

    // Set up middleware for the app
    app.Use(middleware.Logger)
    app.Use(middleware.Recover)
    app.Use(middleware.RequestLogger)

    // Set up the application's routes
    app.GET("/", HomeHandler)
    app.POST("/sort", SortHandler)

    // Start the application
    log.Fatal(app.Serve(":3000"))
}

// HomeHandler is the handler for the root route.
func HomeHandler(c buffalo.Context) error {
    return c.Render(200, buffalo.R.HTML("index.html"))
}

// SortHandler is the handler for sorting an array of integers.
func SortHandler(c buffalo.Context) error {
    var input []int
    if err := c.Bind(&input); err != nil {
        return c.Error(400, err)
    }

    // Perform the sort using the BubbleSort algorithm
    BubbleSort(input)

    // Return the sorted array as JSON
    return c.Render(200, buffalo.R.JSON(input))
}

// BubbleSort is an implementation of the Bubble Sort algorithm.
func BubbleSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
}
