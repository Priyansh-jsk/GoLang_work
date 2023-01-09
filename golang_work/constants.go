package main

import "fmt"

func main() {
	const myfavlang = "Go"
	const myfavanimal = "Tiger"

	const country, code = "India", 91
	const (
		studentId string  = "stu101"
		salary    float32 = 7.0
	)
	const typedStr string = "Hi"

	fmt.Println(myfavlang, myfavanimal, country, code, studentId, salary, typedStr)
}
