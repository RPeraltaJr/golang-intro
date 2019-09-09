package main

import "fmt"

func add(x, y float32) float32 {
	return x + y
}

// * Defining a constant
const x int = 5

func multiple(a, b string) (string, string) {
	return a, b
}

func main() {

	// * Shorthand (option #1)
	// var num1 float64 = 5.6
	// var num2 float64 = 9.5

	// * Shorthand (option #2)
	// var num1, num2 float64 = 5.6, 9.5

	// * Shorthand (option #3)
	// num1, num2 := 5.6, 9.5

	w1, w2 := "Hey", "there!"
	fmt.Println(multiple(w1, w2))

	// var a int = 62
	// var b float64 = float64(a)

	// x := a // x will be type int

}
