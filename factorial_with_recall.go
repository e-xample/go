package main

import "fmt"

func fact(n int) int {
    if n == 0 {
        return 1
    } else {
        return n * fact(n - 1)
    }
}

func main() {
    for i := 0; i < 13; i++ {
        fmt.Println(i, ":", fact(i))
    }
}
