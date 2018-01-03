package main

import "fmt"

type car struct {
	speed    float64
	distance float64
	time     float64
}

func (c car) tim() float64 {
	return float64(c.speed) * (c.distance)
}
func main() {
	det1 := car{
		speed:    100,
		distance: 450,
		time:     60}

	fmt.Println(det1.time)
	fmt.Println(det1.tim())

}
