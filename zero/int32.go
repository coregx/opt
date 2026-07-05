package zero

import (
	"strconv"

	"github.com/coregx/opt/internal"
)

// Int32 is a nullable int32 where 0 is treated as null.
type Int32 struct {
	Option[int32]
}

// NewInt32 creates an Int32 with the given value and validity.
func NewInt32(i int32, valid bool) Int32 {
	return Int32{New(i, valid)}
}

// Int32From creates an Int32 that is null if i is 0.
func Int32From(i int32) Int32 {
	return Int32{From(i)}
}

// Int32FromPtr creates an Int32 from a pointer. Nil results in null.
func Int32FromPtr(i *int32) Int32 {
	return Int32{FromPtr(i)}
}

// Equal reports whether two Int32s are equal.
func (i Int32) Equal(other Int32) bool {
	return i.Option.Equal(other.Option)
}

// MarshalJSON implements json.Marshaler. Marshals to 0 when null.
func (i Int32) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(int64(i.OrZero()), 10)), nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (i *Int32) UnmarshalJSON(data []byte) error {
	n, _, err := internal.UnmarshalIntJSON(data, 32)
	if err != nil {
		return err
	}
	i.V = int32(n)
	i.Valid = int32(n) != 0
	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (i Int32) MarshalText() ([]byte, error) {
	return []byte(strconv.FormatInt(int64(i.OrZero()), 10)), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (i *Int32) UnmarshalText(text []byte) error {
	n, _, err := internal.UnmarshalIntText(text, 32)
	if err != nil {
		return err
	}
	i.V = int32(n)
	i.Valid = int32(n) != 0
	return nil
}
