package main

import "fmt"

func main(){
	var fib1 int = 0 
	var fib2 int = 1
	var soma int
	var i int

	for i = 0; i < 6; i++{
		soma = fib1 + fib2
		fib1 = fib2
		fib2 = soma

		fmt.Printf(" %d ", soma)
	}
	fmt.Println()
}