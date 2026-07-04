package opt

import (
	"encoding/json"
	"fmt"
)

// Byte is a nullable byte (uint8) with JSON marshaling.
type Byte struct {
	Value[byte]
}

// NewByte creates a Byte with the given value and validity.
func NewByte(b byte, valid bool) Byte {
	return Byte{New(b, valid)}
}

// ByteFrom creates a Byte that is always valid.
func ByteFrom(b byte) Byte {
	return Byte{From(b)}
}

// ByteFromPtr creates a Byte from a pointer. Nil results in null.
func ByteFromPtr(b *byte) Byte {
	return Byte{FromPtr(b)}
}

// ByteOrNull creates a valid Byte if b is non-zero, null otherwise.
func ByteOrNull(b byte) Byte {
	return NewByte(b, b != 0)
}

// Equal reports whether two Bytes are equal.
func (b Byte) Equal(other Byte) bool {
	return Equal(b.Value, other.Value)
}

// MarshalJSON implements json.Marshaler.
func (b Byte) MarshalJSON() ([]byte, error) {
	if !b.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(b.V)
}

// UnmarshalJSON implements json.Unmarshaler.
func (b *Byte) UnmarshalJSON(data []byte) error {
	if len(data) > 0 && data[0] == 'n' {
		b.Valid = false
		return nil
	}
	var n int64
	if err := json.Unmarshal(data, &n); err != nil {
		return fmt.Errorf("opt: couldn't unmarshal JSON: %w", err)
	}
	if n < 0 || n > 255 {
		return fmt.Errorf("opt: byte value out of range: %d", n)
	}
	b.V = byte(n)
	b.Valid = true
	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (b Byte) MarshalText() ([]byte, error) {
	if !b.Valid {
		return []byte{}, nil
	}
	return []byte{b.V}, nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (b *Byte) UnmarshalText(text []byte) error {
	if len(text) == 0 || string(text) == "null" {
		b.Valid = false
		return nil
	}
	if len(text) != 1 {
		return fmt.Errorf("opt: invalid Byte text: expected 1 byte, got %d", len(text))
	}
	b.V = text[0]
	b.Valid = true
	return nil
}
