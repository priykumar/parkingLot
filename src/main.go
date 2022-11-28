package main

import (
	"fmt"

	controller "github.com/priykumar/parkingLot/src/controller"
)

func main() {
	fmt.Println("Designing parking lot")

	c := controller.InitController()
	c.CreateParkingLot()
}
