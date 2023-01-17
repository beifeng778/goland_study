package main

import "fmt"

func main() {
	var num1 *int
	fmt.Println(num1)
	num2 := 5
	num1 = &num2
	fmt.Println(num1)
	fmt.Println(&num2)
	fmt.Println(&num1)

}
