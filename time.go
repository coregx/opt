package opt

import (
	"encoding/json"
	"fmt"
	"time"
)

// Time is a nullable time.Time with RFC3339 JSON marshaling.
type Time struct {
	Value[time.Time]
}

// NewTime creates a Time with the given value and validity.
func NewTime(t time.Time, valid bool) Time {
	return Time{New(t, valid)}
}

// TimeFrom creates a Time that is always valid.
func TimeFrom(t time.Time) Time {
	return Time{From(t)}
}

// TimeFromPtr creates a Time from a pointer. Nil results in null.
func TimeFromPtr(t *time.Time) Time {
	return Time{FromPtr(t)}
}

// Equal reports whether two Times represent the same instant (timezone-independent).
func (t Time) Equal(other Time) bool {
	return t.Valid == other.Valid && (!t.Valid || t.V.Equal(other.V))
}

// ExactEqual reports whether two Times are exactly equal (including timezone).
func (t Time) ExactEqual(other Time) bool {
	return t.Valid == other.Valid && (!t.Valid || t.V == other.V)
}

// MarshalJSON implements json.Marshaler.
func (t Time) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}
	return t.V.MarshalJSON()
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *Time) UnmarshalJSON(data []byte) error {
	if len(data) > 0 && data[0] == 'n' {
		t.Valid = false
		return nil
	}
	if err := json.Unmarshal(data, &t.V); err != nil {
		return fmt.Errorf("opt: couldn't unmarshal JSON: %w", err)
	}
	t.Valid = true
	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (t Time) MarshalText() ([]byte, error) {
	if !t.Valid {
		return []byte{}, nil
	}
	return t.V.MarshalText()
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (t *Time) UnmarshalText(text []byte) error {
	str := string(text)
	if str == "" || str == "null" {
		t.Valid = false
		return nil
	}
	if err := t.V.UnmarshalText(text); err != nil {
		return err
	}
	t.Valid = true
	return nil
}
