package main

import "fmt"

func facti(n, a int) int {
    if n == 0 {
        return a
    } else {
        return facti(n - 1, a * n)
    }
}

func main() {
    for i := 0; i < 13; i++ {
        fmt.Println(i, ":", facti(i, 1))
    }
}
