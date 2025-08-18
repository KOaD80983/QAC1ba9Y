// 代码生成时间: 2025-08-18 17:14:09
package main

import (
    "buffalo"
    "github.com/gobuffalo/buffalo-pop/v2/pop"
    "github.com/gobuffalo/nulls"
    "log"
    "os"
)

// Notification contains the message to be sent.
type Notification struct {
    ID        nulls.Int    `db:"id"`
    Message   string       `db:"message"`
    CreatedAt nulls.Time   `db:"created_at"`
    UpdatedAt nulls.Time   `db:"updated_at"`
}

// NotificationService handles business logic for notifications.
type NotificationService struct {
    DB *pop.Connection
}

// NewNotificationService initializes a new NotificationService with a database connection.
func NewNotificationService(db *pop.Connection) *NotificationService {
    return &NotificationService{DB: db}
}

// CreateNotification inserts a new notification into the database.
func (s *NotificationService) CreateNotification(message string) (*Notification, error) {
    if message == "" {
        return nil, buffalo.NewError("Message cannot be empty")
    }
    
    n := &Notification{Message: message}
    err := s.DB.Create(n)
    if err != nil {
        return nil, buffalo.NewError("Failed to create notification: " + err.Error())
    }
    
    return n, nil
}

// Main function to run the Buffalo application.
func main() {
    app := buffalo.Automatic()
    defer app.Close()

    // Set up the database connection.
    db, err := app.DB().Begin
    if err != nil {
        log.Fatal(err)
    }
    defer db.Rollback()

    // Initialize the NotificationService with the database connection.
    notificationService := NewNotificationService(app.DB())

    // Define a route for creating notifications.
    app.POST("/notifications", func(c buffalo.Context) error {
        message := c.Request().FormValue("message")
        notification, err := notificationService.CreateNotification(message)
        if err != nil {
            return c.Error(400, err)
        }
        
        return c.Render(201, r.JSON(notification))
    })

    // Start the server.
    if err := app.Serve(); err != nil {
        log.Fatal(err)
    }
}