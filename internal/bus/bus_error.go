package bus

type BusErr string

const (
	BusErrSeatNumberEmpty   BusErr = "This seat number is empty"
	BusErrSAlleatFull       BusErr = "All seats are full"
	BusSeatNumberFalse      BusErr = "This seat number is false"
	BusCheckSeatNumberSmall BusErr = "this seat number must be biger than 0"
	BusCheckSeatNumberBig   BusErr = "this seat number must be smaller than capacity"
	BusSeatNumberFull       BusErr = "This seat number is full"
)

func (e BusErr) Error() string {
	return string(e)
}
