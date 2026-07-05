package zero

import (
	"strconv"

	"github.com/coregx/opt/internal"
)

// Int is a nullable int64 where 0 is treated as null.
type Int struct {
	Option[int64]
}

// NewInt creates an Int with the given value and validity.
func NewInt(i int64, valid bool) Int {
	return Int{New(i, valid)}
}

// IntFrom creates an Int that is null if i is 0.
func IntFrom(i int64) Int {
	return Int{From(i)}
}

// IntFromPtr creates an Int from a pointer. Nil results in null.
func IntFromPtr(i *int64) Int {
	return Int{FromPtr(i)}
}

// Equal reports whether two Ints are equal.
func (i Int) Equal(other Int) bool {
	return i.Option.Equal(other.Option)
}

// MarshalJSON implements json.Marshaler. Marshals to 0 when null.
func (i Int) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(i.OrZero(), 10)), nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (i *Int) UnmarshalJSON(data []byte) error {
	n, _, err := internal.UnmarshalIntJSON(data, 64)
	if err != nil {
		return err
	}
	i.V = n
	i.Valid = n != 0
	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (i Int) MarshalText() ([]byte, error) {
	return []byte(strconv.FormatInt(i.OrZero(), 10)), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (i *Int) UnmarshalText(text []byte) error {
	n, _, err := internal.UnmarshalIntText(text, 64)
	if err != nil {
		return err
	}
	i.V = n
	i.Valid = n != 0
	return nil
}
