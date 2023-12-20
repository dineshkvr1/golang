// factorial of given number
package main

import "fmt"

func main() {

	no := 5
	fact := no

	for no > 1 {
		no--
		fact = fact * no
	}
	fmt.Println(fact)

}
