package zero

import (
	"errors"
	"fmt"
)

// Bool is a nullable bool where false is treated as null.
type Bool struct {
	Value[bool]
}

// NewBool creates a Bool with the given value and validity.
func NewBool(b bool, valid bool) Bool {
	return Bool{New(b, valid)}
}

// BoolFrom creates a Bool that is null if b is false.
func BoolFrom(b bool) Bool {
	return Bool{From(b)}
}

// BoolFromPtr creates a Bool from a pointer. Nil results in null.
func BoolFromPtr(b *bool) Bool {
	return Bool{FromPtr(b)}
}

// Equal reports whether two Bools are equal.
func (b Bool) Equal(other Bool) bool {
	return b.Value.Equal(other.Value)
}

// MarshalJSON implements json.Marshaler. Marshals to false when null.
func (b Bool) MarshalJSON() ([]byte, error) {
	if !b.Valid || !b.V {
		return []byte("false"), nil
	}
	return []byte("true"), nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (b *Bool) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return errors.New("opt/zero: empty JSON input")
	}
	switch data[0] {
	case 'n':
		b.Valid = false
		return nil
	case 't':
		b.V = true
		b.Valid = true
		return nil
	case 'f':
		b.V = false
		b.Valid = false
		return nil
	default:
		return fmt.Errorf("opt/zero: cannot unmarshal %s into Bool", string(data))
	}
}

// MarshalText implements encoding.TextMarshaler.
func (b Bool) MarshalText() ([]byte, error) {
	if b.V && b.Valid {
		return []byte("true"), nil
	}
	return []byte("false"), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (b *Bool) UnmarshalText(text []byte) error {
	str := string(text)
	switch str {
	case "true":
		b.V = true
		b.Valid = true
	default:
		b.V = false
		b.Valid = false
	}
	return nil
}
