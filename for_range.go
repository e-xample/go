package main

import "fmt"

func main() {
    a := []int{1, 2, 3, 4, 5}
    b := map[string]int{"foo": 10, "bar": 20, "baz": 30}
    for i, v := range a {
        fmt.Println(i, v)
    }
    for k, v := range b {
        fmt.Println(k, v)
    }
}
