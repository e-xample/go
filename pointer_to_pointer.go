package main

import "fmt"

func main() {
    var i int = 100
    var p *int
    var q **int
    p = &i
    q = &p
    fmt.Println(p)
    fmt.Println(*p)
    fmt.Println(q)
    fmt.Println(*q)
    fmt.Println(**q)
}
