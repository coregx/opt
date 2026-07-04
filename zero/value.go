// Package zero provides nullable types where zero values are treated as null.
//
// Unlike package opt where zero values are valid (IntFrom(0) is valid),
// in package zero, zero values are null (IntFrom(0) is invalid).
// Invalid values marshal to their zero representation ("", 0, false),
// not to JSON null.
package zero

import (
	"database/sql"
	"encoding/json"
)

// Value is a generic nullable type where the zero value of T is considered null.
type Value[T comparable] struct {
	sql.Null[T]
}

// New creates a Value with the given value and validity.
func New[T comparable](v T, valid bool) Value[T] {
	return Value[T]{sql.Null[T]{V: v, Valid: valid}}
}

// From creates a Value that is null if v equals the zero value of T.
func From[T comparable](v T) Value[T] {
	var zero T
	return New(v, v != zero)
}

// FromPtr creates a Value from a pointer. Nil results in null.
func FromPtr[T comparable](ptr *T) Value[T] {
	if ptr == nil {
		var zero T
		return New(zero, false)
	}
	return New(*ptr, true)
}

// Or returns the inner value if valid, otherwise the fallback.
func (v Value[T]) Or(fallback T) T {
	if !v.Valid {
		return fallback
	}
	return v.V
}

// OrZero returns the inner value if valid, otherwise the zero value.
func (v Value[T]) OrZero() T {
	if !v.Valid {
		var zero T
		return zero
	}
	return v.V
}

// SetValid sets the value and marks it as valid.
func (v *Value[T]) SetValid(val T) {
	v.V = val
	v.Valid = true
}

// Ptr returns a pointer to the value, or nil if null.
func (v Value[T]) Ptr() *T {
	if !v.Valid {
		return nil
	}
	return &v.V
}

// IsZero returns true when null OR when the value equals zero.
func (v Value[T]) IsZero() bool {
	var zero T
	return !v.Valid || v.V == zero
}

// Equal reports whether two Values are equal by comparing OrZero results.
func (v Value[T]) Equal(other Value[T]) bool {
	return v.OrZero() == other.OrZero()
}

// MarshalJSON implements json.Marshaler.
// Marshals to the zero value of T when null, not to JSON "null".
func (v Value[T]) MarshalJSON() ([]byte, error) {
	if !v.Valid {
		var zero T
		return json.Marshal(zero)
	}
	return json.Marshal(v.V)
}

// UnmarshalJSON implements json.Unmarshaler.
func (v *Value[T]) UnmarshalJSON(data []byte) error {
	if len(data) > 0 && data[0] == 'n' {
		v.Valid = false
		return nil
	}
	if err := json.Unmarshal(data, &v.V); err != nil {
		return err
	}
	var zero T
	v.Valid = v.V != zero
	return nil
}
