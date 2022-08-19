package bus

type busErr string

const (
	BusErrSeatFull          busErr = "The entered seat number is full"
	BusSeatNumberFalse      busErr = "This seat number is false"
	BusCheckSeatNumberSmall busErr = "this seat number must be biger than 0"
	BusCheckSeatNumberBig   busErr = "this seat number must be smaller than capacity"
)

func (e busErr) Error() string {
	return string(e)
}
