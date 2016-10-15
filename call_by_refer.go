package main

import "fmt"

func foo(ary []int) {
    ary[0] *= 10
}

func main() {
    a := []int{1, 2, 3, 4}
    fmt.Println(a)
    foo(a)
    fmt.Println(a)
}
