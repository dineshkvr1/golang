package main

import "fmt"

func main() {
	n := 1
	for n <= 20 {
		if (n%3 == 0) || (n%5 == 0) {
			fmt.Println(n)

		}
		n++
	}
}
