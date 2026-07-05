package zero

import (
	"encoding/json"
	"fmt"
	"time"
)

// Time is a nullable time.Time where the zero time is treated as null.
type Time struct {
	Option[time.Time]
}

// NewTime creates a Time with the given value and validity.
func NewTime(t time.Time, valid bool) Time {
	return Time{New(t, valid)}
}

// TimeFrom creates a Time that is null if t is the zero time.
func TimeFrom(t time.Time) Time {
	return NewTime(t, !t.IsZero())
}

// TimeFromPtr creates a Time from a pointer. Nil results in null.
func TimeFromPtr(t *time.Time) Time {
	if t == nil {
		return NewTime(time.Time{}, false)
	}
	return TimeFrom(*t)
}

// Equal reports whether two Times represent the same instant.
func (t Time) Equal(other Time) bool {
	return t.OrZero().Equal(other.OrZero())
}

// MarshalJSON implements json.Marshaler. Marshals to zero time when null.
func (t Time) MarshalJSON() ([]byte, error) {
	return t.OrZero().MarshalJSON()
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *Time) UnmarshalJSON(data []byte) error {
	if len(data) > 0 && data[0] == 'n' {
		t.Valid = false
		return nil
	}
	if err := json.Unmarshal(data, &t.V); err != nil {
		return fmt.Errorf("opt/zero: couldn't unmarshal JSON: %w", err)
	}
	t.Valid = !t.V.IsZero()
	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (t Time) MarshalText() ([]byte, error) {
	return t.OrZero().MarshalText()
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
	t.Valid = !t.V.IsZero()
	return nil
}
