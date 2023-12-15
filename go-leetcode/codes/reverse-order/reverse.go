package main

import "fmt"

func main() {
	no := 1234
	rem := 0
	for no > 0 {
		rem = no % 10
		no = no / 10
		//fmt.Println(no)
		fmt.Println(rem)
	}

}

//reverse number
