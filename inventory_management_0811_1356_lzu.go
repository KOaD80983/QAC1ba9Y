// 代码生成时间: 2025-08-11 13:56:31
package main

import (
    "buffalo"
    "fmt"
    "github.com/gobuffalo/pop/v6"
    "github.com/gobuffalo/buffalo-plugins/plugins"
    "github.com/gobuffalo/buffalo-plugins/plugins/plugcmd"
    "github.com/markbates/validate"
)

// InventoryItem represents a single item in the inventory.
type InventoryItem struct {
    ID    uint   "json:"id" db:"id" xml:"id""
    Name  string "json:"name" db:"name" xml:"name""
    Quantity int    "json:"quantity" db:"quantity" xml:"quantity""
}

// InventoryResource represents the resource for managing inventory items.
type InventoryResource struct {
    *buffalo.Resource
}

// NewInventoryResource creates a new InventoryResource instance.
func NewInventoryResource(c buffalo.Context) *InventoryResource {
    return &InventoryResource{Resource: buffalo.NewResource(c)}
}

// List returns a list of all inventory items.
func (v *InventoryResource) List(c buffalo.Context) error {
    var items []InventoryItem
    // Use database connection (pop) to fetch items.
    if err := v.DB.All(&items); err != nil {
        return buffalo.NewError(err)
    }
    // Return the list of items as JSON response.
    return c.Render(200, r.JSON(items))
}

// Show returns a single inventory item.
func (v *InventoryResource) Show(c buffalo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return buffalo.NewError(err)
    }
    item := InventoryItem{}
    // Use database connection (pop) to fetch item by ID.
    if err := v.DB.Find(&item, id); err != nil {
        return buffalo.NewError(err)
    }
    // Return the item as JSON response.
    return c.Render(200, r.JSON(item))
}

// Create adds a new inventory item to the database.
func (v *InventoryResource) Create(c buffalo.Context) error {
    var item InventoryItem
    // Unmarshal the request body into the struct.
    if err := c.Bind(&item); err != nil {
        return err
    }
    // Validate the item data.
    if err := item.Validate(); err != nil {
        return err
    }
    // Use database connection (pop) to create the item.
    if err := v.DB.Create(&item); err != nil {
        return buffalo.NewError(err)
    }
    // Return the created item as JSON response.
    return c.Render(201, r.JSON(item))
}

// Update modifies an existing inventory item.
func (v *InventoryResource) Update(c buffalo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return buffalo.NewError(err)
    }
    var item InventoryItem
    // Use database connection (pop) to fetch item by ID.
    if err := v.DB.Find(&item, id); err != nil {
        return buffalo.NewError(err)
    }
    // Unmarshal the request body into the struct.
    if err := c.Bind(&item); err != nil {
        return err
    }
    // Validate the item data.
    if err := item.Validate(); err != nil {
        return err
    }
    // Use database connection (pop) to update the item.
    if err := v.DB.Update(&item); err != nil {
        return buffalo.NewError(err)
    }
    // Return the updated item as JSON response.
    return c.Render(200, r.JSON(item))
}

// Destroy removes an inventory item from the database.
func (v *InventoryResource) Destroy(c buffalo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return buffalo.NewError(err)
    }
    item := InventoryItem{}
    // Use database connection (pop) to remove item by ID.
    if err := v.DB.Destroy(&item, id); err != nil {
        return buffalo.NewError(err)
    }
    // Return a success message as JSON response.
    return c.Render(200, r.String("Item deleted successfully"))
}

// Validate checks if the inventory item data is valid.
func (i *InventoryItem) Validate(tx *pop.Connection) error {
    errors := validate.NewErrors()
    // Validate item name is not empty.
    if i.Name == "" {
        errors.Add("name", "cannot be blank")
    }
    // Validate item quantity is greater than 0.
    if i.Quantity <= 0 {
        errors.Add("quantity", "must be greater than 0")
    }
    // Return errors if any.
    return errors
}

// main is the entry point for the Buffalo application.
func main() {
    // Create a new Buffalo application instance.
    app := buffalo.New(buffalo.Options{})

    // Mount resources.
    app.Resource("/inventory", NewInventoryResource(buffalo.DefaultContext),
        buff.AllowedMethods([]string{"GET", "POST"}),
        buff.AllowedMethods([]string{"GET", "PATCH", "DELETE"}),
    )

    // Start the application.
    if err := app.Serve(); err != nil {
        fmt.Println(err)
    }
}