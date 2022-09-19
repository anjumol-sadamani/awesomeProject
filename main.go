package main

import (
	s "awesomeProject/service"
	"math/rand"
)

func main() {
	var birds = make([]s.Bird, 10)
	n := 'a'
	for i := 0; i < len(birds); i++ {
		birds[i].Height = rand.Intn(10) + 1
		birds[i].Name = string(n)
		birds[i].ShitWeight = 0
		birds[i].FleshWeight = rand.Intn(10) + 1
		n++
	}

	for i := 0; i < 100; i++ {
		switch rand.Intn(3) {
		case 0:
			birds[rand.Intn(9)+1].Eat()
		case 1:
			birds[rand.Intn(9)+1].Fly()
		case 2:
			birds[rand.Intn(9)+1].Shit()
		}
	}

}
