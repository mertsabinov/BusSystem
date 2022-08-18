package db

import (
	bus "BusSystem/internal/bus"
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
		err := d.Save(b)
		chackError(t, err)

		result, err := d.Search(b.Id)
		chackError(t, err)
		got := result
		want := b
		chackDeepEqual(t, got, want)
	})

	t.Run("Db key AlReady Existi", func(t *testing.T) {
		got := d.Save(b)
		want := DbAlReadyExisti
		if got != want {
			t.Errorf("got = %s want = %s", got, want)
		}
	})
}

func TestData_Update(t *testing.T) {
	d := Data{
		"1": {
			Id:             "1",
			Capacity:       30,
			EmptySeat:      []int{},
			DeparturePoint: "test",
			Destination:    "test test",
		},
	}
	newBus := bus.Bus{
		Id:             "1",
		Capacity:       15,
		EmptySeat:      []int{},
		DeparturePoint: "test",
		Destination:    "test test",
	}

	err := d.Update(newBus)
	chackError(t, err)

	result, err := d.Search(newBus.Id)
	chackError(t, err)
	got := result
	want := newBus
	chackDeepEqual(t, got, want)
}

func TestData_Delete(t *testing.T) {
	d := Data{
		"1": {
			Id:             "1",
			Capacity:       30,
			EmptySeat:      []int{},
			DeparturePoint: "test",
			Destination:    "test test",
		},
	}
	err := d.Delete("1")
	chackError(t, err)

	_, got := d.Search("1")
	want := DbKeyNotFound
	if got != want {
		t.Errorf("got = %s want = %s", got, want)
	}
}

func BenchmarkData_Search(b *testing.B) {
	d := Data{
		"1": {
			Id:             "1",
			Capacity:       30,
			EmptySeat:      []int{},
			DeparturePoint: "test",
			Destination:    "test test",
		},
	}
	for i := 0; i < b.N; i++ {
		d.Search("1")
	}
}

func BenchmarkData_Save(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d := Data{
			"1": {
				Id:             "1",
				Capacity:       30,
				EmptySeat:      []int{},
				DeparturePoint: "test",
				Destination:    "test test",
			},
		}
		newBus := bus.Bus{
			Id:             "2",
			Capacity:       15,
			EmptySeat:      []int{},
			DeparturePoint: "test",
			Destination:    "test test",
		}
		d.Save(newBus)
	}
}

func BenchmarkData_Update(b *testing.B) {
	d := Data{
		"1": {
			Id:             "1",
			Capacity:       30,
			EmptySeat:      []int{},
			DeparturePoint: "test",
			Destination:    "test test",
		},
	}
	newBus := bus.Bus{
		Id:             "1",
		Capacity:       15,
		EmptySeat:      []int{},
		DeparturePoint: "test",
		Destination:    "test test",
	}
	for i := 0; i < b.N; i++ {
		d.Update(newBus)
	}
}

func BenchmarkData_Delete(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d := Data{
			"1": {
				Id:             "1",
				Capacity:       30,
				EmptySeat:      []int{},
				DeparturePoint: "test",
				Destination:    "test test",
			},
		}
		d.Delete("1")
	}
}
