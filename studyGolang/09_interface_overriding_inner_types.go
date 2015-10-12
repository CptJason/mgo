// Sample program to show what happens when the outer and inner
// type implement the same interface.
package main

import (
	"fmt"
)

// notifier is an interface that defined notification
// type behavior.
type notifier interface {
	notify3()
}

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// admin represents an admin user with privileges.
type admin struct {
	user
	level string
}

// notify implements a method that can be called via
// a value of type user.
func (u *user) notify3() {
	fmt.Printf("user: Sending user email To %s<%s>\n",
		u.name,
		u.email)
}

// notify implements a method that can be called via
// a value of type Admin.
func (a *admin) notify3() {
	fmt.Printf("User: Sending Admin Email To %s<%s>\n",
		a.name,
		a.email)
}

// main is the entry point for the application.
func main() {
	// Create an admin user.
	ad := admin{
		user: user{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	// Send the admin user a notification.
	// The embedded inner type's implementation of the
	// interface is NOT "promoted" to the outer type.
	sendNotification3(&ad)

	// We can access the inner type's method directly.
	ad.user.notify3()

	// The inner type's method is NOT promoted.
	ad.notify3()
}

// sendNotification accepts values that implement the notifier
// interface and sends notifications.
func sendNotification3(n notifier) {
	n.notify3()
}