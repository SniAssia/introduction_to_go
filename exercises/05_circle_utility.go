package main

import (
	"fmt"
	"math"
	
)
func main(){
	c := Circle{radius : 2.89}
	fmt.Println("the area is : ",c.calc_area())
	fmt.Println("the perime is : ",c.perime())
	

}
type Circle struct {
	radius float64
}
func (c Circle) calc_area() float64{
	return (c.radius*c.radius) * 2 * math.Pi 
}
func (c *Circle) scale(factor float64){
	c.radius+=factor
}
func (c Circle) perime() float64{
	return 2*math.Pi*c.radius
}
	
