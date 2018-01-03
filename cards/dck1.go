package main

import "fmt"

type deck []string

func newDeck() deck {

	card := deck{}

	cardSuits := []string{"spades", "hearts", "diamonds"}
	cardValues := []string{"ace", "two", "three"}
	for _, suit := range cardSuits {
		for _, values := range cardValues {
			card = append(card, suit+"of"+values)
		}
	}
	return card

}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
