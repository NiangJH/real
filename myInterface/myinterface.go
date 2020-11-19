package myInterface

import "fmt"

type Animal interface {
	Eat(food string)
	Drink(ml int)
}

type Sheep struct {
	water int
	food string
}

func (s Sheep)Eat(food string) {
	fmt.Println("eat:" + food)
}

func (s Sheep)Drink(ml int) {
	fmt.Println(ml)
}


func createSheep() {
	sheep := new(Sheep)
	sheep.water = 2
	sheep.food = "grass"

	animal := []Animal{sheep}
	//animal[0].Eat(anm)
	for _, v := range animal {
		v.Eat("xx")
		v.Drink(4)
	}
}