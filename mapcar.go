package main

import "fmt"

func square(x int) int {
	return x * x
}

func cube(x int) int {
	return x * x * x
}

func mapcar(f func(int) int, ary []int) []int {
	buff := make([]int, len(ary))
	for i,v := range ary {
		buff[i] = f(v)
	}
	return buff
}

func main() {
    a := []int{1,2,3,4,5,6,7,8}
    b := mapcar(square, a)
    c := mapcar(cube, a)
    fmt.Println(b)
    fmt.Println(c)
}
