package main

import "fmt"

type vehicle struct {
	make  string
	model string
	year  int
}
type insurable interface {
	caclulateinsurance() int
}
type printable interface {
	details()
}
type car struct {
	vehicle
	numberofdoors int
}

func (c car) caclulateinsurance() int {
	return c.numberofdoors * c.year
}
func (c car) details() {
	fmt.Println("the make is : ",c.make)
	fmt.Println("the model is : ",c.model)
	fmt.Println("the year of creation is : ",c.year)

}
type truck struct {
	vehicle
	playloadcap int 

}

func (c truck) caclulateinsurance() int {
	return c.playloadcap * c.year
}
func (c truck) details() {
	fmt.Println("the make is : ",c.make)
	fmt.Println("the model is : ",c.model)
	fmt.Println("the year of creation is : ",c.year)
	fmt.Println("the insurance is : ",c.caclulateinsurance())
}

func printall(p []printable){
	for _,v := range p {
		v.details()
	}
}
func main() {
	v := vehicle{"make","model",1900}
	c0 := car{v,4}
	c1 := car{v,2}
	c2 := car{v,1}

	c3 := car{v,3}
	c4 := car{v,9}
	t1 := truck{v,23}
	t2 := truck{v,23}
	vehicles := []printable{c1,c2,c3,c4,c0,t1,t2}
	printall(vehicles)

}
