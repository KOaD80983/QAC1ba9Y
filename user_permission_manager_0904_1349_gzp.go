// 代码生成时间: 2025-09-04 13:49:07
package main

import (
    "buffalo" // Buffalo web framework
    "github.com/gobuffalo/buffalo-pop" // Buffalo ORM
    "github.com/buffalo/plush" // Buffalo template helpers
    "github.com/buffalo/x/buffalo-plugins" // Buffalo plugins
    "github.com/buffalo/x/www" // Buffalo web assets
    "log"
)

// User represents a user with permissions
type User struct {
    ID   uint   "db:default_id"
anonymous bool   "db:-"
    Roles  []Role `db:'belongs_to:roles'"
}

// Role represents a role for a user
type Role struct {
    ID      uint   "db:default_id"
    Name    string `db:"size:255"`
    Users   []User `db:'has_many:users'`
    
    // Permissions is a slice of permissions attached to this role
    Permissions []Permission `db:'has_many:permissions,join_fk:role_id'`
}

// Permission represents a permission that can be granted to a role
type Permission struct {
    ID         uint   "db:default_id"
    RoleID     uint   "db:"index"`
    Resource   string `db:"size:255"`
    Action     string `db:"size:255"`
}

// Main function to run the Buffalo application
func main() {
    app := buffalo.Automatic(buffalo.Options{
        Environment: "development",
        ProjectName: "user_permission_manager",
    })
    
    // Define routes for the application
    app.GET("/users/:id/permissions", userPermissionsHandler)
    app.POST("/roles/:id/permissions", addPermissionToRoleHandler)
    app.DELETE("/roles/:id/permissions/:permission_id", removePermissionFromRoleHandler)

    // Start the application
    app.Serve()
}

// userPermissionsHandler returns the permissions for a given user ID
func userPermissionsHandler(c buffalo.Context) error {
    userID, err := c.ParamInt("id")
    if err != nil {
        return c.Error(401, err)
    }
    
    tx := c.Value("tx