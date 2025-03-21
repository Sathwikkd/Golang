package main

import "fmt"

func channel(ch chan string) {
	ch <- "hi Go this is sathwik KD"
}

func main() {
	ch := make(chan string)
	 go channel(ch)

	message := <-ch
	fmt.Println(message)
}
