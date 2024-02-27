package main

import (
	"fmt"
	"time"
)

type Notification struct {
	UserID  int
	Message string
}

func sendEmailAsync(userID int, message string) {
	time.Sleep(2 * time.Second)
	fmt.Printf("Email notification sent to user %d: %s\n", userID, message)
}

func main() {
	notifications := []Notification{
		{101, "Your order has been confirmed."},
		{102, "Your account has been created."},
		{103, "Your payment was sucessful."},
	}

	for _, notification := range notifications {
		go sendEmailAsync(notification.UserID, notification.Message)
	}
	fmt.Println("Main application continues...")
	time.Sleep(3 * time.Second)
	fmt.Println("Main application finished.")
}
