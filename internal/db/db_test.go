package db

import (
	"BusSystem/internal/bus"
	"reflect"
	"testing"
)

func chackError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func chackDeepEqual(t *testing.T, got bus.Bus, want bus.Bus) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got = %v want = %v", got, want)
	}
}

func TestData_Search(t *testing.T) {
	d := Data{
		"1": {
			Id:             "1",
			Capacity:       30,
			EmptySeat:      []int{},
			DeparturePoint: "test",
			Destination:    "test test",
		},
	}

	b := bus.Bus{
		Id:             "1",
		Capacity:       30,
		EmptySeat:      []int{},
		DeparturePoint: "test",
		Destination:    "test test",
	}

	result, err := d.Search("1")
	chackError(t, err)
	got := result
	want := b
	chackDeepEqual(t, got, want)
}

func TestData_Save(t *testing.T) {
	d := Data{}
	b := bus.Bus{
		Id:             "1",
		Capacity:       30,
		EmptySeat:      []int{},
		DeparturePoint: "test",
		Destination:    "test test",
	}

	t.Run("Db Save", func(t *testing.T) {
		err := d.Save(b, b.Id)
		chackError(t, err)

		result, err := d.Search(b.Id)
		chackError(t, err)
		got := result
		want := b
		chackDeepEqual(t, got, want)
	})

	t.Run("Db key AlReady Existi", func(t *testing.T) {
		got := d.Save(b, b.Id)
		want := DbAlReadyExisti
		if got != want {
			t.Errorf("got = %s want = %s", got, want)
		}
	})

}
