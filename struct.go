package main

import "fmt"

type Student struct {
	name string
	age  int
}

func main() {
	stu1 := Student{}
	stu2 := Student{}
	var stu3 Student

	stu1.name = "sath"
	stu2.name = "kd"

	stu2.age = 1
	stu1.age = 2

	stu3.age = 5

	fmt.Println(stu1.name, stu3.age)

	structFunc(stu1)
	structFunc(stu2)

}


func structFunc(strvar Student){
	fmt.Println(strvar.name)
	fmt.Println(strvar.age)
}
