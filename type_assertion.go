package main

import "fmt"

// 整数の合計値を求める
func sumOfInt(ary []interface{}) int {
    sum := 0
    for _, x := range ary {
        v, ok := x.(int)
        if ok { sum += v }
    }
    return sum
}

// 実数の合計値を求める
func sumOfFloat(ary []interface{}) float64 {
    sum := 0.0
    for _, x := range ary {
        v, ok := x.(float64)
        if ok { sum += v }
    }
    return sum
}

func main() {
    a := []interface{}{1, 1.1, "abc", 2, 2.2, "def", 3, 3.3}
    fmt.Println(sumOfInt(a))
    fmt.Println(sumOfFloat(a))
}
