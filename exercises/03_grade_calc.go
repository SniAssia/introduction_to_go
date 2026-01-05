package main

import "fmt"
		

func main() {
	average :=0.0
	i:=0.0
	for {
		var grady float64
		fmt.Println("enter the grade  ",": ")
		fmt.Scanf("%g",&grady)
		if (grady==-1.0){
			break
		}
		average=average+grady
		i+=1

	}
	average/=i
	fmt.Println("the average is : ",average)

	switch  {
	case average>=90: 
		fmt.Println(average,"A")
	case average<90 && average >= 80 : 
		fmt.Println(average,"B")
	case average<80 && average >= 70 : 
		fmt.Println(average,"C")
	case average<70 && average >= 60 : 
		fmt.Println(average,"D")
	case average < 60 : 
		fmt.Println(average,"E")
	default : 
		fmt.Println("your average is not valid")
	}
	
	


}
