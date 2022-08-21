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
	err := b.CheckSeatNumber(seatNumber)
	if err != nil {
		return true, err
	}

	result, ok := b.EmptySeat[seatNumber]
	if !ok {
		return true, BusSeatNumberFalse
	}
	return result, nil
}

func (b Bus) SeatRegistration(seatNumber int) error {

	result, err := b.SeatSearch(seatNumber)
	if err != nil {
		return err
	}
	if result {
		return BusSeatNumberFull
	}

	b.EmptySeat[seatNumber] = true

	return nil
}

func (b Bus) VehicleEmpty() (bool, error) {
	for _, value := range b.EmptySeat {
		if value == false {
			return true, nil
		}
	}
	return false, BusErrSAlleatFull
}

func (b Bus) SeatRemove(seatNumber int) error {
	return nil
}
