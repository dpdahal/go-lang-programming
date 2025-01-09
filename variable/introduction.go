package main

import "fmt"

func main() {
	// What is variable?
	// A variable is a storage location for holding a value.
	// The value of a variable can be changed and reused many times.
	// The type of a variable determines how much space it occupies in storage and how the bit pattern stored is interpreted.
	// The type of a variable is defined by the keyword var followed by the name of the variable and the type.
	// rule for naming variables:
	// 1. Variable names must start with a letter or an underscore.
	// 2. Variable names cannot start with a number.
	// 3. Variable names can only contain alpha-numeric characters and underscores (A-z, 0-9, and _ ).
	// 4. Variable names are case-sensitive (name, Name, and NAME are three different variables).
	// 5. Variable names should be descriptive and meaningful.
	// 6. Variable names should not be a keyword.
	// 7. Variable names should not contain any spaces.
	// 8. Variable names should not contain any special characters.

	// Declaring a variable
	// var name type
	var name string = "John Doe"
	fmt.Println(name)
	fmt.Println("Hello, World!")
	//name := "Hello Sophia"
	//fmt.Println(name)
	var Names string = "John Doe"
	var names string = "Hari"
	println(Names)
	println(names)
}
