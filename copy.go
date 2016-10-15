package main

import "fmt"

func main() {

	a := []int{1,2,3,4,5,6,7,8}
	b := make([]int, 10)
	c := make([]int, 4)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	copy(b,a)
	copy(c,a)
	fmt.Println(b)
	fmt.Println(c)
}
