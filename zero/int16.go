package zero

import (
	"strconv"

	"github.com/coregx/opt/internal"
)

// Int16 is a nullable int16 where 0 is treated as null.
type Int16 struct {
	Option[int16]
}

// NewInt16 creates an Int16 with the given value and validity.
func NewInt16(i int16, valid bool) Int16 {
	return Int16{New(i, valid)}
}

// Int16From creates an Int16 that is null if i is 0.
func Int16From(i int16) Int16 {
	return Int16{From(i)}
}

// Int16FromPtr creates an Int16 from a pointer. Nil results in null.
func Int16FromPtr(i *int16) Int16 {
	return Int16{FromPtr(i)}
}

// Equal reports whether two Int16s are equal.
func (i Int16) Equal(other Int16) bool {
	return i.Option.Equal(other.Option)
}

// MarshalJSON implements json.Marshaler. Marshals to 0 when null.
func (i Int16) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(int64(i.OrZero()), 10)), nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (i *Int16) UnmarshalJSON(data []byte) error {
	n, _, err := internal.UnmarshalIntJSON(data, 16)
	if err != nil {
		return err
	}
	i.V = int16(n)
	i.Valid = int16(n) != 0
	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (i Int16) MarshalText() ([]byte, error) {
	return []byte(strconv.FormatInt(int64(i.OrZero()), 10)), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (i *Int16) UnmarshalText(text []byte) error {
	n, _, err := internal.UnmarshalIntText(text, 16)
	if err != nil {
		return err
	}
	i.V = int16(n)
	i.Valid = int16(n) != 0
	return nil
}
