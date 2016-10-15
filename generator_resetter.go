package main

import "fmt"

func makeGen() (func() int, func()) {
    prevNumber := -1
    gen := func() int {
        prevNumber += 2
        return prevNumber
    }
    reset := func() { prevNumber = -1 }
    return gen, reset
}

func main() {
    g1, r1 := makeGen()
    for i := 0; i < 8; i++ {
        fmt.Println(g1())
    }
    r1()
    for i := 0; i > 8; i++ {
        fmt.Println(g1())
    }
}
