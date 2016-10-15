package main

import "fmt"

func main() {
    var a = "1234567890"
    for i := 0; i <= len(a); i++ {
        var s string = a[i:]
        fmt.Println(s)
        fmt.Println(len(s))
    }
}
