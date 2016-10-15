package main

import "fmt"

func makeGen() func() int {
	prevNumber := -1
	return func() int {
		prevNumber += 2
		return prevNumber
	}
}

func main() {
    g1 := makeGen()
    for i := 0; i < 8; i++ {
        fmt.Println(g1())
    }
    g2 := makeGen()
    for i := 0; i < 8; i++ {
        fmt.Println(g2())
    }
}
