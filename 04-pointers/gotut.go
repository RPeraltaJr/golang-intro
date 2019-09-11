package main

import "fmt"

// * Use ampersand (&) to point to the memory address
// * Use asterisk (*) to get the value

func main() {
	x := 15
	a := &x // memory address
	fmt.Println(a)
	fmt.Println(*a)
	*a = 5
	fmt.Println(x)
	*a = *a * *a
	fmt.Println(x)
	fmt.Println(*a)
}
