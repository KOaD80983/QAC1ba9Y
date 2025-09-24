// 代码生成时间: 2025-09-24 16:09:25
It provides an endpoint to upload and extract zip files.
*/

package main

import (
    "log"
    "net/http"
    "os"
    "path/filepath"
    "io/ioutil"
    "strings"
    "archive/zip"

    "github.com/gobuffalo/buffalo"
)

// unzipHandler handles the file upload and extraction process.
func unzipHandler(c buffalo.Context) error {
    // Get the zip file from the request's form data.
    file, header, err := c.Request().FormFile("file")
    if err != nil {
        return buffalo.NewError("Error retrieving file from request", http.StatusBadRequest)
    }
    defer file.Close()

    // Get the destination directory from the request's form data.
    destination := c.Request().PostFormValue("destination")
    if destination == "" {
        destination = "./extracted"
    }

    // Ensure the destination directory exists or create it.
    if _, err := os.Stat(destination); os.IsNotExist(err) {
        if err := os.MkdirAll(destination, 0755); err != nil {
            return buffalo.NewError("Error creating destination directory", http.StatusInternalServerError)
        }
    }

    // Open the zip file.
    zipReader, err := zip.OpenReader(header.Filename)
    if err != nil {
        return buffalo.NewError("Error opening zip file", http.StatusInternalServerError)
    }
    defer zipReader.Close()

    // Loop through the files in the zip and extract them.
    for _, f := range zipReader.File {
        rc, err := f.Open()
        if err != nil {
            return buffalo.NewError("Error opening file inside zip", http.StatusInternalServerError)
        }
        defer rc.Close()

        // Create the full path to the file.
        path := filepath.Join(destination, f.Name)
        if f.FileInfo().IsDir() {
            // Make directory.
            if err := os.MkdirAll(path, 0755); err != nil {
                return buffalo.NewError("Error creating directory", http.StatusInternalServerError)
            }
        } else {
            // Create file.
            f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
            if err != nil {
                return buffalo.NewError("Error creating file", http.StatusInternalServerError)
            }
            defer f.Close()
            _, err = io.Copy(f, rc)
            if err != nil {
                return buffalo.NewError("Error copying file", http.StatusInternalServerError)
            }
        }
    }

    // Return a success message.
    return c.Render(200, buffalo.RenderOptions{"json": map[string]interface{}{"message": "Files extracted successfully"}})
}

func main() {
    // Create a new Buffalo app.
    app := buffalo.Automatic()

    // Register the unzipHandler with the app to handle POST requests to /unzip.
    app.POST("/unzip", unzipHandler)

    // Run the app.
    app.Serve()
}
