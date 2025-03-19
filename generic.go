package main

import "fmt"

func PrintSlice[T any](s []T) {
	for _,i := range s {
		fmt.Println(i)
	}
}

func main() {
	intSlice := []int{1, 2, 3}
	stringSlice := []string{"hello", "world"}

	PrintSlice[int](intSlice)
	PrintSlice[string](stringSlice)
}
