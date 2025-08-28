// 代码生成时间: 2025-08-28 23:28:25
package main

import (
    "bufio"
    "fmt"
    "log"
    "net/http"
    "os"
    "os/exec"
    "strconv"
    "strings"
    "time"
)

// PerformanceTest holds configuration for the test
type PerformanceTest struct {
    URL     string
    Threads int
    Duration time.Duration
}

// NewPerformanceTest creates a new PerformanceTest instance
func NewPerformanceTest(url string, threads int, duration time.Duration) *PerformanceTest {
    return &PerformanceTest{
        URL:     url,
        Threads: threads,
        Duration: duration,
    }
}

// Run starts the performance test
func (pt *PerformanceTest) Run() error {
    // Create a channel to signal when each thread is done
    done := make(chan bool, pt.Threads)
    defer close(done)

    // Start a timer to control the duration of the test
    endTime := time.Now().Add(pt.Duration)

    for i := 0; i < pt.Threads; i++ {
        go func(id int) {
            for time.Now().Before(endTime) {
                // Make a request to the given URL
                resp, err := http.Get(pt.URL)
                if err != nil {
                    log.Printf("Thread %d: Error making request: %v", id, err)
                    return
                }
                defer resp.Body.Close()

                // Check if the response is successful
                if resp.StatusCode != http.StatusOK {
                    log.Printf("Thread %d: Non-200 status code: %d", id, resp.StatusCode)
                    return
                }
            }
            done <- true
        }(i)
    }

    // Wait for all threads to complete or the duration to end
    for i := 0; i < pt.Threads; i++ {
        <-done
    }
    return nil
}

func main() {
    // Define the configuration for the performance test
    url := "http://localhost:3000"
    threads, _ := strconv.Atoi(os.Getenv("THREADS"))
    duration, _ := time.ParseDuration(os.Getenv("DURATION"))
    test := NewPerformanceTest(url, threads, duration)

    // Run the performance test
    if err := test.Run(); err != nil {
        fmt.Printf("Performance test failed: %v
", err)
    } else {
        fmt.Println("Performance test completed successfully.")
    }
}
