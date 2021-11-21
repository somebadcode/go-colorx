package mathx

import (
	"math"
)

// DefaultPrecision is intended for comparing numbers that are close to 0 and 1 but want to maintain some precision.
const DefaultPrecision = 1e-9

// Equal will compare two floats using DefaultPrecision and return a bool.
func Equal(x, y float64) bool {
	return EqualP(x, y, DefaultPrecision)
}

// EqualP will compare two floats using the specified precision and return a bool.
func EqualP(x, y, p float64) bool {
	return math.Abs(x-y) <= p
}
