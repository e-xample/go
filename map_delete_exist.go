package main

import "fmt"

func main() {
    a := make(map[string]int)
    a["foo"] = 10
    a["bar"] = 20
    a["baz"] = 30
    fmt.Println(a)
    delete(a, "baz")
    v, ok := a["baz"]
    fmt.Println(v)
    fmt.Println(ok)
    fmt.Println(a)
}
