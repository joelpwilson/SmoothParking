package lot

//ParkingLot is the object where cars can find a parking spot to park in
type ParkingLot struct {
	Name string         //Name given to the parking lot
	Rows map[string]Row //Rows available in this parking lot. Accessed by their name. Ex. "Left", "Street", "NorthWall"
}

//CreateParkingLot initializes and returns a parking lot, assigns the given name
func CreateParkingLot(name string) *ParkingLot {
	pL := new(ParkingLot)
	pL.SetLotName(name)
	pL.Rows = make(map[string]Row)
	return pL
}

//SetLotName sets the parking lot name to a given name
func (p *ParkingLot) SetLotName(name string) {
	p.Name = name
}

//AddRow adds a row to the parking lot with the given name, id
func (p *ParkingLot) AddRow(name string, numStalls uint32) {
	row := new(Row)
	var i uint32
	for i = 0; i < numStalls; i++ {
		stall := new(Stall)
		stall.MakeNewStall(i)
		row.Stalls = append(row.Stalls, *stall)
	}
	p.Rows[name] = *row
}

//LotPrinter is an object that prints out the parking lot to the console for visualization
type LotPrinter struct {
	pl *ParkingLot
}

//NewLotPrinter initializes and returns a new LotPri
func NewLotPrinter(p *ParkingLot) *LotPrinter {
	l := new(LotPrinter)
	l.pl = p
	return l
}
