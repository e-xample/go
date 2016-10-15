package main

import "fmt"

func square(x int) int {
	return x * x
}

func cube(x int) int {
	return x * x * x
}

func mapcar(f func(int) int) func([]int) []int {
	return func(ary []int) []int {
		buff := make([]int, len(ary))
		for i,v := range ary {
			buff[i] = f(v)
		}
		return buff
	}
}

func main() {
    a := []int{1,2,3,4,5,6,7,8}
    squareAry := mapcar(square)
    cubeAry := mapcar(cube)
    fmt.Println(squareAry(a))
    fmt.Println(cubeAry(a))
    fmt.Println(mapcar(square)(a))
    fmt.Println(mapcar(cube)(a))
}
