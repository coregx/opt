package opt

import (
	"strconv"

	"github.com/coregx/opt/internal"
)

// Int32 is a nullable int32 with optimized JSON marshaling.
type Int32 struct {
	Option[int32]
}

// NewInt32 creates an Int32 with the given value and validity.
func NewInt32(i int32, valid bool) Int32 {
	return Int32{New(i, valid)}
}

// Int32From creates an Int32 that is always valid.
func Int32From(i int32) Int32 {
	return Int32{From(i)}
}

// Int32FromPtr creates an Int32 from a pointer. Nil results in null.
func Int32FromPtr(i *int32) Int32 {
	return Int32{FromPtr(i)}
}

// Int32OrNull creates a valid Int32 if n is non-zero, null otherwise.
func Int32OrNull(n int32) Int32 {
	return NewInt32(n, n != 0)
}

// Equal reports whether two Int32s are equal.
func (i Int32) Equal(other Int32) bool {
	return Equal(i.Option, other.Option)
}

// MarshalJSON implements json.Marshaler.
func (i Int32) MarshalJSON() ([]byte, error) {
	if !i.Valid {
		return []byte("null"), nil
	}
	return []byte(strconv.FormatInt(int64(i.V), 10)), nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (i *Int32) UnmarshalJSON(data []byte) error {
	n, valid, err := internal.UnmarshalIntJSON(data, 32)
	if err != nil {
		return err
	}
	i.V = int32(n)
	i.Valid = valid
	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (i Int32) MarshalText() ([]byte, error) {
	if !i.Valid {
		return []byte{}, nil
	}
	return []byte(strconv.FormatInt(int64(i.V), 10)), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (i *Int32) UnmarshalText(text []byte) error {
	n, valid, err := internal.UnmarshalIntText(text, 32)
	if err != nil {
		return err
	}
	i.V = int32(n)
	i.Valid = valid
	return nil
}
