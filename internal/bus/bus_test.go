package bus

import (
	"reflect"
	"testing"
)

var testBus = Bus{
	Id:             "1",
	Capacity:       30,
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
