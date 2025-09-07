// 代码生成时间: 2025-09-08 06:57:01
package main

import (
    "buffalo"
    "github.com/gobuffalo/buffalo-pop/v2/pop"
    "github.com/gobuffalo/envy"
    "github.com/gobuffalo/packd"
    "github.com/gobuffalo/packr/v2"
    "log"
    "net/http"
    "os"
    "time"
)

// main is the entry point for the application.
func main() {
    // env is used to help switch settings based on the environment.
    env := envy.MustGet("BUFFALO_ENV", "development")
    if env == "development" {
        // In development, we want to enable reloading.
        // We also set the address and the port to serve on.
        app := buffalo.Automatic(buffalo.Options{})
        app.Serve()
    } else {
        // In production, we want to set the address and port as well,
        // but we want to load the assets from the filesystem.
        app := buffalo.Automatic(buffalo.Options{
            Env:  env,
            Addr: "0.0.0.0:3000",
            AssetsBox: packr.New("app:assets", path.Join("app", "assets"))})
        // Start the application.
        if err := app.Serve(); err != nil {
            log.Fatal(err)
        }
    }
}

// App is the application struct.
type App struct {
    *buffalo.App
}

// NewApp creates a new instance of the application.
func NewApp(opts ...buffalo.Option) *App {
    a := &App{
        App: buffalo.New(opts...),
    }
    a.App.Middleware.Insert(corsHandler, 0)
    a.App.Middleware.Insert(middleware.RequestLogger, 0)
    a.App.Middleware.Insert(middleware.Recovery, 0)
    a.App.Middleware.Insert(middleware.RequestID, 0)
    a.App.Middleware.Insert(middleware.RequestBodyLogger, 0)
    return a
}

// corsHandler handles CORS
func corsHandler(next buffalo.Handler) buffalo.Handler {
    return func(c buffalo.Context) error {
        c.Response().Header().Set("Access-Control-Allow-Origin", "*")
        return next(c)
    }
}

// DB is a convenience method to get the database connection.
func DB(c buffalo.Context) *pop.Connection {
    return c.Value("db").(*pop.Connection)
}

// Router is where we define our routes and their corresponding handlers.
func Router(app *buffalo.App) {
    // Define our resources
    app.Resource("/items", NewItemResource())
}

// ItemResource is a resource for managing items.
type ItemResource struct {
    // Standard Buffalo resource fields
}

// NewItemResource is a method to create a new instance of the ItemResource
func NewItemResource() *ItemResource {
    return &ItemResource{}
}

// List gets all items and is responded to by the GET /items endpoint.
func (v *ItemResource) List(c buffalo.Context) error {
    items := []Item{}
    // Get all items from the DB
    tx := DB(c).Where("1 = 1").All(&items)
    if tx.Error != nil {
        return BuffaloError{"error": tx.Error.Error()}
    }
    return c.Render(200, r.JSON(items))
}

// Show gets the data for one item and is responded to by the GET /items/{item_id} endpoint.
func (v *ItemResource) Show(c buffalo.Context) error {
    // Get the item by the item_id parameter from the URL.
    id := c.Param("item_id")
    var item Item
    // Check if the item exists in the database.
    tx := DB(c).Where("id = ?