package main

import "fmt"

func main() {
	i := 1
	j := 10
	y := 0
	for i <= 5 {
		y = i * j
		fmt.Println(y)
		i++
		fmt.Println("i value:", i)
		j--
		fmt.Println("j value:", j)
	}

}
