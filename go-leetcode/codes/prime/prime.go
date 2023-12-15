package main

import "fmt"

func main() {
	no := 4
	i := 2
	boolean_prime := true

	if no%i == 0 {
		fmt.Println("not prime number")
		boolean_prime = false
	}
	if boolean_prime {
		fmt.Println("prime number")
	}

}
