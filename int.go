package opt

import (
	"strconv"

	"github.com/coregx/opt/internal"
)

// Int is a nullable int64 with optimized JSON marshaling.
// It accepts both numbers and string-encoded numbers in JSON.
type Int struct {
	Value[int64]
}

// NewInt creates an Int with the given value and validity.
func NewInt(i int64, valid bool) Int {
	return Int{New(i, valid)}
}

// IntFrom creates an Int that is always valid.
func IntFrom(i int64) Int {
	return Int{From(i)}
}

// IntFromPtr creates an Int from a pointer. Nil results in null.
func IntFromPtr(i *int64) Int {
	return Int{FromPtr(i)}
}

// Equal reports whether two Ints are equal.
func (i Int) Equal(other Int) bool {
	return Equal(i.Value, other.Value)
}

// MarshalJSON implements json.Marshaler.
func (i Int) MarshalJSON() ([]byte, error) {
	if !i.Valid {
		return []byte("null"), nil
	}
	return []byte(strconv.FormatInt(i.V, 10)), nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (i *Int) UnmarshalJSON(data []byte) error {
	n, valid, err := internal.UnmarshalIntJSON(data, 64)
	if err != nil {
		return err
	}
	i.V = n
	i.Valid = valid
	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (i Int) MarshalText() ([]byte, error) {
	if !i.Valid {
		return []byte{}, nil
	}
	return []byte(strconv.FormatInt(i.V, 10)), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (i *Int) UnmarshalText(text []byte) error {
	n, valid, err := internal.UnmarshalIntText(text, 64)
	if err != nil {
		return err
	}
	i.V = n
	i.Valid = valid
	return nil
}
