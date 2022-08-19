package bus

type Bus struct {
	Id             string
	Capacity       int
	EmptySeat      map[int]bool
	DeparturePoint string
	Destination    string
}

func (b Bus) CheckSeatNumber(seatNumber int) error {
	if seatNumber < 1 {
		return BusCheckSeatNumberSmall
	}
	if seatNumber > b.Capacity {
		return BusCheckSeatNumberBig
	}
	return nil
}

func (b Bus) SeatSearch(seatNumber int) (bool, error) {
	result, ok := b.EmptySeat[seatNumber]
	if !ok {
		return true, BusSeatNumberFalse
	}
	return result, nil
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
