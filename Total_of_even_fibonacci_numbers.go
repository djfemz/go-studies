package main

import (
	"fmt"
	"strconv"
)

func fib(n int) int {
	if n <= 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
func sum(x ...int) int {
	total := 0
	for _, v := range x {
		if v%2 == 0 {
			total += v
		}
	}
	return total
}

func main() {
	fib(2)
	fmt.Println(fib(2))
	for i := 0; i < 33; i++ {
		fmt.Print(strconv.Itoa(fib(i)) + " ")
	}
	xi := []int{1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368, 75025, 121393, 196418, 317811, 514229, 832040, 1346269, 2178309, 3524578}
	a := sum(xi...)
	fmt.Println("\n The total of even fibonacci terms less the 4 million, is", a)
}
