package main

import "fmt"
// we only use ch to send data 
// to receive we use ch <- chan int
// bidirectonel channel ch chan int 
func generatePrimes(n int,ch chan<- int){
	defer close(ch)
	primes := make([]bool, n)
	primes[0] = false
	primes[1]=false 

	for i := 2; i < n; i++  {
		primes[i]=true
	}
	for i := 2; i < n; i++ {
		
		
		if primes[i] == true{
			for j := i*i;j<n;j+=i{
				primes[j]=false
			}

		}
	}
	for i := 0; i < n; i++ {
		if primes[i]==true{
			
			ch <- i
		}
	}

}

func printprimes(ch <- chan int){
	for elem := range ch {
		fmt.Println("the current number is : ",elem)

	}


}
func main(){
	ch := make(chan int)


	go generatePrimes(15,ch)
	printprimes(ch)

}
