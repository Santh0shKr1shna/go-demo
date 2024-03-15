package parking

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewParkingCostructorWithValidValues(t *testing.T) {
	s := 12

	parkLotTest, err := NewParking(s)

	assert.NotNil(t, parkLotTest, ErrorNilObjectReturned)
	assert.Nil(t, err, ErrorExpectedNilObjectOrPointer)
}

func TestNewParkingCostructorWithInvalidValues(t *testing.T) {
	s := -12

	parkLotTest, err := NewParking(s)

	assert.Nil(t, parkLotTest, ErrorExpectedNilObjectOrPointer)
	assert.NotNil(t, err, ErrorNilObjectReturned)
}

func TestShouldReturnTrueOnCallingIsVacantMethodWithNewObject(t *testing.T) {
	size := 10

	parklottest, _ := NewParking(size)

	assert.True(t, parklottest.isVacant(), ErrorExpectedTrue)
}

func TestShouldReturnFalseOnCallingIsVacantWithFullyOccupiedLot(t *testing.T) {
	size := 2

	parklotTest, _ := NewParking(size)

	parklotTest.Occupy()
	parklotTest.Occupy()

	assert.False(t, parklotTest.isVacant(), ErrorExpectedFalse)
}

func TestShouldNotReturnErrorOnRunningOccupyOnNewObject(t *testing.T) {
	size := 2

	parklotTest, _ := NewParking(size)

	err := parklotTest.Occupy()

	assert.Nil(t, err, ErrorUnexpectedError)
}

func TestShouldReturnErrorOnRunningOccupyMethodOnNonVacantObject(t *testing.T) {
	size := 2

	parklotTest, _ := NewParking(size)

	parklotTest.Occupy()
	parklotTest.Occupy()

	err := parklotTest.Occupy()

	assert.NotNil(t, err)

	assert.Equal(t, ErrorNoVacantSpots, err.Error(), ErrorUnexpectedError)
}