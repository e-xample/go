package main

import (
    "fmt"
    "math"
)

type Point struct {
    x, y float64
}

type Point3d struct {
    x, y, z float64
}

func newPoint(x, y float64) *Point {
    p := new(Point)
    p.x, p.y = x, y
    return p
}

func newPoint3d(x, y, z float64) *Point3d {
    p := new(Point3d)
    p.x, p.y, p.z = x, y, z
    return p
}

func (p *Point) distance(q *Point) float64 {
    dx := p.x - q.x
    dy := p.y - q.y
    return math.Sqrt(dx * dx + dy * dy)
}

func (p *Point3d) distance(q *Point3d) float64 {
    dx := p.x - q.x
    dy := p.y - q.y
    dz := p.z - q.z
    return math.Sqrt(dx * dx + dy * dy + dz * dz)
}

func main() {
    p1 := newPoint(0, 0)
    p2 := newPoint(10, 10)
    q1 := newPoint3d(0, 0, 0)
    q2 := newPoint3d(10, 10, 10)
    fmt.Println(p1.distance(p2))
    fmt.Println(q1.distance(q2))
}
