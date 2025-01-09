package main

import "fmt"

func main() {
	fmt.Println("Integers Data Types")
	var a int = 10
	var b int = 20
	fmt.Println("a:", a, "b:", b)
	fmt.Println("Float Data Types")
	var c float32 = 10.5
	var d float32 = 20.5
	fmt.Println("c:", c, "d:", d)
	fmt.Println("Complex Data Types")
	var e complex64 = 10 + 5i
	var f complex64 = 20 + 5i
	println("e:", e, "f:", f)

	fmt.Println("String Data Types")
	var g string = "Hello"
	var h string = "World"
	fmt.Println("g:", g, "h:", h)

	fmt.Println("Boolean Data Types")
	var i bool = true
	var j bool = false
	fmt.Println("i:", i, "j:", j)

	fmt.Println("Array Data Types")
	var k [5]int = [5]int{1, 2, 3, 4, 5}
	var l [5]string = [5]string{"a", "b", "c", "d", "e"}
	fmt.Println("k:", k, "l:", l)

	fmt.Println("Slice Data Types")
	var m []int = []int{1, 2, 3, 4, 5}
	var n []string = []string{"a", "b", "c", "d", "e"}
	fmt.Println("m:", m, "n:", n)

	fmt.Println("Map Data Types")
	var o map[string]int = map[string]int{"a": 1, "b": 2, "c": 3}
	var p map[string]string = map[string]string{"a": "apple", "b": "ball", "c": "cat"}
	fmt.Println("o:", o, "p:", p)

}
