package main

import "fmt"

type details struct {
	name   string
	rollno string
	class  string
	age    int
	dob    uint16
}

func main() {
	st1 := details{name: prince, rollno: sst, class: final, age: 21, dob: 7}
	fmt.Println(st1.dob)
}
