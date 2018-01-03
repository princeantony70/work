package main

import "fmt"

type deck []string

func food() {

	combination := deck{}

	mdish := []string{"briyani", "meals", "tandoori"}
	sdish := []string{"chillychicke", "kootu", "lemon"}

	for _, m := range mdish {
		for _, s := range sdish {
			combination = append(card, m+"with"+s)
		}
	}

	return combination
}

func (d deck) print() {
	for i, combination = range d {
		fmt.Println(i, combination)

	}
}
