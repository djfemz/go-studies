package main

import "fmt"

func main() {
	total := 0
	for v := 2; v < 1000; v++ {
		if v%3 == 0 || v%5 == 0 {
			total += v
		}
	}
	fmt.Println(total)
}
