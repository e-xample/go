package main

import "fmt"

                             // 間違い
func swap(x *int, y *int) {  // func swap(x int, y int) {
    tmp := *x                //     tmp := x
    *x = *y                  //     x = y
    *y = tmp                 //     y = tmp
}                            // }

func timesArray(n int, ary *[8]int) {
    for i := 0; i < len(*ary); i++ {
        ary[i] *= n
    }
}

func main() {
    var a int = 10
    var b int = 20
    var c [8]int = [8]int{1,2,3,4,5,6,7,8}
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
    swap(&a, &b)
    timesArray(10, &c)
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
}
