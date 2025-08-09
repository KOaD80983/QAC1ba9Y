// 代码生成时间: 2025-08-10 05:09:03
// folder_organizer.go

package main

import (
    "bufio"
    "flag"
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"
    "time"
)

// FolderOrganizer is a struct that holds the configuration for organizing folders.
type FolderOrganizer struct {
    SourcePath string
    DestinationPath string
    FileNameTemplate string
    DryRun bool
}

// NewFolderOrganizer creates a new instance of FolderOrganizer.
func NewFolderOrganizer(sourcePath, destinationPath, fileNameTemplate string, dryRun bool) *FolderOrganizer {
    return &FolderOrganizer{
        SourcePath: sourcePath,
        DestinationPath: destinationPath,
        FileNameTemplate: fileNameTemplate,
        DryRun: dryRun,
    }
}

// Organize runs the folder organizing process.
func (f *FolderOrganizer) Organize() error {
    files, err := os.ReadDir(f.SourcePath)
    if err != nil {
        return err
    }

    for _, file := range files {
        if file.IsDir() {
            continue
