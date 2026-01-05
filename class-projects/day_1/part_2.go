package main

import ("fmt"
		"math/rand/v2")

func main(){
	random := rand.IntN(100)
	//fmt.Println(random)
	var i int
	var attempts int 
	fmt.Println("enter number of attempts : ")
	fmt.Scanf("%d",&attempts)

	for {
		var number int 
		fmt.Println("enter your guess here : ")
		fmt.Scanf("%d",&number)
		
		if number <= 100 && number >= 0 {
			if  number > random {
				fmt.Println("the guess is too high")
				i+=1
			}else if number < random {
				fmt.Println("the guess is too low")
				i+=1
			}else {
				fmt.Println("the guess is correct")
				i+=1
				break
			}
		}else {
			fmt.Println("out of range , enter a valid number ")
		}
		if i>=attempts{
			fmt.Println("game over, you lost !!")
			break
		}
		
	}
	fmt.Println("the number of valid attempts is : ",i)
}
