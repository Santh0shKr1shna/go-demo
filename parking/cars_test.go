package parking

import (
    "testing"
    "github.com/stretchr/testify/assert"
)
 
func TestNewUnParkingCostructorWithValidValues(t *testing.T) {
    carTest, err := NewCar("123D5","santho")
    assert.NotNil(t, carTest, ErrorNilObjectReturned)
    assert.Nil(t, err,ErrorExpectedNilObjectOrPointer)
}
 
func TestNewUnParkingCostructorWithInvalidValues(t *testing.T) {
    carTest, err := NewCar("","")
 
    assert.Nil(t, carTest, ErrorExpectedNilObjectOrPointer)
    assert.NotNil(t, err, ErrorNilObjectReturned)
}