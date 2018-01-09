package main

func main() {

	type A struct{ name string }
	type B A

	type (
		C string
		D map[string]int
	)
}
