package car

//Car is the object that will be used in traffic situations for filling stalls in a parking lot
type Car struct {
	license string //License plate number for the car
}

//SetLicensePlateNum sets the license plate number for a car
func (c *Car) setLicensePlateNum(num string) {
	c.license = num
}

//GetLicensePlateNum returns the license plate number for a car
func (c *Car) GetLicensePlateNum() string {
	return c.license
}

//MakeNewCar creates a new car object and assigns the license number to it
func MakeNewCar(l string) *Car {
	car := new(Car)
	car.setLicensePlateNum(l)
	return car
}
