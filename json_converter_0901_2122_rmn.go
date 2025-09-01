// 代码生成时间: 2025-09-01 21:22:55
package main

import (
    "buffalo"
    "encoding/json"
    "net/http"
)

// JsonConverter is a struct that contains the necessary properties for the converter.
type JsonConverter struct{}

// ConvertJSON handles HTTP requests and converts JSON data.
func (c *JsonConverter) ConvertJSON(w http.ResponseWriter, r *http.Request) error {
    // Decode the JSON data from the request body into a map for flexibility.
    var jsonData map[string]interface{}
    if err := json.NewDecoder(r.Body).Decode(&jsonData); err != nil {
        // Return a 400 status code and an error message if decoding fails.
        return buffalo.NewError(err, http.StatusBadRequest)
    }
    defer r.Body.Close()

    // Convert the map back to JSON and write it to the response.
    responseJSON, err := json.Marshal(jsonData)
    if err != nil {
        // Return a 500 status code and an error message if encoding fails.
        return buffalo.NewError(err, http.StatusInternalServerError)
    }

    // Set the content type of the response to JSON.
    w.Header().Set("Content-Type", "application/json")
    // Write the JSON response to the client.
    _, _ = w.Write(responseJSON)
    return nil
}

// main is the entry point of the application.
func main() {
    // Create a new Buffalo application instance.
    app := buffalo.Automatic(buffalo.Options{})

    // Register the JSON conversion handler.
    app.GET("/convert", func(c buffalo.Context) error {
        return (&JsonConverter{}).ConvertJSON(c.Response(), c.Request())
    })

    // Start the Buffalo application.
    app.Serve()
}