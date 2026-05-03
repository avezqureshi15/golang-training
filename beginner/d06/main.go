package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type InvalidDimensionError struct {
	Shape string
	Msg   string
}

func (e InvalidDimensionError) Error() string {
	return fmt.Sprintf("%s error: %s", e.Shape, e.Msg)
}

type Circle struct {
	Radius float64
}

func NewCircle(r float64) (*Circle, error) {
	if r <= 0 {
		return nil, InvalidDimensionError{"Circle", "radius must be > 0"}
	}
	return &Circle{Radius: r}, nil
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Width   float64
	Heights float64
}

func NewRectangle(w float64, h float64) (*Rectangle, error) {

	if w <= 0 || h <= 0 {
		return nil, InvalidDimensionError{"Rectangle", "width/height must be > 0"}
	}

	return &Rectangle{Width: w, Heights: h}, nil
}

func (r * Rectangle) Area() float64{
	return r.Width * r.Heights
}

func (r * Rectangle) Perimeter() float64{
	return 2 * (r.Heights + r.Heights)
}

type Triangle struct{
	A float64
	B float64
	C float64 
}

func NewTriangle(a float64, b float64, c float64) (*Triangle,error){
	if a <= 0 || b<= 0 || c<=0 {
		return nil, InvalidDimensionError{"Trianle","sides must be >0"}
	}
	if a+b <= c || a+c <= b || b+c <= a{
		return nil , InvalidDimensionError{"Triangle","invalid triangle sides"}
	}
	return &Triangle{A:a,B:b,C:c},nil 
}

func (t *Triangle) Area() float64 {
	s := (t.A + t.B + t.C) / 2
	return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}

func (t *Triangle) Perimeter() float64{
	return t.A + t.B + t.C
}

func TotalArea(shapes []Shape) float64{
	total := 0.0
	for _, s:= range shapes{
		total += s.Area()
	}
	return total
}

func main(){
	c,_ := NewCircle(5)
	r,_ := NewRectangle(4,6)
	t,_ := NewTriangle(3,4,5)

	shapes := []Shape{c,r,t}

	fmt.Println("Total Area ",TotalArea(shapes))
}