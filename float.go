package opt

import (
	"encoding/json"
	"math"
	"reflect"
	"strconv"

	"github.com/coregx/opt/internal"
)

// Float is a nullable float64 with optimized JSON marshaling.
// It rejects Inf and NaN during JSON serialization.
type Float struct {
	Value[float64]
}

// NewFloat creates a Float with the given value and validity.
func NewFloat(f float64, valid bool) Float {
	return Float{New(f, valid)}
}

// FloatFrom creates a Float that is always valid.
func FloatFrom(f float64) Float {
	return Float{From(f)}
}

// FloatFromPtr creates a Float from a pointer. Nil results in null.
func FloatFromPtr(f *float64) Float {
	return Float{FromPtr(f)}
}

// FloatOrNull creates a valid Float if f is non-zero, null otherwise.
func FloatOrNull(f float64) Float {
	return NewFloat(f, f != 0)
}

// Equal reports whether two Floats are equal.
func (f Float) Equal(other Float) bool {
	return Equal(f.Value, other.Value)
}

// MarshalJSON implements json.Marshaler. Rejects Inf and NaN.
func (f Float) MarshalJSON() ([]byte, error) {
	if !f.Valid {
		return []byte("null"), nil
	}
	if math.IsInf(f.V, 0) || math.IsNaN(f.V) {
		return nil, &json.UnsupportedValueError{
			Value: reflect.ValueOf(f.V),
			Str:   strconv.FormatFloat(f.V, 'g', -1, 64),
		}
	}
	return []byte(strconv.FormatFloat(f.V, 'f', -1, 64)), nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (f *Float) UnmarshalJSON(data []byte) error {
	n, valid, err := internal.UnmarshalFloatJSON(data)
	if err != nil {
		return err
	}
	f.V = n
	f.Valid = valid
	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (f Float) MarshalText() ([]byte, error) {
	if !f.Valid {
		return []byte{}, nil
	}
	return []byte(strconv.FormatFloat(f.V, 'f', -1, 64)), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (f *Float) UnmarshalText(text []byte) error {
	n, valid, err := internal.UnmarshalFloatText(text)
	if err != nil {
		return err
	}
	f.V = n
	f.Valid = valid
	return nil
}
