package main

import "fmt"

func main() {
	no1 := 3
	no2 := 8

	var max int
	if no1 > no2 {
		max = no1
	} else {
		max = no2
	}
	fmt.Println("maximum:", max)
	//For is Go's "while"
	//lcm := 0
	for max > 0 {
		fmt.Println(max)
		if (max%no1 == 0) && (max%no2 == 0) {
			lcm := max
			fmt.Println(lcm)
			/* terminate the loop using break statement */
			break
		}
		max++
	}
	//fmt.Println(lcm)

	//fmt.Println(lcm)

}
