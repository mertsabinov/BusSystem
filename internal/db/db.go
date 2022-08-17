package db

import (
	"BusSystem/internal/bus"
)

type Data map[string]bus.Bus

func (d Data) Search(id string) (bus.Bus, error) {
	result, ok := d[id]
	if !ok {
		return bus.Bus{}, DbKeyNotFound
	}
	return result, nil
}

func (d Data) Save(vehicle bus.Bus) error {
	_, err := d.Search(vehicle.Id)

	switch err {
	case DbKeyNotFound:
		d[vehicle.Id] = vehicle
		break
	case nil:
		return DbAlReadyExisti
	}

	return nil
}

func (d Data) Update(vehicle bus.Bus) error {
	_, err := d.Search(vehicle.Id)
	switch err {
	case DbKeyNotFound:
		return DbKeyNotFound
	case nil:
		d[vehicle.Id] = vehicle
		break

	}
	return nil
}

func (d Data) Delete(vehicle bus.Bus, id string) error {
	return nil
}
