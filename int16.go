package opt

import (
	"strconv"

	"github.com/coregx/opt/internal"
)

// Int16 is a nullable int16 with optimized JSON marshaling.
type Int16 struct {
	Value[int16]
}

// NewInt16 creates an Int16 with the given value and validity.
func NewInt16(i int16, valid bool) Int16 {
	return Int16{New(i, valid)}
}

// Int16From creates an Int16 that is always valid.
func Int16From(i int16) Int16 {
	return Int16{From(i)}
}

// Int16FromPtr creates an Int16 from a pointer. Nil results in null.
func Int16FromPtr(i *int16) Int16 {
	return Int16{FromPtr(i)}
}

// Int16OrNull creates a valid Int16 if n is non-zero, null otherwise.
func Int16OrNull(n int16) Int16 {
	return NewInt16(n, n != 0)
}

// Equal reports whether two Int16s are equal.
func (i Int16) Equal(other Int16) bool {
	return Equal(i.Value, other.Value)
}

// MarshalJSON implements json.Marshaler.
func (i Int16) MarshalJSON() ([]byte, error) {
	if !i.Valid {
		return []byte("null"), nil
	}
	return []byte(strconv.FormatInt(int64(i.V), 10)), nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (i *Int16) UnmarshalJSON(data []byte) error {
	n, valid, err := internal.UnmarshalIntJSON(data, 16)
	if err != nil {
		return err
	}
	i.V = int16(n)
	i.Valid = valid
	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (i Int16) MarshalText() ([]byte, error) {
	if !i.Valid {
		return []byte{}, nil
	}
	return []byte(strconv.FormatInt(int64(i.V), 10)), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (i *Int16) UnmarshalText(text []byte) error {
	n, valid, err := internal.UnmarshalIntText(text, 16)
	if err != nil {
		return err
	}
	i.V = int16(n)
	i.Valid = valid
	return nil
}
