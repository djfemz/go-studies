package main

import (
	"fmt"
)

func main() {
	x := make([]int, 10, 20)
	for i := 1; i <= 100; i++ {
		if 100 % i == 0 {
			x = append(x, i)
			fmt.Print(x)
		}

	}
}
