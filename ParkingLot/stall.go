package lot

import (
	"github.com/SmoothParking/car"
)

type Stall {
	id		 uint32 //Used to identify each stall
	occupied bool   //Tells if stall has car parked in it or not
	car		 *Car    //Car parked in stall
}

//Sets the id of the stall 
func (s *Stall) SetID(id uint32) {
	s.id = id
}

//Returns the ID of the stall
func (s *Stall) GetID() id uint32 {
	return s.id
}

//Sets the stall occupied status to true
func (s *Stall) setOccupied() {
	s.occupied = true
}

//Sets the stall occupied status to false
func (s *Stall) setEmpty(){
	s.occupied = false
}

//Returns if stall is occupied by a car
func (s *Stall) IsOccupied() bool {
	return s.occupied
}

//Sets license plate number in the stall
func (s *Stall) setLicensePlateNum(l string) {
	s.license = l
}

//Get license plate number from stall
func (s *stall) GetLicensePlateNum() string{
	return s.license
}

//Assigns a car to a stall
func (s *Stall) ParkCar(car Car){
	s.car = car 
	s.setOccupied(true)
}

//Clears stall from car, returns that car
func (s *Stall) LeaveStall() *Car {
	s.setOccupied(false)
	return s.Car
}