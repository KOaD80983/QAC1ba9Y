// 代码生成时间: 2025-09-12 06:31:13
package main

import (
    "bufio"
    "fmt"
    "image"
    "image/jpeg"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "strconv"
    "strings"

    "github.com/bmizerany/pat"
    "github.com/nfnt/resize"
)

// ImageResizer defines a struct for handling image resizing
type ImageResizer struct {
    Width  int
    Height int
}

// NewImageResizer creates an instance of ImageResizer
func NewImageResizer(width, height int) *ImageResizer {
    return &ImageResizer{Width: width, Height: height}
}

// ResizeImage resizes an image to the specified dimensions
func (ir *ImageResizer) ResizeImage(inputPath, outputPath string) error {
    file, err := os.Open(inputPath)
    if err != nil {
        return err
    }
    defer file.Close()

    img, _, err := image.Decode(file)
    if err != nil {
        return err
    }

    resizedImg := resize.Resize(ir.Width, ir.Height, img, resize.Lanczos3)
    outputPath = strings.TrimSuffix(outputPath, filepath.Ext(outputPath)) + "_resized.jpg"

    outFile, err := os.Create(outputPath)
    if err != nil {
        return err
    }
    defer outFile.Close()

    err = jpeg.Encode(outFile, resizedImg, nil)
    if err != nil {
        return err
    }

    return nil
}

// ProcessImageDirectory resizes all images within a specified directory
func (ir *ImageResizer) ProcessImageDirectory(directoryPath string) error {
    dir, err := os.Open(directoryPath)
    if err != nil {
        return err
    }
    defer dir.Close()

    files, err := dir.Readdir(-1)
    if err != nil {
        return err
    }

    for _, file := range files {
        if !file.IsDir() {
            inputFile := filepath.Join(directoryPath, file.Name())
            outputFile := filepath.Join(directoryPath, "resized_" + file.Name())
            err := ir.ResizeImage(inputFile, outputFile)
            if err != nil {
                return err
            }
        }
    }
    return nil
}

// SetupRoutes sets up the routes for the image resizing application
func SetupRoutes(resizer *ImageResizer, m *pat.PatternServeMux) {
    m.Get("/resize/{width}/{height}", func(w http.ResponseWriter, r *http.Request) {
        vars := patURLVars(r)
        width, _ := strconv.Atoi(vars["width"])
        height, _ := strconv.Atoi(vars["height"])
        newResizer := NewImageResizer(width, height)
        directoryPath := r.FormValue("dir")
        err := newResizer.ProcessImageDirectory(directoryPath)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        fmt.Fprintf(w, "Images resized successfully")
    })
}

// patURLVars extracts URL variables from an HTTP request
func patURLVars(r *http.Request) map[string]string {
    parser := pat.Get()
    pattern, _ := parser.Find(r.Method, r.URL.Path)
    vars := make(map[string]string)
    if pattern != nil && pattern.Vars != nil {
        for _, varName := range pattern.Vars.Names() {
            vars[varName] = pattern.Vars.Value(varName)
        }
    }
    return vars
}

func main() {
    M := pat.New()
    resizer := NewImageResizer(800, 600) // Default resizing dimensions
    SetupRoutes(resizer, M)
    M.ServeHTTP()
}
