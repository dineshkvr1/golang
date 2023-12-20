// square root of a given number

package main

import "fmt"

func main() {
	no := 256
	div := no / 2
	i := 1
	for i < div {
		m := no / i
		if m == i {
			fmt.Println(m)
			break
		}
		i++

	}

}
