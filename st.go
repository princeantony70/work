package main

import "fmt"

type next struct {
	one   int
	two   int
	three int
	four  int
}

func main() {
	maths := next{one: 1, two: 2, three: 3, four: 4}

	fmt.Println(maths)

}
