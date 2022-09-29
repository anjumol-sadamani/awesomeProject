package service

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Bird struct {
	Height      int
	Name        string
	FleshWeight int
	ShitWeight  int
}

type FoodGenerator struct {
	Name     string
	Quantity int
}

func (b *Bird) Fly(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(5 * time.Second)
	fmt.Println(b.Name+" is flying and having fleshWeight", b.FleshWeight)
}

func (b *Bird) Eat(wg *sync.WaitGroup, f *FoodGenerator) {
	defer wg.Done()
	food := rand.Intn(10) + 1
	if f.Quantity >= food {
		f.Quantity = f.Quantity - 10
		fmt.Println(b.Name+" is eating ", food, " kilo food")
		fmt.Println("Remaining food in bin is ", f.Quantity)
		//time.Sleep(5 * time.Second)
		fleshPercentage := 0.5
		b.FleshWeight = b.FleshWeight + int(fleshPercentage*float64(food))
		b.ShitWeight = b.ShitWeight + int((1-fleshPercentage)*float64(food))

		if b.FleshWeight > 100 {
			b.FleshWeight = 100
			fmt.Println(b.Name + "'s stomach is full.. can't eat anymore!!")
		}
		fmt.Println(b.Name+" ate. Now fleshWeight is ", b.FleshWeight, " and shitWeight is ", b.ShitWeight)
	} else {
		fmt.Println("Only had ", f.Quantity, " kilo food. ", b.Name, " need ", food, " kilo food!!")
	}

}

func (b *Bird) Shit(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(5 * time.Second)
	shit := rand.Intn(10) + 1
	if b.ShitWeight-shit > 0 {
		b.ShitWeight = b.ShitWeight - shit
		fmt.Println(b.Name+" shat", shit, " kilo. Now shitWeight is ", b.ShitWeight)
	}
}

func (f *FoodGenerator) GenerateFood(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		fmt.Println("generating food and storing in bin")
		f.Quantity = f.Quantity + 10
		f.Name = "seeds"
		fmt.Println("food in bin is ", f.Quantity)
		time.Sleep(5 * time.Second)
	}
	//for f.Quantity+10 <= 100 {
	//	fmt.Println("generating food and storing in bin")
	//	f.Quantity = f.Quantity + 10
	//	f.Name = "seeds"
	//	fmt.Println("food in bin is ", f.Quantity)
	//	time.Sleep(5 * time.Second)
	//}

}

func (b *Bird) DoSomething(wg *sync.WaitGroup, f *FoodGenerator) {

	go b.Fly(wg)
	go b.Eat(wg, f)
	go b.Shit(wg)

}
