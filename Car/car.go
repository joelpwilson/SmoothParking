package car

//Car is the object that will be used in traffic situations for filling stalls in a parking lot
type Car struct {
	license string //License plate number for the car
}

//SetLicensePlateNum sets the license plate number for a car
func (c *Car) SetLicensePlateNum(num string) {
	c.license = num
}

//GetLicensePlateNum returns the license plate number for a car
func (c *Car) GetLicensePlateNum() string {
	return c.license
}
