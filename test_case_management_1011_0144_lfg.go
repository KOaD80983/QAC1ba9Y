// 代码生成时间: 2025-10-11 01:44:31
package main

import (
    "buffalo"
    "github.com/gobuffalo/buffalo-pop/v2/pop"
    "github.com/markbates/inflect"
    "log"
)

// TestCase represents a Test Case entity.
type TestCase struct {
    ID        uint   `db:"id"`
    Title     string `db:"title"`
    Content   string `db:"content"`
    CreatedAt string `db:"created_at"`
    UpdatedAt string `db:"updated_at"`
}

// CreateTest cases handlers
func CreateTestCase(c buffalo.Context) error {
    // Start a transaction
    tx, err := pop.Begin(c.Request())
    if err != nil {
        return err
    }
    defer tx.Rollback()

    // Decode the struct
    var tc TestCase
    if err := c.Request().ParseForm(); err != nil {
        return err
    }
    if err := c.Bind(&tc); err != nil {
        return err
    }
    if err := tx.Create(&tc); err != nil {
        return err
    }
    if err := tx.Commit(); err != nil {
        return err
    }
    return c.Render(201, r.JSON(tc))
}

// ReadTest cases handlers
func ReadTestCase(c buffalo.Context) error {
    // Get the DB connection from the context
    tx, ok := c.Value("tx").(*pop.Connection)
    if !ok {
        return errors.New("no transaction found")
    }

    // Allocate an empty TestCase
    tc := &TestCase{}

    // To find the TestCase the parameter "id" is needed
    params := c.Params()
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        return err
    }

    // Query the database for the TestCase
    err = tx.Eager().Find(tc, id)
    if err != nil {
        return err
    }

    // Render the TestCase with status code
    return c.Render(200, r.JSON(tc))
}

// UpdateTest cases handlers
func UpdateTestCase(c buffalo.Context) error {
    // Get the DB connection from the context
    tx, ok := c.Value("tx").(*pop.Connection)
    if !ok {
        return errors.New("no transaction found")
    }

    // Allocate an empty TestCase
    tc := &TestCase{}

    // To find the TestCase the parameter "id" is needed
    params := c.Params()
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        return err
    }

    // Query the database for the TestCase
    err = tx.Find(tc, id)
    if err != nil {
        return err
    }
    // Decode the struct
    if err := c.Bind(tc); err != nil {
        return err
    }
    if err := tx.Update(tc); err != nil {
        return err
    }
    return c.Render(200, r.JSON(tc))
}

// DeleteTest cases handlers
func DeleteTestCase(c buffalo.Context) error {
    // Get the DB connection from the context
    tx, ok := c.Value("tx").(*pop.Connection)
    if !ok {
        return errors.New("no transaction found")
    }

    // To find the TestCase the parameter "id" is needed
    params := c.Params()
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        return err
    }

    // Query the database for the TestCase
    tc := &TestCase{}
    err = tx.Find(tc, id)
    if err != nil {
        return err
    }

    // Delete the TestCase
    if err := tx.Destroy(tc); err != nil {
        return err
    }
    return c.Render(200, r.JSON(tc))
}

// Main function to setup the Buffalo application
func main() {
    // Initialize the Buffalo application
    app := buffalo.Automatic()
    app.GET("/test_cases/:id", ReadTestCase)
    app.POST("/test_cases", CreateTestCase)
    app.PUT("/test_cases/:id", UpdateTestCase)
    app.DELETE("/test_cases/:id", DeleteTestCase)

    // Set the database connection
    app.Use(pop.ProvideConnection(
        inflect.Underscore("test_case"),
        pop.ConnectionDetails{
            DatabaseURL: "sqlite3:test.db",
        },
    ))

    // Start the Buffalo application
    if err := app.Start(); err != nil {
        log.Fatal(err)
    }
}