package lot

type Row {
	stalls []Stall //Stalls available in the row
}

type ParkingLot {
	name 	string 		   //Name given to the parking lot
	rows	map[string]Row //Rows available in this parking lot. Accessed by their name. Ex. "Left", "Street", "NorthWall"
}

//Initializes and returns a parking lot, assigns the given name
func CreateParkingLot(name string) *ParkingLot {
	pL := new(ParkingLot)
	pl.SetLotName(name)
	pl.rows = make(map[string]Row)
}

//Sets the parking lot name to a given name
func (p *ParkingLot) SetLotName(name string) {
	p.name = name
}

//Adds a row to the parking lot with the given name, id
func (p	*ParkingLot) AddRow(name string, numStalls uint32) {
	row := new(Row)
	for i := 0; i < numStalls; i++ {

		row.stalls.append()
	}
}