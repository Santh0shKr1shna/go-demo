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

	car1, _ := NewCar("TN02", "Girish")
	car2, _ := NewCar("TN05", "Phani")

	parklotTest, _ := NewParking(size)

	parklotTest.Park(*car1)
	parklotTest.Park(*car2)

	assert.False(t, parklotTest.isVacant(), ErrorExpectedFalse)
}

func TestShouldNotReturnErrorOnRunningOccupyOnNewObject(t *testing.T) {
	size := 2

	car1, _ := NewCar("TN12", "Nandhan")

	parklotTest, _ := NewParking(size)

	err := parklotTest.Park(*car1)

	assert.Nil(t, err, ErrorUnexpectedError)
}

func TestShouldReturnErrorOnRunningOccupyMethodOnNonVacantObject(t *testing.T) {
	size := 2

	parklotTest, _ := NewParking(size)

	car1, _ := NewCar("TN02", "Girish")
	car2, _ := NewCar("TN05", "Phani")
	car3, _ := NewCar("TN12", "Nandhan")

	parklotTest.Park(*car1)
	parklotTest.Park(*car2)

	err := parklotTest.Park(*car3)

	assert.NotNil(t, err)

	assert.Equal(t, ErrorNoVacantSpots, err.Error(), ErrorUnexpectedError)
}

// TODO: contains - check both cases



// TODO: Unpark - check three cases:
// 		1) with 0 occupancy
// 		2) without particular car in lot
// 		3) positive flow

func TestShouldReturnErrorToUnparkWith0Occupancy(t *testing.T) {
	// Arrange
	size := 2
	parklotTest, _ := NewParking(size)
	car1, _ := NewCar("TN02", "Girish")

	// Act
	err := parklotTest.UnPark(*car1)

	assert.NotNil(t, err)
	assert.Equal(t, ErrorNoCarsAreParked, err.Error())
}

func TestShouldReturnErrorAsTheParticularCarIsNotInTheLot(t *testing.T) {
	// Arrange
	size := 2
	parklotTest, _ := NewParking(size)
	car1, _ := NewCar("TN02", "Girish")
	car2, _ := NewCar("TN05", "Phani")

	parklotTest.Park(*car1)

	// Act
	err := parklotTest.UnPark(*car2)

	// Assert
	assert.NotNil(t, err)
	assert.Equal(t, ErrorThisCarIsNotParkedHere, err.Error())
}

func TestShouldCheckTheExistenceOfCarUsingContainsBeforeAndAfterUnparking(t *testing.T) {
	// Arrange
	size := 2
	parklotTest, _ := NewParking(size)
	demoCar, _ := NewCar("TN02", "Girish")
	actorCar, _ := NewCar("TN05", "Phani")

	parklotTest.Park(*demoCar)
	parklotTest.Park(*actorCar)

	// Assert existence before acting
	assert.True(t, parklotTest.Contains(*actorCar))

	// Act
	parklotTest.UnPark(*actorCar)

	// Assert
	assert.False(t, parklotTest.Contains(*actorCar))
}

func TestShouldVerifyParkingLotOwnerGetsNotified(t *testing.T) {
	// arrange
	size := 2
	parklotTest, _ := NewParking(size)
	demoCar1, _ := NewCar("TN02", "Girish")
	demoCar2, _ := NewCar("TN05", "Phani")

	owner := Owner{}

	parklotTest.register(&owner)

	// act
	parklotTest.Park(*demoCar1)
	parklotTest.Park(*demoCar2)

	// assert
	assert.Equal(t, owner.msg, "Parking lot full")
}

func TestShouldVerifyTrafficPoliceGetsNotified(t *testing.T) {
	// arrange
	size := 2
	parklotTest, _ := NewParking(size)
	demoCar1, _ := NewCar("TN02", "Girish")
	demoCar2, _ := NewCar("TN05", "Phani")

	police := TrafficPolice{}

	parklotTest.register(&police)

	parklotTest.Park(*demoCar1)
	// verification
	assert.NotEqual(t, police.msg, "Parking lot full")

	// action
	parklotTest.Park(*demoCar2)

	// assert
	assert.Equal(t, police.msg, "Parking lot full")
}

func TestShouldVerifyBothTrafficPoliceAndOwnerAreNotified(t *testing.T) {
	size := 2
	parklotTest, _ := NewParking(size)
	demoCar1, _ := NewCar("TN02", "Girish")

	owner := Owner{}
	police := TrafficPolice{}

	parklotTest.register(&owner)
	parklotTest.register(&police)

	parklotTest.Park(*demoCar1)

	assert.NotEqual(t, police.msg, "Parking lot full")
	assert.NotEqual(t, owner.msg, "Parking lot full")
}