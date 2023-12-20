// addition of first and second number

package main

import "fmt"

func main() {

	f := 0
	s := 1
	total := 0
	for i := 0; i < 8; i++ {
		total = f + s
		f = s
		s = total
		fmt.Println(total)
	}

}
