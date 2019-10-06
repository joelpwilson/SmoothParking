package lot

//ParkingLot is the object where cars can find a parking spot to park in
type ParkingLot struct {
	name string         //Name given to the parking lot
	rows map[string]Row //Rows available in this parking lot. Accessed by their name. Ex. "Left", "Street", "NorthWall"
}

//CreateParkingLot initializes and returns a parking lot, assigns the given name
func CreateParkingLot(name string) *ParkingLot {
	pL := new(ParkingLot)
	pL.SetLotName(name)
	return pL
}

//SetLotName sets the parking lot name to a given name
func (p *ParkingLot) SetLotName(name string) {
	p.name = name
}

//AddRow adds a row to the parking lot with the given name, id
func (p *ParkingLot) AddRow(name string, numStalls uint32) {
	row := new(Row)
	var i uint32
	for i = 0; i < numStalls; i++ {
		stall := new(Stall)
		stall.MakeNewStall(i)
		row.stalls = append(row.stalls, *stall)
	}
	p.rows[name] = *row
}
