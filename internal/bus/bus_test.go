package bus

import (
	"reflect"
	"testing"
)

var testBus = Bus{
	Id:             "1",
	Capacity:       15,
	EmptySeat:      map[int]bool{},
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

func TestBus_EmptySeatInıt(t *testing.T) {
	err := testBus.EmptySeatInıt()
	checkError(t, err)
	got := testBus.EmptySeat
	want := map[int]bool{
		1:  false,
		2:  false,
		3:  false,
		4:  false,
		5:  false,
		6:  false,
		7:  false,
		8:  false,
		9:  false,
		10: false,
		11: false,
		12: false,
		13: false,
		14: false,
		15: false,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got = %v want = %v", got, want)
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
	got, err := testBus.SeatSearch(11)
	checkError(t, err)

	want := false
	if got != want {
		t.Errorf("got = %t want = %t", got, want)
	}
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
		got := testBus.SeatRegistration(2)
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
		for i := 1; i <= testBus.Capacity; i++ {
			testBus.EmptySeat[i] = true
		}
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
