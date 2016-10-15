package main

import "fmt"

func main() {
    a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    sum1 := 0
    for i := 0; i < len(a); i++ {
        sum1 += a[i]
    }
    fmt.Println(sum1)
    sum2 := 0
    for _, v := range a {
        sum2 += v
    }
    fmt.Println(sum2)
}
