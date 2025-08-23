// 代码生成时间: 2025-08-23 13:16:05
package main

import (
    "buffalo"
    "github.com/gobuffalo/buffalo/generators"
    "github.com/gobuffalo/pop/v6"
    "github.com/gobuffalo/buffalo/meta/inflect"
    "github.com/markbates/oncer"
    "log"
    "github.com/gobuffalo/buffalo"
    "github.com/gobuffalo/buffalo/generators/assets/gtm/gtm"
    "github.com/gobuffalo/buffalo/generators/assets/gtm/templates"
)

// CachePolicyGenerator represents the generator for creating a cache policy
type CachePolicyGenerator struct {
    generator.Base
}

// Generate will generate the code for a cache policy
func (c *CachePolicyGenerator) Generate() error {
    if err := c.checkForExistingCode(); err != nil {
        return err
    }
    oncer.Do("cache_policy", func() {
        if err := c.generate(); err != nil {
            log.Fatal(err)
        }
    })
    return nil
}

// generate creates the cache policy code
func (c *CachePolicyGenerator) generate() error {
    // Define the cache policy code
    cachePolicyCode := `
// CachePolicy defines the strategy for caching
type CachePolicy struct {
    // MaxAge defines the maximum age of the cache in seconds
    MaxAge int
    // Private indicates if the cache is private or shared
    Private bool
    // NoStore indicates if the cache should store the data
    NoStore bool
}

// NewCachePolicy creates a new CachePolicy instance
func NewCachePolicy(maxAge int, private bool, noStore bool) *CachePolicy {
    return &CachePolicy{
        MaxAge: maxAge,
        Private: private,
        NoStore: noStore,
    }
}

// GetCacheControlHeader returns the cache control header based on the policy
func (cp *CachePolicy) GetCacheControlHeader() string {
    var header string
    if cp.Private {
        header = "private"
    } else {
        header = "public"
    }
    if cp.NoStore {
        header += ", no-store"
    } else {
        header += ", max-age=" + strconv.Itoa(cp.MaxAge)
    }
    return header
}

`
    // Create a file for the cache policy code
    cFile, err := c.Create("cache_policy.go")
    if err != nil {
        return err
    }
    defer cFile.Close()

    // Write the cache policy code to the file
    if _, err := cFile.Write([]byte(cachePolicyCode)); err != nil {
        return err
    }

    return nil
}

// checkForExistingCode checks if the cache policy code already exists
func (c *CachePolicyGenerator) checkForExistingCode() error {
    // Check if the cache_policy.go file already exists
    if _, err := c.App.Fs.Stat("cache_policy.go"); err == nil {
        return generators.ErrorWithHint(
            "The cache_policy.go file already exists.",
            "You can regenerate this file by deleting the existing file or using the --force flag.",
        )
    }
    return nil
}

func main() {
    c := CachePolicyGenerator{}
    if err := c.Generate(); err != nil {
        log.Fatal(err)
    }
}
