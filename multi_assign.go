package main

import "fmt"

func main() {
    var a, b int
    c := make([]int, 4)
    a, b = 10, 20
    c[0], c[1], c[2], c[3] = 10, 20, 30, 40
    fmt.Println(a, b)
    fmt.Println(c)
    a, b = b, a
    fmt.Println(a, b)
    c[0], c[3] = c[3], c[0]
    fmt.Println(c)
}
