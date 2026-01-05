package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	var count int
	for {
		count+=1
		best_score := 1000
		random := rand.IntN(100)
		//fmt.Println(random)
		var i int
		var attempts int
		fmt.Println("enter number of attempts: ")
		fmt.Scanf("%d\n", &attempts)
		for {
			var number int
			fmt.Println("enter your guess here : ")
			fmt.Scanf("%d\n", &number)

			if number <= 100 && number >= 0 {
				if number > random {
					fmt.Println("the guess is too high")
					i += 1
				} else if number < random {
					fmt.Println("the guess is too low")
					i += 1
				} else {
					fmt.Println("the guess is correct")
					if count==1{
						best_score= i+1
					}
					if best_score > i  {
						best_score = i+1
					}
					i += 1
					break
				}
			} else {
				fmt.Println("out of range , enter a valid number ")
			}
			if i >= attempts {
				fmt.Println("game over, you lost !!")
				break
			}

		}

		fmt.Println("the number of valid attempts is : ", i)
		
		fmt.Println("do you wanna play again?(y,n)")
		var c rune
		fmt.Scanf("%c\n", &c)
		if c == 'n' {
			fmt.Println("your best score is : ", best_score,"N.B : if you saw a high score , it means you didn't succeed any step")
			fmt.Println("review this quiz (easy,medium,hard)")
			var intention string
			fmt.Scanf("%s\n",intention)
			fmt.Println("your answer was : ",intention)
			break
		}
	}
}
