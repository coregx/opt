package opt

import (
	"encoding/json"
	"fmt"
)

// String is a nullable string with optimized JSON and Text marshaling.
type String struct {
	Value[string]
}

// NewString creates a String with the given value and validity.
func NewString(s string, valid bool) String {
	return String{New(s, valid)}
}

// StringFrom creates a String that is always valid.
func StringFrom(s string) String {
	return String{From(s)}
}

// StringFromPtr creates a String from a pointer. Nil results in null.
func StringFromPtr(s *string) String {
	return String{FromPtr(s)}
}

// Equal reports whether two Strings are equal (both null, or same value).
func (s String) Equal(other String) bool {
	return Equal(s.Value, other.Value)
}

// MarshalJSON implements json.Marshaler.
func (s String) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.V)
}

// UnmarshalJSON implements json.Unmarshaler.
func (s *String) UnmarshalJSON(data []byte) error {
	if len(data) > 0 && data[0] == 'n' {
		s.Valid = false
		return nil
	}
	if err := json.Unmarshal(data, &s.V); err != nil {
		return fmt.Errorf("opt: couldn't unmarshal JSON: %w", err)
	}
	s.Valid = true
	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (s String) MarshalText() ([]byte, error) {
	if !s.Valid {
		return []byte{}, nil
	}
	return []byte(s.V), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *String) UnmarshalText(text []byte) error {
	s.V = string(text)
	s.Valid = s.V != ""
	return nil
}
