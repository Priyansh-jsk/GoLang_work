package main

import (
	"fmt"
)

// func main() {
// 	fmt.Println("one")
// 	defer fmt.Println("three")
// 	fmt.Println("two")
// }

// Example2
func main() {
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
}
