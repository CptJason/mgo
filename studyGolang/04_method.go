// compiler supports them.
package main

import (
	"fmt"
)

// user defines a user in the program.
type user struct {
	name  string
	email string
}

//
// method有两种类型的receivers: Value & Pointer
//

// notify implements a method with a value receiver.
func (u user) notify() {
	fmt.Printf("User: Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}

// changeEmail implements a method with a pointer receiver.
func (u *user) changeEmail(email string) {
	u.email = email
}

// main is the entry point for the application.
func main() {
	// Values of type user can be used to call methods
	// declared with a value receiver.
	bill := user{"Bill", "bill@email.com"}
	bill.notify()
	bill.changeEmail("bill@gmail.com")

	// Pointers of type user can also be used to methods
	// declared with a value receiver.
	lisa := &user{"Lisa", "lisa@email.com"}
	lisa.notify()
	lisa.changeEmail("lisa@gmail.com")
}
