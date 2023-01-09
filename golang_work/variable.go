package main

import "fmt"

func main() {
	var arr1 = [...]int{1, 2, 3}
	arrdemo := [3]int{10, 20, 30}
	var a bool = true
	var b int = 29
	var c float32 = 3.14
	var d string = "priyansh"
	fmt.Println(arr1)
	fmt.Println(arrdemo[1])
	fmt.Println("Bool", a)
	fmt.Println("Int val", b)
	fmt.Println("float value", c)
	fmt.Println("String value", d)
}
