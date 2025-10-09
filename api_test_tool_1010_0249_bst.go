// 代码生成时间: 2025-10-10 02:49:25
// api_test_tool.go
package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
)

// ApiResponse 定义API响应的结构
type ApiResponse struct {
    Status  string `json:"status"`
    Message string `json:"message"`
    Data    interface{} `json:"data"`
}

// TestAPI 测试API的功能
func TestAPI(url, method string, body []byte) *ApiResponse {
    var response *http.Response
    var err error
    
    // 根据请求方法创建HTTP请求
    switch method {
    case "GET":
        response, err = http.Get(url)
    case "POST":
        response, err = http.Post(url, "application/json", bytes.NewBuffer(body))
    case "PUT":
        response, err = http.Put(url, "application/json", bytes.NewBuffer(body))
    case "DELETE":
        response, err = http.Delete(url)
    default:
        return &ApiResponse{Status: "error", Message: "Invalid method"}
    }
    
    // 错误处理
    if err != nil {
        return &ApiResponse{Status: "error", Message: err.Error()}
    }
    defer response.Body.Close()
    
    // 读取响应体
    var responseData ApiResponse
    if err := json.NewDecoder(response.Body).Decode(&responseData); err != nil {
        return &ApiResponse{Status: "error", Message: err.Error()}
    }
    
    return &responseData
}

// main 程序入口
func main() {
    APIMethods := []string{"GET", "POST", "PUT", "DELETE"}
    APIURL := "http://example.com/api"
    requestData := []byte(`{"key": "value"}`)

    for _, method := range APIMethods {
        fmt.Printf("Testing %s method...
", method)
        result := TestAPI(APIURL, method, requestData)
        
        // 打印响应状态，消息和数据
        fmt.Printf("Status: %s, Message: %s, Data: %+v
", result.Status, result.Message, result.Data)
    }
}
