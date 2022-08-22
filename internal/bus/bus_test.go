package bus

import (
	"reflect"
	"testing"
)

var testBus = Bus{
	Id:             "1",
	Capacity:       30,
	EmptySeat:      map[int]bool{2: false, 10: false, 11: true},
	DeparturePoint: "test",
	Destination:    "test test",
}

func checkDeepEquel(t *testing.T, got Bus, want Bus) {
	t.Helper()
	if reflect.DeepEqual(got, want) {
		t.Errorf("got = %v, want = %v", got, want)
	}
}
func checkError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func check(t *testing.T, got int, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got = %d want = %d", got, want)
	}
}

func TestBus_checkSeatNumber(t *testing.T) {
	t.Run("Check seat number", func(t *testing.T) {
		err := testBus.CheckSeatNumber(15)
		checkError(t, err)
	})

	t.Run("Check seat number big", func(t *testing.T) {
		got := testBus.CheckSeatNumber(32)
		want := BusCheckSeatNumberBig
		if got != want {
			t.Errorf("got = %s want = %s", got, want)
		}
	})

	t.Run("Check seat number small", func(t *testing.T) {
		got := testBus.CheckSeatNumber(0)
		want := BusCheckSeatNumberSmall
		if got != want {
			t.Errorf("got = %s want = %s", got, want)
		}
	})
}

func TestBus_SeatSearch(t *testing.T) {

	t.Run("Seat Search", func(t *testing.T) {
		got, err := testBus.SeatSearch(11)
		checkError(t, err)

		want := true
		if got != want {
			t.Errorf("got = %t want = %t", got, want)
		}
	})

	t.Run("Seat search error", func(t *testing.T) {
		_, got := testBus.SeatSearch(1)

		want := BusSeatNumberFalse
		if got != want {
			t.Errorf("got = %s want = %s", got, want)
		}
	})
}

func TestBus_SeatRegistration(t *testing.T) {
	t.Run("Seat Registarion", func(t *testing.T) {
		err := testBus.SeatRegistration(2)
		checkError(t, err)

		got, err := testBus.SeatSearch(2)
		checkError(t, err)
		want := true

		if got != want {
			t.Errorf("got = %t want = %t", got, want)
		}
	})

	t.Run("Seat number full", func(t *testing.T) {
		got := testBus.SeatRegistration(11)
		want := BusSeatNumberFull

		if got != want {
			t.Errorf("got = %s want = %s", got, want)
		}
	})

}

func TestBus_VehicleEmpty(t *testing.T) {
	t.Run("Vehicle empty", func(t *testing.T) {
		got, err := testBus.VehicleEmpty()
		checkError(t, err)
		want := true
		if got != want {
			t.Errorf("got = %t want = %t", got, want)
		}
	})

	t.Run("Error : All seats are full", func(t *testing.T) {
		err2 := testBus.SeatRegistration(10)
		checkError(t, err2)

		_, got := testBus.VehicleEmpty()
		want := BusErrSAlleatFull
		if got != want {
			t.Errorf("got = %s want = %s", got, want)
		}
	})
}

func TestBus_SeatRemove(t *testing.T) {
	t.Run("Seat number remove", func(t *testing.T) {
		err := testBus.SeatRemove(2)
		checkError(t, err)

		got, err := testBus.SeatSearch(2)
		checkError(t, err)
		want := false
		if got != want {
			t.Errorf("got = %t want = %t", got, want)
		}
	})

	t.Run("Error: seat number is empty", func(t *testing.T) {
		got := testBus.SeatRemove(2)
		want := BusErrSeatNumberEmpty

		if got != want {
			t.Errorf("got = %s want = %s", got, want)
		}
	})
}
