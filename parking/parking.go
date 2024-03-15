package parking

import "errors"

type ParkingLot struct {
	capacity int
	occupied int
	carsParked []Car
}

func NewParking(size int) (*ParkingLot, error) {
	if size > 0 {
		lots := ParkingLot {
			capacity: size,
			occupied: 0,
			carsParked: make([]Car, 0),
		}
		return &lots, nil
	}

	return nil, errors.New(ErrorInvalidLotSize)
}

func (p ParkingLot) isVacant() bool {
	return p.occupied < p.capacity
}

func (p *ParkingLot) Park(car Car) error {
	if p.isVacant() {
		p.carsParked = append(p.carsParked, car)
		p.occupied += 1
	} else {
		return errors.New(ErrorNoVacantSpots)
	}

	return nil
}

func (p *ParkingLot) UnPark(car Car) error {
	if p.occupied == 0 {
		return errors.New(ErrorNoCarsAreParked)
	}
	
}