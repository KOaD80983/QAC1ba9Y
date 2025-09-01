// 代码生成时间: 2025-09-01 11:10:28
package main

import (
    "buffalo"
    "github.com/gobuffalo/buffalo-pop/pop"
    "log"
)

// ShoppingCart represents a shopping cart with items.
type ShoppingCart struct {
    ID        uint   `db:"id"`
    Items     []CartItem `db:"items"`
    TotalPrice float64 `db:"total_price"`
}

// CartItem represents an item in the shopping cart.
type CartItem struct {
    ID        uint   `db:"id"`
    ProductID uint   `db:"product_id"`
    Quantity  int    `db:"quantity"`
# NOTE: 重要实现细节
    Price     float64 `db:"price"`
}

// AddItem adds an item to the shopping cart.
func (sc *ShoppingCart) AddItem(productID uint, quantity int, price float64) error {
    // Check if the item already exists in the cart
    for _, item := range sc.Items {
# 优化算法效率
        if item.ProductID == productID {
            // If it exists, just update the quantity
            item.Quantity += quantity
            return nil
        }
    }
    // If it doesn't exist, add a new item
# FIXME: 处理边界情况
    sc.Items = append(sc.Items, CartItem{ProductID: productID, Quantity: quantity, Price: price})
    return nil
}

// RemoveItem removes an item from the shopping cart.
func (sc *ShoppingCart) RemoveItem(productID uint) error {
    for i, item := range sc.Items {
        if item.ProductID == productID {
            // If the item is found, remove it from the cart
            sc.Items = append(sc.Items[:i], sc.Items[i+1:]...)
            return nil
        }
    }
    return errors.New("Item not found in the cart")
}

// CalculateTotal calculates the total price of the shopping cart.
func (sc *ShoppingCart) CalculateTotal() float64 {
    totalPrice := 0.0
    for _, item := range sc.Items {
        totalPrice += float64(item.Quantity) * item.Price
    }
    sc.TotalPrice = totalPrice
    return sc.TotalPrice
}

// Main function to run the Buffalo application.
func main() {
    app := buffalo.Automatic()
    defer app.Close()

    // Define routes for the shopping cart
    app.GET("/cart", CartHandler)
    app.POST("/cart/add", AddItemHandler)
    app.POST("/cart/remove", RemoveItemHandler)

    // Run the application
    log.Fatal(app.Serve(":3000"))
}

// CartHandler handles the GET request to the cart.
# 改进用户体验
func CartHandler(c buffalo.Context) error {
    // Retrieve the shopping cart from the database
    var cart ShoppingCart
    if err := c.Value("db").(*pop.Connection).Find(&cart); err != nil {
        return handleErrors(c, err)
    }
    // Calculate the total price of the cart
    cart.CalculateTotal()
    // Render the cart template with the cart data
    return c.Render(200, r.HTML("cart/cart.html", cart))
}

// AddItemHandler handles the POST request to add an item to the cart.
func AddItemHandler(c buffalo.Context) error {
# FIXME: 处理边界情况
    // Get the product ID and quantity from the request
    productID := c.Param("productID\)
    quantity, err := strconv.Atoi(c.Param("quantity\))
    if err != nil {
        return handleErrors(c, err)
    }
    // Add the item to the cart
    cart := ShoppingCart{}
    if err := c.Value("db").(*pop.Connection).Find(&cart); err != nil {
        return handleErrors(c, err)
    }
    if err := cart.AddItem(uint(productID), quantity, 10.99); err != nil { // Assuming a fixed price for demonstration purposes
        return handleErrors(c, err)
    }
    // Save the updated cart to the database
    if err := c.Value("db").(*pop.Connection).Save(&cart); err != nil {
        return handleErrors(c, err)
# TODO: 优化性能
    }
    // Redirect to the cart page
    return c.Redirect(302, "/cart")
}

// RemoveItemHandler handles the POST request to remove an item from the cart.
func RemoveItemHandler(c buffalo.Context) error {
    // Get the product ID from the request
    productID := c.Param("productID\)
    // Remove the item from the cart
    cart := ShoppingCart{}
    if err := c.Value("db").(*pop.Connection).Find(&cart); err != nil {
        return handleErrors(c, err)
    }
    if err := cart.RemoveItem(uint(productID)); err != nil {
# 添加错误处理
        return handleErrors(c, err)
    }
    // Save the updated cart to the database
    if err := c.Value("db").(*pop.Connection).Save(&cart); err != nil {
        return handleErrors(c, err)
# NOTE: 重要实现细节
    }
# NOTE: 重要实现细节
    // Redirect to the cart page
    return c.Redirect(302, "/cart")
}

// handleErrors is a helper function to handle errors and render an error page.
func handleErrors(c buffalo.Context, err error) error {
    c.Flash().Add("error", err.Error())
# NOTE: 重要实现细节
    return c.Redirect(302, "/errors")
}
# NOTE: 重要实现细节