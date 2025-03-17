package main

import "fmt"

func DebitAction() func(int) int {
	Amount := 100
	DebitFunc := func(debitAmount int) int {
		Amount -= debitAmount
		return Amount
	}
	return DebitFunc
}

func main() {
	shashi := DebitAction()
	sath := DebitAction()
	fmt.Println(shashi(10))
	fmt.Println(shashi(25))
	fmt.Println(sath(10))
}
