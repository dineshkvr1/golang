package main

import "fmt"

//find no of digits in given number

func main() {

	no := 12345
	count := 0
	for no > 0 {
		no = no / 10
		fmt.Println(no)
		count++
		//fmt.Println(count)
	}
	fmt.Println(count)

}
