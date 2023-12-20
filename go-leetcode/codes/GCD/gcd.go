package main

import "fmt"

func main() {
	no1 := 12 //common divisor - 12 6 3 2
	no2 := 30 // common diversor-30 15 10 6 5 3 2 -- 6 is greater common divisor

	var min int
	if no1 < no2 {
		min = no1
	} else {
		min = no2
	}
	fmt.Println("minimum:", min)
	//For is Go's "while"
	//lcm := 0
	for min > 2 {
		//fmt.Println(min)
		if (no1%min == 0) && (no2%min == 0) {
			gcd := min
			fmt.Println(gcd)
			/* terminate the loop using break statement */
			break
		}
		min--
	}
	//fmt.Println(lcm)

	//fmt.Println(lcm)

}
