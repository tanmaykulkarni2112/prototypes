package main

import (
	"fmt"

	"github.com/tanmaykulkarni2112/prototypes/factoryPattern/carfactory"
)

func main() {

	car := carfactory.GetCar("sedan")

	if car == nil {
		fmt.Println("Invalid car type")
		return
	}

	fmt.Println(car.GetName())
}