package parking

import (
	"fmt"
	"strconv"

	lot "github.com/SmoothParking/ParkingLot"
)

//TODO: incorporate UI into this for more phone-app-like feel

//LotPrinter is an object that prints out the parking lot to the console for visualization
type LotPrinter struct {
	pl *lot.ParkingLot
}

//NewLotPrinter initializes and returns a new LotPri
func NewLotPrinter(p *lot.ParkingLot) *LotPrinter {
	l := new(LotPrinter)
	l.pl = p
	return l
}

//PrintRow prints out a row showing empty or full stalls. Id is printed above the stall
//Example Below:
//
//--NorthWall--
//------------------------------------------------------
//    _0_  _1_  _2_  _3_  _4_
//  ||   ||   ||   ||   ||   ||
//  || X ||   || X ||   ||   ||
//  ||   ||   ||   ||   ||   ||
//
func (l *LotPrinter) PrintRow(name string) {
	stalls := l.pl.Rows[name].Stalls
	consoleWidth := 40
	var numIter int
	var remainder int
	if len(stalls) < consoleWidth {
		numIter = len(stalls)
	} else {
		remainder = len(stalls) % consoleWidth
		numIter = consoleWidth
	}
	for j := 0; j < (len(stalls)/consoleWidth)+1; j++ {
		if j == (len(stalls) / consoleWidth) {
			numIter = remainder
		}
		if numIter != 0 {
			fmt.Println("")
			fmt.Print("--")
			for i := j * consoleWidth; i < numIter+j*consoleWidth; i++ {
				fmt.Print("-----")
			}
			fmt.Println("")

			fmt.Println("--| " + name + " |--")
			fmt.Print("--")
			for i := j * consoleWidth; i < numIter+j*consoleWidth; i++ {
				fmt.Print("-----")
			}
			fmt.Println("")
			fmt.Print("    ")
			for i := j * consoleWidth; i < numIter+j*consoleWidth; i++ {
				if i < 9 {
					fmt.Print("_" + strconv.Itoa(int(stalls[i].GetID())+1) + "_  ")
				} else if i < 99 {
					fmt.Print("_" + strconv.Itoa(int(stalls[i].GetID())+1) + "  ")
				} else {
					fmt.Print(strconv.Itoa(int(stalls[i].GetID()+1)) + "  ")
				}
			}

			fmt.Println("")
			fmt.Print("  ||")
			for i := j * consoleWidth; i < numIter+j*consoleWidth; i++ {
				fmt.Print("   ||")
			}

			fmt.Println("")
			fmt.Print("  ||")
			for i := j * consoleWidth; i < numIter+j*consoleWidth; i++ {
				fmt.Print(" ")
				if stalls[i].IsOccupied() {
					fmt.Print("X")
				} else {
					fmt.Print(" ")
				}
				fmt.Print(" ||")
			}

			fmt.Println("")
			fmt.Print("  ||")
			for i := j * consoleWidth; i < numIter+j*consoleWidth; i++ {
				fmt.Print("   ||")
			}
			fmt.Println("")
		}
	}
}
