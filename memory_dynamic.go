package main

import "fmt"

func main() {
    var p *int = new(int)
    var q *float64 = new(float64)
    var a *[8]int = new([8]int)
    fmt.Println(p)
    fmt.Println(*p)
    fmt.Println(q)
    fmt.Println(*q)
    fmt.Println(a)
    fmt.Println(*a)
    *p = 100
    *q = 1.2345
    a[0] = 10
    a[7] = 80
    fmt.Println(*p)
    fmt.Println(*q)
    fmt.Println(*a)
}
