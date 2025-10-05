// 代码生成时间: 2025-10-05 19:25:02
package main

import (
    "buffalo"
    "buffalo/render"
    "github.com/markbates/inflect"
    "log"
)

// License represents a software license
type License struct {
    ID     uint   `db:"id"`
    Name   string `db:"name"`
    Key    string `db:"key"`
    UserID uint   `db:"user_id"`
}

// LicensesResource is the resource for handling licenses
type LicensesResource struct{}

// List handles the GET request for licenses
func (l *LicensesResource) List(c buffalo.Context) error {
    // Get the DB connection from the context
    tx := c.Value("tx").(*pop.Connection)
    // Query the database for all licenses
    licenses := []License{}
    // Use error handling
    if err := tx.Where("user_id = ?", c.Session().Get("current_user_id")).All(&licenses); err != nil {
        return buffalo.NewError(err, 500)
    }
    // Render the licenses with the HTML template
    return c.Render(200, render.HTML("licenses/index.html", licenses))
}

// Show handles the GET request to show a single license
func (l *LicensesResource) Show(c buffalo.Context) error {
    // Get the license ID from the params
    id := c.Param("id")
    // Get the DB connection from the context
    tx := c.Value("tx").(*pop.Connection)
    // Query the database for the license
    var license License
    // Use error handling
    if err := tx.Where("id = ?", id).First(&license); err != nil {
        return buffalo.NewError(err, 500)
    }
    // Render the license with the HTML template
    return c.Render(200, render.HTML("licenses/show.html", license))
}

// New handles the GET request for creating a new license
func (l *LicensesResource) New(c buffalo.Context) error {
    // Render the form for creating a new license
    return c.Render(200, render.HTML("licenses/new.html"))
}

// Create handles the POST request for creating a new license
func (l *LicensesResource) Create(c buffalo.Context) error {
    // Get the DB connection from the context
    tx := c.Value("tx").(*pop.Connection)
    // Create a new License struct
    var license License
    // Bind the license to the context
    if err := c.Bind(&license); err != nil {
        return err
    }
    // Validate the license
    if err := tx.ValidateAndCreate(&license); err != nil {
        return buffalo.NewError(err, 500)
    }
    // Redirect to the license show page
    return c.Redirect(302, "/licenses/"+license.ID.String())
}

// Edit handles the GET request for editing a license
func (l *LicensesResource) Edit(c buffalo.Context) error {
    // Get the license ID from the params
    id := c.Param("id")
    // Get the DB connection from the context
    tx := c.Value("tx").(*pop.Connection)
    // Query the database for the license
    var license License
    // Use error handling
    if err := tx.Where("id = ?", id).First(&license); err != nil {
        return buffalo.NewError(err, 500)
    }
    // Render the form for editing the license
    return c.Render(200, render.HTML("licenses/edit.html", license))
}

// Update handles the PUT/PATCH request for updating a license
func (l *LicensesResource) Update(c buffalo.Context) error {
    // Get the license ID from the params
    id := c.Param("id\)
    // Get the DB connection from the context
    tx := c.Value("tx").(*pop.Connection)
    // Query the database for the license
    var license License
    // Use error handling
    if err := tx.Where("id = ?", id).First(&license); err != nil {
        return buffalo.NewError(err, 500)
    }
    // Bind the license to the context
    if err := c.Bind(&license); err != nil {
        return err
    }
    // Validate the license
    if err := tx.ValidateAndSave(&license); err != nil {
        return buffalo.NewError(err, 500)
    }
    // Redirect to the license show page
    return c.Redirect(302, "/licenses/"+license.ID.String())
}

// Destroy handles the DELETE request for deleting a license
func (l *LicensesResource) Destroy(c buffalo.Context) error {
    // Get the license ID from the params
    id := c.Param("id\)
    // Get the DB connection from the context
    tx := c.Value("tx").(*pop.Connection)
    // Query the database for the license
    var license License
    // Use error handling
    if err := tx.Where("id = ?", id).First(&license); err != nil {
        return buffalo.NewError(err, 500)
    }
    // Delete the license
    if err := tx.Destroy(&license); err != nil {
        return buffalo.NewError(err, 500)
    }
    // Redirect to the licenses list page
    return c.Redirect(302, "/licenses")
}

// main is the entry point for the application
func main() {
    // Define the application
    app := buffalo.New(buffalo.Options{
        Env:                  "development",
        SessionStoreKey:      "session",
        SessionStoreSecret:   "secret",
        StaticFilesRootPath: "public",
        TemplatesRootPath:   "templates",
    })

    // Set the DB connection string
    app.ENV["DATABASE_URL"] = "postgres://user:password@localhost/dbname?sslmode=disable"

    // Setup the Pop/SQRL database connection
    app.Use(pop.Pop())

    // Set the authenticate middleware
    app.Use(middleware.Authenticate)

    // Register the resources
    app.Resource("/licenses", &LicensesResource{})

    // Start the application
    app.Serve()
}
