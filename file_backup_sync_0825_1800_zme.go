// 代码生成时间: 2025-08-25 18:00:45
// file_backup_sync.go

package main

import (
    "bufio"
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "strings"
    "sync"
    "time"
)

// 文件备份同步工具
type BackupSyncTool struct {
    sourceDir string
    targetDir string
    syncChan  chan struct{}
    once      sync.Once
}

// NewBackupSyncTool 创建一个新的文件备份同步工具实例
func NewBackupSyncTool(sourceDir, targetDir string) *BackupSyncTool {
    return &BackupSyncTool{
        sourceDir: sourceDir,
        targetDir: targetDir,
        syncChan:  make(chan struct{}, 1),
    }
}

// Start 开始同步文件
func (t *BackupSyncTool) Start() {
    go t.syncFiles()
}

// syncFiles 同步文件的逻辑
func (t *BackupSyncTool) syncFiles() {
    for {
        select {
        case <-t.syncChan:
            err := t.backupFiles()
            if err != nil {
                log.Printf("Error during file backup: %v", err)
            }
        case <-time.After(10 * time.Minute): // 每10分钟同步一次文件
            err := t.backupFiles()
            if err != nil {
                log.Printf("Error during file backup: %v", err)
            }
        }
    }
}

// backupFiles 备份文件
func (t *BackupSyncTool) backupFiles() error {
    // 遍历源目录
    err := filepath.Walk(t.sourceDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() {
            // 构建目标文件路径
            relPath, err := filepath.Rel(t.sourceDir, path)
            if err != nil {
                return err
            }
            targetPath := filepath.Join(t.targetDir, relPath)

            // 创建目标目录
            targetDir := filepath.Dir(targetPath)
            if err := os.MkdirAll(targetDir, 0755); err != nil {
                return err
            }

            // 读取源文件内容
            file, err := os.Open(path)
            if err != nil {
                return err
            }
            defer file.Close()

            // 创建目标文件
            dest, err := os.Create(targetPath)
            if err != nil {
                return err
            }
            defer dest.Close()

            // 复制文件内容
            _, err = io.Copy(dest, file)
            return err
        }
        return nil
    })
    return err
}

// TriggerSync 触发文件同步
func (t *BackupSyncTool) TriggerSync() {
    t.once.Do(func() {
        t.syncChan <- struct{}{}
    })
}

func main() {
    // 示例使用
    source := "/path/to/source"
    target := "/path/to/target"
    tool := NewBackupSyncTool(source, target)
    tool.Start()
    // 触发同步
    tool.TriggerSync()
    // 等待用户输入以退出程序
    fmt.Println("Press enter to exit...")
    fmt.Scanln()
}
