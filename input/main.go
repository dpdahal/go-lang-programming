package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// input value from user
	//var name string
	//print("Enter your name: ")
	////_, _ = fmt.Scanln(&name)
	//fmt.Scanln(&name)
	//fmt.Println("Hello", name)

	// Using bufio for more advanced input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your age: ")
	age, _ := reader.ReadString('\n') // Reads input until the newline
	fmt.Printf("You are %s years old!\n", age)

}
