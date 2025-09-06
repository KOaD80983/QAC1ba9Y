// 代码生成时间: 2025-09-07 04:37:29
 * It allows users to organize their directories by moving files into a specified structure.
 *
 * @author Your Name
 * @version 1.0
 */

package main

import (
    "log"
    "net/http"
    "os"
    "path/filepath"
    "sort"

    "github.com/gobuffalo/buffalo"
)

// FolderOrganizer is the main handler for the application
type FolderOrganizer struct{}

// New creates a new FolderOrganizer
func (f *FolderOrganizer) New() buffalo.Handler {
    return func(c buffalo.Context) error {
        // Get the directory path from the query parameter
        dirPath := c.Param("dir")
        if dirPath == "" {
            return buffalo.NewError("Directory path is required")
        }

        // Check if the directory exists
        if _, err := os.Stat(dirPath); os.IsNotExist(err) {
            return buffalo.NewError("Directory does not exist")
        }

        // Attempt to organize the directory
        if err := f.organizeDirectory(dirPath); err != nil {
            return err
        }

        // Return a success message
        return c.Render(200, buffalo.R.JSON(map[string]string{
            "message": "Directory organized successfully",
        }))
    }
}

// organizeDirectory sorts the files in the given directory into a specified structure
func (f *FolderOrganizer) organizeDirectory(dirPath string) error {
    // Get a list of all files in the directory
    files, err := os.ReadDir(dirPath)
    if err != nil {
        return err
    }

    // Sort the files by name
    sort.Slice(files, func(i, j int) bool {
        return files[i].Name() < files[j].Name()
    })

    // Create a new directory structure
    for _, file := range files {
        // Define the new path for the file
        newDirPath := filepath.Join(dirPath, file.Name())

        // Check if the file is not already in the correct location
        if !file.IsDir() && filepath.Ext(file.Name()) != "" {
            // Move the file to the new directory
            if err := os.Rename(filepath.Join(dirPath, file.Name()), newDirPath); err != nil {
                return err
            }
        }
    }

    return nil
}

func main() {
    app := buffalo.New(buffalo.Options{})
    app.GET("/organize", FolderOrganizer{}.New())
   	defer app.Serve()
    log.Fatal(http.ListenAndServe(":3000", app))
}