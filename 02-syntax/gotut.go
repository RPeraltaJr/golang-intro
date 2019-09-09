package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func foo() {
	fmt.Println("The square root of 4 is", math.Sqrt(4))
}

func randomNumber() {
	rand.Seed(time.Now().UnixNano()) // prevents getting the same number from rand.Intn() method
	fmt.Println("A number from 1-100", rand.Intn(100))
}

// * main()is always going to run
func main() {
	foo()
	randomNumber()
}

/*
	* Documentation:
	$ godoc math/rand Intn
*/
