package main

import "fmt"

func isEven(x int) bool {
    return x % 2 == 0
}

func isOdd(x int) bool {
    return x % 2 != 0
}

func removeIf(f func(int) bool, ary []int) []int {
	buff := make([]int, 0)
	for _,v := range ary {
		if !f(v) {
			buff = append(buff, v)
		}
	}
	return buff
}

func main() {
    a := []int{1,2,3,4,5,6,7,8}
    d := removeIf(isEven, a)
    e := removeIf(isOdd, a)
    fmt.Println(d)
    fmt.Println(e)
}
