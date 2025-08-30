// 代码生成时间: 2025-08-30 17:57:55
package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"

    "github.com/markbates/buffalo"
    "github.com/markbates/buffalo/generators"
)

// TestdataGenerator is a struct that holds the data for the generator.
type TestdataGenerator struct {
    *generators.DefaultGenerator
}

// NewTestdataGenerator creates a new instance of TestdataGenerator.
func NewTestdataGenerator() *TestdataGenerator {
    return &TestdataGenerator{
        DefaultGenerator: generators.DefaultGenerator{
            Run: true,
        },
    }
}

// Run is a method that runs the test data generator.
func (g *TestdataGenerator) Run(args []string, root string, flags buffalo.Flags) error {
    if len(args) != 1 {
        return fmt.Errorf("a single argument is required: <testdata_file>.")
    }
    testdataFileName := args[0]

    // Check if the file already exists and bail out if it does.
    if _, err := os.Stat(filepath.Join(root, testdataFileName)); err == nil {
        return fmt.Errorf("testdata file '%s' already exists", testdataFileName)
    }

    // Create a new test data file.
    file, err := os.Create(filepath.Join(root, testdataFileName))
    if err != nil {
        return fmt.Errorf("error creating testdata file: %s", err)
    }
    defer file.Close()

    // Generate the test data file content.
    if _, err := file.WriteString(testdataContent()); err != nil {
        return fmt.Errorf("error writing testdata file: %s", err)
    }

    log.Println("Test data file generated successfully.")
    return nil
}

// testdataContent returns the content of the test data file.
func testdataContent() string {
    return `// This is a generated test data file.

// Test data example
var testData = []struct {
    Name    string
    Email   string
    Address string
}{
    // Add your test data here.
    {
        Name:    "John Doe",
        Email:   "john@example.com",
        Address: "123 Main St",
    },
    // Add more test data as needed.
}
`
}

func main() {
    g := NewTestdataGenerator()
    if err := g.Run(os.Args[1:], "", buffalo.Flags{}); err != nil {
        log.Fatal(err)
    }
}
