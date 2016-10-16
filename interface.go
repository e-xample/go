package main

import (
	"fmt"
	"math"
)

// Type: Point
type Point interface {
	distance0() float64
}

// 2-dimension
type Point2d struct {
	x, y float64
}

func newPoint2d(x, y float64) *Point2d {
	p := new(Point2d)
	p.x, p.y = x, y
	return p
}

func (p *Point2d) distance0() float64 {
	return math.Sqrt(p.x * p.x + p.y * p.y)
}

// 3-dimension
type Point3d struct {
	x, y, z float64
}

func newPoint3d(x, y, z float64) *Point3d {
	p := new(Point3d)
	p.x, p.y, p.z = x, y, z
	return p
}

func (p *Point3d) distance0() float64 {
	return math.Sqrt(p.x * p.x + p.y * p.y + p.z * p.z)
}

// sum
func sumOfDistance(ary []Point) float64 {
	sum := 0.0
	for _, p := range ary {
		sum += p.distance0()
	}
	return sum
}

func main() {
	a := []Point{
		newPoint2d(0, 0), newPoint2d(10, 10),
		newPoint3d(0, 0, 0), newPoint3d(10, 10, 10),
	}
	fmt.Println(a[0].distance0())
	fmt.Println(a[1].distance0())
	fmt.Println(a[2].distance0())
	fmt.Println(a[3].distance0())
	fmt.Println(sumOfDistance(a))
}


