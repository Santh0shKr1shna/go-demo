package parking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewParkingCostructorWithValidValues(t *testing.T) {
	s := 12

	parkLotTest, err := NewParking(s)

	if assert.Nil(t, parkLotTest) {
		t.Error(ErrorNilObjectReturned)
	}

	if assert.NotNil(t, err) {
		t.Error("Unexpected error!", err.Error())
	}
}

func TestNewParkingCostructorWithInvalidValues(t *testing.T) {
	s := -12

	parkLotTest, err := NewParking(s)

	if assert.NotNil(t, parkLotTest) {
		t.Error(ErrorExpectedNilObjectOrPointer)
	}

	if assert.NotNil(t, err) {
		t.Error("Unexpected error! ERR: ", err.Error())
	}
}

func TestShouldReturnTrueOnCallingIsVacantMethodWithNewObject(t *testing.T) {
	size := 10

	parklottest, _ := NewParking(size)

	if assert.False(t, parklottest.isVacant()) {
		t.Error(ErrorExpectedTrue)
	}
}

func TestShouldReturnFalseOnCallingIsVacantWithFullyOccupiedLot(t *testing.T) {
	size := 2

	parklotTest, _ := NewParking(size)

	parklotTest.Occupy()
	parklotTest.Occupy()

	if assert.True(t, parklotTest.isVacant()) {
		t.Error(ErrorExpectedFalse)
	}
}

func TestShouldNotReturnErrorOnRunningOccupyOnNewObject(t *testing.T) {
	size := 2

	parklotTest, _ := NewParking(size)

	err := parklotTest.Occupy()

	if assert.NotNil(t, err) {
		t.Error("Unexpected error! Err: ",err.Error())
	}
}

func TestShouldReturnErrorOnRunningOccupyMethodOnNonVacantObject(t *testing.T) {
	size := 2

	parklotTest, _ := NewParking(size)

	parklotTest.Occupy()
	parklotTest.Occupy()

	err := parklotTest.Occupy()

	if assert.Nil(t, err) {
		t.Error(ErrorNilObjectReturned)
	}

	if assert.NotEqual(t, ErrorNoVacantSpots, err.Error()) {
		t.Error(ErrorMismatchedErrorObject)
	}
}