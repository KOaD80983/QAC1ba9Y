// 代码生成时间: 2025-09-20 06:30:00
and follows Go best practices for maintainability and scalability.
*/

package main

import (
    "github.com/gobuffalo/buffalo"
    "github.com/gobuffalo/buffalo/middleware"
    "net/http"
)

// ApiResponseFormatter represents the structure of a formatted API response.
type ApiResponseFormatter struct {
    Status  string      `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data"`
    Errors  []string    `json: