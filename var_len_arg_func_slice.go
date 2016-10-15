package main

import "fmt"

func foo(x int, y... int) {
    fmt.Println(x, y)
}

func main() {
    a := []int{1, 2, 3}
    b := []int{4, 5, 6}
    foo(0, a...)
    fmt.Println(append(a, b...))
}
