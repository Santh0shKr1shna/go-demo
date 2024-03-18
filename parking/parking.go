package parking

import (
	"errors"
)

type ParkingLot struct {
	capacity int
	occupied int
	carsParked []Car
	observers []Observer
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

func (p ParkingLot) Contains(car Car) bool {
	for _, c := range p.carsParked {
		if c.numPlate == car.numPlate {
			return true
		}
	}
	return false
}

func (p *ParkingLot) Park(car Car) error {
	if !p.isVacant() {
		return errors.New(ErrorNoVacantSpots)
		
	} else if p.Contains(car){
		return errors.New(ErrorCarAlreadyInLot)
	} else {
		p.carsParked = append(p.carsParked, car)
		p.occupied += 1
	}

	if p.occupied == p.capacity {
		p.notify()
	}

	return nil
}

func (p *ParkingLot) UnPark(car Car) error {
	if p.occupied == 0 {
		return errors.New(ErrorNoCarsAreParked)
	} else if !p.Contains(car) {
		return errors.New(ErrorThisCarIsNotParkedHere)
	} else {
		for i, c := range p.carsParked {
			if c.numPlate == car.numPlate {
				p.carsParked = append(p.carsParked[:i], p.carsParked[i+1:]...)
				p.occupied -= 1
				break
			}
		}
	}

	return nil
}

func (p *ParkingLot) register(o Observer) {
	p.observers = append(p.observers, o)
}

func (p *ParkingLot) notify() {
	for _, obs := range p.observers {
		obs.Update()
	}
}