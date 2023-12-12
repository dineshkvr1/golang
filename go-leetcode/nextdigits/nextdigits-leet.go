package main

import "fmt"

func plusOne(digits []int) []int {

	n := len(digits)
	//fmt.Println(n)
	for i := n - 1; i >= 0; i-- {

		//fmt.Println(i)
		//fmt.Println(digits[i])
		if digits[i] < 9 {
			digits[i]++
			//fmt.Println(digits[i])
			return digits

		}
		digits[i] = 0
		fmt.Println(digits[i])

	}
	digits = append([]int{1}, digits...)
	fmt.Println(digits)
	return digits
}

func main() {

	one := plusOne([]int{1, 2, 3, 4, 9})
	fmt.Println(one)

}
