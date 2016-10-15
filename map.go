package main

import "fmt"

var a map[string]int = map[string]int{"foo": 10, "bar": 20}

func main() {
    fmt.Println(a)
    fmt.Println(len(a))
    fmt.Println(a["foo"])
    a["foo"] = 100
    fmt.Println(a["foo"])
    a["baz"] = 30
    fmt.Println(a["baz"])
    fmt.Println(a)
    fmt.Println(len(a))
}
