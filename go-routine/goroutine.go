package main

import (
	"fmt"
	"time"
)

func printName() {
	fmt.Println("hello this is sathwik from vithsutra")
	time.Sleep(500 * time.Millisecond)
	for i := 'A'; i < 'z'; i++ {
		fmt.Println(string(i))
	}
}

func printUsn() {
	fmt.Println("my usn is 4AL21IS047")
	time.Sleep(1 * time.Millisecond)

}

func main() {
	go printUsn()
	go printName()
	time.Sleep(3 * time.Second)
	//fmt.Print("hi")
}
