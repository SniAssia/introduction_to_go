package main

import (
	"fmt"
	"log"
	"math"
)

type Circle struct {
	radius float64
}

// every struct implement empty interface
type Rectangle struct {
	height float64
	width  float64
}
type Triangle struct {
	sidea float64 
	sideb float64
	sidec float64

}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) Perime() float64 {
	return math.Pi * 2 * c.radius

}

func (c Triangle) Area() float64 {
	return c.sidea * c.sideb * c.sidec
}
func (c Triangle) Perime() float64 {
	return c.sidea + c.sideb + c.sidec 

}

func (c Rectangle) Area() float64 {
	return c.width * c.height
}
func (c Rectangle) Perime() float64 {
	return c.width * 2 + c.height* 2 

}

type Shape interface {
	Area() float64
	Perime() float64
}

func printshapedetails(s Shape) {
	switch  s.(type) {
	case Circle : 
		fmt.Println("the area is : ",s.Area())
		fmt.Println("the perimeter is : ",s.Perime())
	
	case Rectangle : 
		fmt.Println("the area is : ",s.Area())
		fmt.Println("the perimeter is : ",s.Perime())
	case Triangle : 
		fmt.Println("the area is : ",s.Area())
		fmt.Println("the perimeter is : ",s.Perime())
	default :
		log.Fatal("the argument passed doesnt implements the current interface")
	}
	


}
func total_area(shapes ...Shape) {
	fmt.Println("ddd")
}
func main() {
	c := Circle{23.9}
	printshapedetails(c)

	t := Triangle{12.8,78.9,90.9}
	printshapedetails(t)

	r := Rectangle{23.9,34.9}
	printshapedetails(r)

	}

}
