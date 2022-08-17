package db

import (
	"BusSystem/internal/bus"
	"reflect"
	"testing"
)

func TestData_Save(t *testing.T) {
	d := Data{}
	b := bus.Bus{
		Id: "1",
		Capacity: 30,
		EmptySeat: []int,
		DeparturePoint: "test",
		Destination: "test test",
	}

	d.Save(b,"1")

	got,_ := d.Search(b.Id)
	want := b

	if !reflect.DeepEqual(got,b){
		t.Errorf("got = %q want = %q",got,want)
	}
}
