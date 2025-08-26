// 代码生成时间: 2025-08-26 08:12:33
package main

import (
    "bufio"
    "fmt"
    "os"
    "log"
    "path/filepath"
    "strings"
)

// LogEntry represents a single log entry.
type LogEntry struct {
    Timestamp string // Log timestamp
    Level     string // Log level (INFO, ERROR, etc.)
    Message   string // Log message
}

// ParseLogFile parses a log file and returns a slice of LogEntry structs.
func ParseLogFile(filePath string) ([]LogEntry, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, fmt.Errorf("error opening file: %w", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var logEntries []LogEntry

    for scanner.Scan() {
        line := scanner.Text()
        // Split the line into timestamp, level, and message
        parts := strings.Split(line, " ")
        if len(parts) < 3 {
            continue // Skip lines that don't have enough parts
        }

        timestamp := parts[0] + " " + parts[1]
        level := parts[2]
        message := strings.Join(parts[3:], " ")

        logEntry := LogEntry{Timestamp: timestamp, Level: level, Message: message}
        logEntries = append(logEntries, logEntry)
    }

    if err := scanner.Err(); err != nil {
        return nil, fmt.Errorf("error reading file: %w", err)
    }

    return logEntries, nil
}

// PrintLogEntries prints the log entries to the console.
func PrintLogEntries(logEntries []LogEntry) {
    for _, entry := range logEntries {
        fmt.Printf("%s [%s] %s
", entry.Timestamp, entry.Level, entry.Message)
    }
}

func main() {
    // Replace with your log file path
    logFilePath := "path/to/your/logfile.log"

    logEntries, err := ParseLogFile(logFilePath)
    if err != nil {
        log.Fatalf("error parsing log file: %s", err)
    }

    PrintLogEntries(logEntries)
}
