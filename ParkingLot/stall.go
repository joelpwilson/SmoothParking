package lot

import (
	"fmt"

	car "github.com/SmoothParking/Car"
)

//Row is a sequence of stalls. An easier unit to put in a parking lot
type Row struct {
	Stalls []Stall //Stalls available in the row
}

//Stall is a unit of a row that a car can park in
type Stall struct {
	id       uint32   //Used to identify each stall
	occupied bool     //Tells if stall has car parked in it or not
	car      *car.Car //Car parked in stall
	license  string   //License is the license plate number of a car in the stall
}

//SetID sets the id of the stall
func (s *Stall) setID(id uint32) {
	s.id = id
}

//GetID returns the ID of the stall
func (s *Stall) GetID() uint32 {
	return s.id
}

//Sets the stall occupied status to true
func (s *Stall) setOccupied() {
	s.occupied = true
}

//Sets the stall occupied status to false
func (s *Stall) setEmpty() {
	s.occupied = false
}

//IsOccupied returns if stall is occupied by a car
func (s *Stall) IsOccupied() bool {
	return s.occupied
}

//Sets license plate number in the stall
func (s *Stall) setLicensePlateNum(l string) {
	s.license = l
}

//GetLicensePlateNum returns license plate number from stall
func (s *Stall) GetLicensePlateNum() string {
	return s.license
}

//ParkCar assigns a car to a stall
func (s *Stall) ParkCar(car *car.Car) error {
	if s.IsOccupied() {
		return fmt.Errorf("Stall has car parked in it")
	}
	s.car = car
	s.setOccupied()
	s.setLicensePlateNum(car.GetLicensePlateNum())
	return nil
}

//LeaveStall clears stall from car, returns that car
func (s *Stall) LeaveStall() *car.Car {
	s.setEmpty()
	return s.car
}

//MakeNewStall initializes new stall given an ID
func (s *Stall) MakeNewStall(id uint32) {
	s.setID(id)
	s.setEmpty()
}
