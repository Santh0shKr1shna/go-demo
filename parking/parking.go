package parking

import "errors"

type ParkingLot struct {
	capacity int
	occupied int
}

func NewParking(size int) (*ParkingLot, error) {
	if size > 0 {
		lots := ParkingLot {
			capacity: size,
			occupied: 0,
		}
		return &lots, nil
	}

	return nil, errors.New(ErrorInvalidLotSize)
}

func (p ParkingLot) isVacant() bool {
	return p.occupied < p.capacity
}

func (p *ParkingLot) Occupy() error {
	if p.isVacant() {
		p.occupied += 1
	} else {
		return errors.New(ErrorNoVacantSpots)
	}

	return nil
}