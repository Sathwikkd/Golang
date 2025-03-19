package main

import "fmt"

func main() {
	a := map[string]string{"name": "sath", "village": "kadasur", "post": "kadasur"}
	fmt.Println(a)

	a["post"] = "humcha"
	fmt.Println(a)

	delete(a, "name")
	fmt.Println(a)
}
