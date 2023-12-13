package main

import "fmt"

func main() {
	n := 1
	y := 0
	for n <= 5 {
		y = (n * (n + 1))
		fmt.Println(y)
		n++
	}

}

//output 2 6 12 20 30
