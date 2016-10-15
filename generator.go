package main

import "fmt"

var prevNumber = -1

func genOddNumber() int {
    prevNumber += 2
    return prevNumber
}

func main() {
    for i := 0; i < 8; i++ {
        fmt.Println(genOddNumber())
    }
}
