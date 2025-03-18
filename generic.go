package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func genericSum[T constraints.Ordered](numbers []T) T {
	var result T
	for i := range numbers {
		result += numbers[i]
	}
	return result
}
func main() {
	numbers := []int{1, 2, 3, 4, 5}
	sum:=genericSum(numbers)
	fmt.Println(sum)
	
}
