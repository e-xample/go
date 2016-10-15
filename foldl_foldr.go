package main

import "fmt"

func add(x, y int) int {
    return x + y
}

func foldl(f func(int, int) int, a int, ary []int) int {
	for _,v := range ary {
		a = f(a, v)
	}
	return a
}

func foldr(f func(int, int) int, a int, ary []int) int {
	for i := len(ary)-1; i >= 0; i-- {
		a = f(a, ary[i])
	}
	return a
}

func main() {
    a := []int{1,2,3,4,5,6,7,8}
    fmt.Println(foldl(add, 0, a))
    fmt.Println(foldr(add, 0, a))
}
