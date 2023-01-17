package main

import "fmt"

func main() {
	s := make([]int, 3, 4)
	s[0] = 1
	s[1] = 1
	s[2] = 1
	s[3] = 1

	fmt.Println(s)

}
