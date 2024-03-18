package parking

import "errors"

type Car struct {
	numPlate string
	owner string
}

func NewCar(num, owner string) (*Car, error) {
	if num == "" || owner == "" {
		return nil, errors.New(ErrorEmptyStrings)
	}

	return &Car{
		numPlate: num,
		owner: owner,
	}, nil
}