package main

import (
	"math/rand"
	"time"

	car "github.com/SmoothParking/Car"
	lot "github.com/SmoothParking/ParkingLot"
	parking "github.com/SmoothParking/Visualize"
)

func main() {
	numStalls := 100
	pl := lot.CreateParkingLot("NewLot")
	pl.AddRow("NorthWall", 100)

	// cars := list.New()
	// c := car.MakeNewCar("a")
	// cars.PushBack(c)

	lp := parking.NewLotPrinter(pl)
	lp.PrintRow("NorthWall")

	//Sets the seed to be different each run time
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < numStalls/2; i++ {
		randStall := rand.Intn(numStalls)
		c := car.MakeNewCar("a")
		if !pl.Rows["NorthWall"].Stalls[randStall].IsOccupied() {
			pl.Rows["NorthWall"].Stalls[randStall].ParkCar(c)
		}
	}
	lp.PrintRow("NorthWall")
}
