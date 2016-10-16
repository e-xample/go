
package main

import "fmt"

type Num interface {
    number()
}

type Int int

func (n Int) number() {}

type Real float64

func (n Real) number() {}

func sumOfNum(ary []Num) (Int, Real) {
    var sumi Int  = 0
    var sumr Real = 0.0
    for _, x := range ary {
        switch v := x.(type) {
        case Int:  sumi += v
        case Real: sumr += v
        }
    }
    return sumi, sumr
}

func main() {
    var ary []Num = []Num{
        Int(1), Real(1.1), Int(2), Real(2.2), Int(3), Real(3.3),
    }
    a, b := sumOfNum(ary)
    fmt.Println(a, b)
}
