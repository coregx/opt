package zero

import (
	"encoding/json"
	"math"
	"reflect"
	"strconv"

	"github.com/coregx/opt/internal"
)

// Float is a nullable float64 where 0 is treated as null.
type Float struct {
	Option[float64]
}

// NewFloat creates a Float with the given value and validity.
func NewFloat(f float64, valid bool) Float {
	return Float{New(f, valid)}
}

// FloatFrom creates a Float that is null if f is 0.
func FloatFrom(f float64) Float {
	return Float{From(f)}
}

// FloatFromPtr creates a Float from a pointer. Nil results in null.
func FloatFromPtr(f *float64) Float {
	return Float{FromPtr(f)}
}

// Equal reports whether two Floats are equal.
func (f Float) Equal(other Float) bool {
	return f.Option.Equal(other.Option)
}

// MarshalJSON implements json.Marshaler. Marshals to 0 when null.
func (f Float) MarshalJSON() ([]byte, error) {
	v := f.OrZero()
	if math.IsInf(v, 0) || math.IsNaN(v) {
		return nil, &json.UnsupportedValueError{
			Value: reflect.ValueOf(v),
			Str:   strconv.FormatFloat(v, 'g', -1, 64),
		}
	}
	return []byte(strconv.FormatFloat(v, 'f', -1, 64)), nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (f *Float) UnmarshalJSON(data []byte) error {
	n, _, err := internal.UnmarshalFloatJSON(data)
	if err != nil {
		return err
	}
	f.V = n
	f.Valid = n != 0
	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (f Float) MarshalText() ([]byte, error) {
	return []byte(strconv.FormatFloat(f.OrZero(), 'f', -1, 64)), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (f *Float) UnmarshalText(text []byte) error {
	n, _, err := internal.UnmarshalFloatText(text)
	if err != nil {
		return err
	}
	f.V = n
	f.Valid = n != 0
	return nil
}
