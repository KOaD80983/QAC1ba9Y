// 代码生成时间: 2025-09-07 12:28:46
package main

import (
    "buffalo"
    "buffalo/render"
    "github.com/markbates/pkg/logging"
    "net/http"
)

// MathCalculator represents a set of mathematical operations.
type MathCalculator struct {
    // You can add more properties if needed
}

// NewMathCalculator creates a new instance of MathCalculator.
func NewMathCalculator() *MathCalculator {
    return &MathCalculator{}
}

// Add performs addition of two numbers.
func (c *MathCalculator) Add(a, b int) (int, error) {
    if a < 0 || b < 0 {
        return 0, buffalo.NewError("Input cannot be negative", http.StatusBadRequest)
    }
    return a + b, nil
}

// Subtract performs subtraction of two numbers.
func (c *MathCalculator) Subtract(a, b int) (int, error) {
    if a < 0 || b < 0 {
        return 0, buffalo.NewError("Input cannot be negative", http.StatusBadRequest)
    }
    if a < b {
        return 0, buffalo.NewError("First number must be greater than or equal to the second", http.StatusBadRequest)
    }
    return a - b, nil
}

// Multiply performs multiplication of two numbers.
func (c *MathCalculator) Multiply(a, b int) (int, error) {
    if a < 0 || b < 0 {
        return 0, buffalo.NewError("Input cannot be negative", http.StatusBadRequest)
    }
    return a * b, nil
}

// Divide performs division of two numbers.
func (c *MathCalculator) Divide(a, b int) (int, error) {
    if a < 0 || b < 0 {
        return 0, buffalo.NewError("Input cannot be negative", http.StatusBadRequest)
    }
    if b == 0 {
        return 0, buffalo.NewError("Cannot divide by zero", http.StatusBadRequest)
    }
    return a / b, nil
}

// MathController handles requests related to mathematical operations.
type MathController struct {
    *buffalo.Context
    Trans render.Renderer
    Calc *MathCalculator
}

// Add action performs addition.
func (c MathController) Add() error {
    a := c.Param("a")
    b := c.Param("b")
    
    numA, err := strconv.Atoi(a)
    if err != nil {
        return buffalo.NewError("Invalid input for a", http.StatusBadRequest)
    }
    numB, err := strconv.Atoi(b)
    if err != nil {
        return buffalo.NewError("Invalid input for b", http.StatusBadRequest)
    }
    
    result, err := c.Calc.Add(numA, numB)
    if err != nil {
        return err
    }
    c.SetOK().Render(render.JSON(map[string]int{"result": result}))
    return nil
}

// Subtract action performs subtraction.
func (c MathController) Subtract() error {
    a := c.Param("a")
    b := c.Param("b")
    
    numA, err := strconv.Atoi(a)
    if err != nil {
        return buffalo.NewError("Invalid input for a", http.StatusBadRequest)
    }
    numB, err := strconv.Atoi(b)
    if err != nil {
        return buffalo.NewError("Invalid input for b", http.StatusBadRequest)
    }
    
    result, err := c.Calc.Subtract(numA, numB)
    if err != nil {
        return err
    }
    c.SetOK().Render(render.JSON(map[string]int{"result": result}))
    return nil
}

// Multiply action performs multiplication.
func (c MathController) Multiply() error {
    a := c.Param("a")
    b := c.Param("b")
    
    numA, err := strconv.Atoi(a)
    if err != nil {
        return buffalo.NewError("Invalid input for a", http.StatusBadRequest)
    }
    numB, err := strconv.Atoi(b)
    if err != nil {
        return buffalo.NewError("Invalid input for b", http.StatusBadRequest)
    }
    
    result, err := c.Calc.Multiply(numA, numB)
    if err != nil {
        return err
    }
    c.SetOK().Render(render.JSON(map[string]int{"result": result}))
    return nil
}

// Divide action performs division.
func (c MathController) Divide() error {
    a := c.Param("a")
    b := c.Param("b")
    
    numA, err := strconv.Atoi(a)
    if err != nil {
        return buffalo.NewError("Invalid input for a", http.StatusBadRequest)
    }
    numB, err := strconv.Atoi(b)
    if err != nil {
        return buffalo.NewError("Invalid input for b", http.StatusBadRequest)
    }
    
    result, err := c.Calc.Divide(numA, numB)
    if err != nil {
        return err
    }
    c.SetOK().Render(render.JSON(map[string]int{"result": result}))
    return nil
}

// main function initializes the Buffalo application.
func main() {
    app := buffalo.Automatic(buffalo.Options{
        Preinitializers: []buffalo.Preinitializer{
            func(e *buffalo.Env) error {
                calc := NewMathCalculator()
                e.Set(calc)
                return nil
            },
        },
    })

    // Define routes
    app.GET("/add/:a/:b", func(c buffalo.Context) error {
        return (MathController{Context: c, Calc: c.Value("mathCalculator").(*MathCalculator)}).Add()
    })
    
    app.GET("/subtract/:a/:b", func(c buffalo.Context) error {
        return (MathController{Context: c, Calc: c.Value("mathCalculator").(*MathCalculator)}).Subtract()
    })
    
    app.GET("/multiply/:a/:b", func(c buffalo.Context) error {
        return (MathController{Context: c, Calc: c.Value("mathCalculator").(*MathCalculator)}.Multiply())
    })
    
    app.GET("/divide/:a/:b", func(c buffalo.Context) error {
        return (MathController{Context: c, Calc: c.Value("mathCalculator").(*MathCalculator)}).Divide()
    })

    // Run the application
    app.Serve()
}
