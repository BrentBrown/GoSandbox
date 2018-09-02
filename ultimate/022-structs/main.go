// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct type to maintain information about a user (name, email and age).
// Create a value of this type, initialize with values and display each field.
//
// Declare and initialize an anonymous struct type with the same three fields. Display the value.
package main

// Add imports.
import "fmt"

// user represents a user in the system.
type user struct {
	name  string
	email string
	age   int
}

func main() {

	// Declare variable of type user and init using a struct literal.
	u := user{
		name:  "John",
		email: "j@example.com",
		age:   42,
	}

	// Display the field values.
	fmt.Println(u)

	// Declare a variable using an anonymous struct.
	uu := struct {
		name  string
		email string
		age   int
	}{"Jane", "jn@example.com", 42}

	// Display the field values.
	fmt.Println(uu)
}
