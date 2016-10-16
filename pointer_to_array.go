package main

import "fmt"

func main() {
    var a [8]int = [8]int{1,2,3,4,5,6,7,8}
    var b [8]int = [8]int{10,20,30,40,50,60,70,80}
    var p *[8]int = &a
    p1, p2 := &a[0], &a[1]
    fmt.Println(p)
    fmt.Println(*p)
    fmt.Println(p1)
    fmt.Println(*p1)
    fmt.Println(p[0])
    fmt.Println(p2)
    fmt.Println(*p2)
    fmt.Println(p[1])
    *p1 = 10
    *p2 = 20
    fmt.Println(a)
    p = &b;
    fmt.Println(*p)
}
