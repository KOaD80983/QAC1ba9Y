// 代码生成时间: 2025-08-01 03:34:42
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

// BackupSyncTool 结构体封装了备份和同步文件所需的信息
type BackupSyncTool struct {
    SourceDir  string
    DestinationDir string
    LastSyncTime time.Time
}

// NewBackupSyncTool 创建一个新的备份同步工具实例
func NewBackupSyncTool(sourceDir, destinationDir string) *BackupSyncTool {
    return &BackupSyncTool{
        SourceDir: sourceDir,
        DestinationDir: destinationDir,
        LastSyncTime: time.Now(),
    }
}

// Sync 同步源目录和目标目录之间的文件
func (bst *BackupSyncTool) Sync() error {
    files, err := os.ReadDir(bst.SourceDir)
    if err != nil {
        return fmt.Errorf("failed to read source directory: %w", err)
    }
    for _, file := range files {
        if file.IsDir() {
            continue
        }
        sourceFilePath := filepath.Join(bst.SourceDir, file.Name())
        destinationFilePath := filepath.Join(bst.DestinationDir, file.Name())
        err = bst.copyFileIfModified(sourceFilePath, destinationFilePath)
        if err != nil {
            return fmt.Errorf("failed to sync file %s: %w", file.Name(), err)
        }
    }
    return nil
}

// copyFileIfModified 如果源文件自上次备份后被修改，则复制到目标目录
func (bst *BackupSyncTool) copyFileIfModified(source, destination string) error {
    sourceFileInfo, err := os.Stat(source)
    if err != nil {
        return fmt.Errorf("failed to get source file info: %w", err)
    }
    destinationFileInfo, err := os.Stat(destination)
    var needCopy bool
    if os.IsNotExist(err) {
        needCopy = true
    } else {
        needCopy = sourceFileInfo.ModTime().After(bst.LastSyncTime) || destinationFileInfo.ModTime().After(bst.LastSyncTime)
    }
    if needCopy {
        err = bst.copyFile(source, destination)
        if err != nil {
            return fmt.Errorf("failed to copy file: %w", err)
        }
    }
    return nil
}

// copyFile 复制文件从源到目标
func (bst *BackupSyncTool) copyFile(source, destination string) error {
    sourceFile, err := os.Open(source)
    if err != nil {
        return fmt.Errorf("failed to open source file: %w", err)
    }
    defer sourceFile.Close()
    destinationFile, err := os.Create(destination)
    if err != nil {
        return fmt.Errorf("failed to create destination file: %w", err)
    }
    defer destinationFile.Close()
    _, err = io.Copy(destinationFile, sourceFile)
    if err != nil {
        return fmt.Errorf("failed to copy file content: %w", err)
    }
    return nil
}

func main() {
    var sourceDir, destinationDir string
    flag.StringVar(&sourceDir, "source", "./source", "source directory path")
    flag.StringVar(&destinationDir, "destination", "./destination", "destination directory path")
    flag.Parse()

    tool := NewBackupSyncTool(sourceDir, destinationDir)
    err := tool.Sync()
    if err != nil {
        log.Fatalf("error during file sync: %v", err)
    }
    fmt.Println("File sync completed successfully.")
}