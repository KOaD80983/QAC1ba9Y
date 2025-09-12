// 代码生成时间: 2025-09-13 06:54:25
package main

import (
    "encoding/hex"
    "fmt"
    "golang.org/x/crypto/bcrypt"
    "log"
    "os"
)

// PasswordTool 提供密码加密和解密功能
type PasswordTool struct{}

// Encrypt 将明文密码加密
func (tool PasswordTool) Encrypt(password string) (string, error) {
    // 生成密码的哈希值（bcrypt加密）
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return hex.EncodeToString(bytes), nil
}

// Decrypt 将加密密码解密成明文
// 注意：bcrypt加密是不可逆的，因此这里的解密实际上是验证密码
func (tool PasswordTool) Decrypt(hashedPassword, password string) bool {
    // 将十六进制字符串转换回字节切片
    byteHash := []byte(hashedPassword)
    err := bcrypt.CompareHashAndPassword(byteHash, []byte(password))
    return err == nil
}

func main() {
    tool := PasswordTool{}

    // 示例密码
    password := "mysecretpassword"

    // 加密密码
    encryptedPassword, err := tool.Encrypt(password)
    if err != nil {
        log.Fatalf("Error encrypting password: %s", err)
    }
    fmt.Printf("Encrypted password: %s
", encryptedPassword)

    // 尝试解密密码
    isMatch := tool.Decrypt(encryptedPassword, password)
    if isMatch {
        fmt.Println("The password matches the encrypted password.")
    } else {
        fmt.Println("The password does not match the encrypted password.")
    }
}

// Usage: 该程序提供了一个简单的密码加密和解密（验证）工具。
// 运行程序将展示如何对密码进行加密并验证明文密码是否与加密后的密码匹配。
