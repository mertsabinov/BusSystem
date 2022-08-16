package bus

type Bus struct {
	Id             string
	Capacity       int
	EmptySeat      []int
	DeparturePoint string
	Destination    string
}

func (b Bus) ReserveSeat(seatNumber int) error {
	return nil
}

func (b Bus) VehicleEmpty() bool {
	return true
}

func (b Bus) SeatRemove(seatNumber int) error {
	return nil
}
