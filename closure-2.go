package main

import "fmt"

func createCounter() func() int {
	count := 0
	increment := func() int {
		count++
		return count
	}
	return increment
}

func main() {
	counter0 := createCounter()
	counter1 := createCounter()
	fmt.Println(counter0())
	fmt.Println(counter0())
	fmt.Println(counter1())
	fmt.Println(counter1())
	fmt.Println(counter1())
	fmt.Println(counter1())

}
