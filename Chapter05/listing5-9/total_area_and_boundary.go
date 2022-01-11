package main

import (
	"fmt"
	"math"
)

var pentagonAreaFactor = 0.25 * math.Sqrt(5*(5+2*math.Sqrt(5)))

type Shape interface {
	GetAreaInSqMeters() float64
	GetPerimeterInMeters() float64
}

type Circle struct {
	RadiusMeters float64
}

func (c Circle) GetAreaInSqMeters() float64 {
	return math.Pi * c.RadiusMeters * c.RadiusMeters
}

func (c Circle) GetPerimeterInMeters() float64 {
	return math.Pi * 2 * c.RadiusMeters
}

type Rectangle struct {
	WidthMeters, HeightMeters float64
}

func (r Rectangle) GetAreaInSqMeters() float64 {
	return r.WidthMeters * r.HeightMeters
}

func (r Rectangle) GetPerimeterInMeters() float64 {
	return 2 * (r.WidthMeters + r.HeightMeters)
}

type IsoscelesTriangle struct {
	BaseMeters, HeightMeters float64
}

func (t IsoscelesTriangle) GetAreaInSqMeters() float64 {
	return 0.5 * t.BaseMeters * t.HeightMeters
}

func (t IsoscelesTriangle) GetPerimeterInMeters() float64 {
	slantSide := math.Sqrt(t.BaseMeters*t.BaseMeters + t.HeightMeters*t.HeightMeters)
	return t.BaseMeters + slantSide + slantSide
}

type RegularPentagon struct {
	SideLengthMeters float64
}

func (p RegularPentagon) GetAreaInSqMeters() float64 {
	return pentagonAreaFactor * p.SideLengthMeters * p.SideLengthMeters
}

func (p RegularPentagon) GetPerimeterInMeters() float64 {
	return 5 * p.SideLengthMeters
}

func GetTotalAreaInSqMeters(shapes []Shape) (totalAreaInSqMeters float64) {
	for _, shape := range shapes {
		totalAreaInSqMeters += shape.GetAreaInSqMeters()
	}
	return
}

func GetTotalPerimeterInMeters(shapes []Shape) (totalPerimeterInMeters float64) {
	for _, shape := range shapes {
		totalPerimeterInMeters += shape.GetPerimeterInMeters()
	}
	return
}

func main() {
	circle := Circle{RadiusMeters: 10.5}
	rectangle1 := Rectangle{WidthMeters: 5.5, HeightMeters: 2.2}
	rectangle2 := Rectangle{WidthMeters: 4, HeightMeters: 6.5}
	isosTriangle := IsoscelesTriangle{BaseMeters: 10, HeightMeters: 4.2}
	rPentagon := RegularPentagon{SideLengthMeters: 20.4}

	shapes := []Shape{circle, rectangle1, rectangle2, isosTriangle, rPentagon}
	fmt.Printf("Total area: %.2f square meters\n", GetTotalAreaInSqMeters(shapes))
	fmt.Printf("Total boundary length: %.2f meters\n", GetTotalPerimeterInMeters(shapes))
}
