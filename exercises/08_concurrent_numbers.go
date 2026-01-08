package main

import ("fmt"
"sync")

func squareNumber(a int) {
	s := a * a
	fmt.Println("square of ", a, "is : ",s)
}
func main() {
	var wg sync.WaitGroup
	inte1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := range inte1 {
		wg.Add(1)
		go func(){
			squareNumber(inte1[i])
			defer wg.Done()
		}()
	}
	wg.Wait()

}
