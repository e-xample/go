package main

import "fmt"

func divMod(x, y int) (int, int) {
    return x / y, x % y
}

/*
func divMod(x, y int) (p, q int) {
    p, q = x / y, x % y
    return 
}
*/

func main() {
    p, q := divMod(10, 3)
    fmt.Println(p)
    fmt.Println(q)
}
