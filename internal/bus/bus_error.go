package bus

type BusError string

const (
	BusErrSeatNumberEmpty   BusError = "This seat number is empty"
	BusErrSAlleatFull       BusError = "All seats are full"
	BusSeatNumberFalse      BusError = "This seat number is false"
	BusCheckSeatNumberSmall BusError = "this seat number must be biger than 0"
	BusCheckSeatNumberBig   BusError = "this seat number must be smaller than capacity"
	BusSeatNumberFull       BusError = "This seat number is full"
)

func (e BusError) Error() string {
	return string(e)
}
