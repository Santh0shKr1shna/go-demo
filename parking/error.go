package parking

const (
	ErrorInvalidLotSize = "parking lot size cannot be zero or negative"
	ErrorNoVacantSpots = "no vacant spots available. Try again later"
	ErrorNoCarsAreParked = "no cars are parked; the lot is empty"
	ErrorEmptyStrings = "empty strings received as input"

	ErrorNilObjectReturned = "expected non-nil object"
	ErrorExpectedNilObjectOrPointer = "expected nil object or pointer"
	ErrorExpectedTrue = "expected true value"
	ErrorExpectedFalse = "expected false value"
	ErrorMismatchedErrorObject = "expected different error"
	ErrorUnexpectedError = "unexpected error"
)