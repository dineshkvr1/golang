package main

import "fmt"

func isPalindrome(x int) bool {
	res := 0
	y := x
	//fmt.Println(y)
	for y > 0 {
		sum := y % 10
		//fmt.Println(sum)
		res = res*10 + sum
		//fmt.Println(res)
		y = y / 10
		fmt.Println(y)
	}
	if res == x {
		return true
	} else {
		return false
	}
}

func main() {

	two := isPalindrome(10)
	fmt.Println(two)

}
