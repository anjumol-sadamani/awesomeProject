package main

import (
	s "awesomeProject/service"
	"math/rand"
	"sync"
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
	wg := new(sync.WaitGroup)

	f := new(s.FoodGenerator)
	go f.GenerateFood(wg)

	for i := 0; i < 10; i++ {
		wg.Add(3)
		birds[rand.Intn(9)+1].DoSomething(wg)
		wg.Wait()
	}

}
