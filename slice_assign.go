package main

import "fmt"

func main() {
    var a []int
    b := []int{1,2,3,4,5,6,7,8}
    a = b
    fmt.Println(a)
    fmt.Println(b)
    a[0] = 10
    b[7] = 80
    fmt.Println(a)
    fmt.Println(b)

    a = append(a, 100)
    fmt.Println(a)
    fmt.Println(b)
    a[0] = 200
    fmt.Println(a)
    fmt.Println(b)
}
