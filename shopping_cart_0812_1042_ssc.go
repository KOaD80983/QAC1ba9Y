// 代码生成时间: 2025-08-12 10:42:39
package main

import (
    "buffalo"
    "github.com/markbates/buffalo(buffalo)"
    "github.com/markbates/buffalo-pop/pop"
)

// CartItem represents an item in the shopping cart
type CartItem struct {
    ID       uint   `json:"id"`
    Product  string `json:"product"`
    Quantity uint   `json:"quantity"`
    Price    float64 `json:"price"`
}

// Cart represents the shopping cart
type Cart struct {
    Items []CartItem `json:"items"`
}

// AddItem adds an item to the cart
func (c *Cart) AddItem(item CartItem) error {
    // Check if item already exists in the cart
    for _, i := range c.Items {
        if i.Product == item.Product {
            // If item exists, increase the quantity
            i.Quantity += item.Quantity
            return nil
        }
    }
    // If item does not exist, add it to the cart
    c.Items = append(c.Items, item)
    return nil
}

// RemoveItem removes an item from the cart by product name
func (c *Cart) RemoveItem(product string) error {
    for i, item := range c.Items {
        if item.Product == product {
            // If item is found, remove it from the cart
            c.Items = append(c.Items[:i], c.Items[i+1:]...)
            return nil
        }
    }
    return errors.New("Item not found in the cart")
}

// UpdateQuantity updates the quantity of an item in the cart
func (c *Cart) UpdateQuantity(product string, quantity uint) error {
    for i, item := range c.Items {
        if item.Product == product {
            c.Items[i].Quantity = quantity
            return nil
        }
    }
    return errors.New("Item not found in the cart")
}

// CartResource represents the resource for the shopping cart
type CartResource struct {
    DB *pop.Connection
}

// Create adds a new item to the cart
func (a CartResource) Create(c buffalo.Context) error {
    var item CartItem
    if err := c.Bind(&item); err != nil {
        return err
    }
    c.Cart.AddItem(item)
    return c.Render(200, buffalo.JSON(item))
}

// Show retrieves the current cart
func (a CartResource) Show(c buffalo.Context) error {
    return c.Render(200, buffalo.JSON(c.Cart.Items))
}

// Destroy removes an item from the cart
func (a CartResource) Destroy(c buffalo.Context) error {
    product := c.Param("product")
    if err := c.Cart.RemoveItem(product); err != nil {
        return c.Error(404, err)
    }
    return c.Render(200, buffalo.JSON(c.Cart.Items))
}

// Update updates the quantity of an item in the cart
func (a CartResource) Update(c buffalo.Context) error {
    product := c.Param("product")
    var quantity uint
    if err := c.Bind(&quantity); err != nil {
        return err
    }
    if err := c.Cart.UpdateQuantity(product, quantity); err != nil {
        return c.Error(404, err)
    }
    return c.Render(200, buffalo.JSON(c.Cart.Items))
}

func main() {
    app := buffalo.New(buffalo.Options{
        Env:        buffalo.EnvConfig{"buffalo.env", "development"},
        Session:    buffalo.SessionStore{},
        TemplateBox: buffalo.NewTmplBox("templates"),
    })

    app.Resource("/cart", CartResource{})
    app.Serve()
}