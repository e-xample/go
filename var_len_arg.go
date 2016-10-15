package main

import "fmt"

func foo1(a int, args ...int) {
    fmt.Println(a, args)
}

func foo0(args ...int) {
    fmt.Println(args)
}

func main() {
    foo0()
    foo0(1)
    foo0(1, 2)
    foo0(1, 2, 3)
    foo1(1)
    foo1(1, 2)
    foo1(1, 2, 3)
    foo1(1, 2, 3, 4)
}
