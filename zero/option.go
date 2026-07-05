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

// Option is a generic nullable type where the zero value of T is considered null.
type Option[T comparable] struct {
	sql.Null[T]
}

// New creates an Option with the given value and validity.
func New[T comparable](v T, valid bool) Option[T] {
	return Option[T]{sql.Null[T]{V: v, Valid: valid}}
}

// From creates an Option that is null if v equals the zero value of T.
func From[T comparable](v T) Option[T] {
	var zero T
	return New(v, v != zero)
}

// FromPtr creates an Option from a pointer. Nil results in null.
func FromPtr[T comparable](ptr *T) Option[T] {
	if ptr == nil {
		var zero T
		return New(zero, false)
	}
	return New(*ptr, true)
}

// Or returns the inner value if valid, otherwise the fallback.
func (v Option[T]) Or(fallback T) T {
	if !v.Valid {
		return fallback
	}
	return v.V
}

// OrZero returns the inner value if valid, otherwise the zero value.
func (v Option[T]) OrZero() T {
	if !v.Valid {
		var zero T
		return zero
	}
	return v.V
}

// SetValid sets the value and marks it as valid.
func (v *Option[T]) SetValid(val T) {
	v.V = val
	v.Valid = true
}

// Ptr returns a pointer to the value, or nil if null.
func (v Option[T]) Ptr() *T {
	if !v.Valid {
		return nil
	}
	return &v.V
}

// IsZero returns true when null OR when the value equals zero.
func (v Option[T]) IsZero() bool {
	var zero T
	return !v.Valid || v.V == zero
}

// Equal reports whether two Options are equal by comparing OrZero results.
func (v Option[T]) Equal(other Option[T]) bool {
	return v.OrZero() == other.OrZero()
}

// MarshalJSON implements json.Marshaler.
// Marshals to the zero value of T when null, not to JSON "null".
func (v Option[T]) MarshalJSON() ([]byte, error) {
	if !v.Valid {
		var zero T
		return json.Marshal(zero)
	}
	return json.Marshal(v.V)
}

// UnmarshalJSON implements json.Unmarshaler.
func (v *Option[T]) UnmarshalJSON(data []byte) error {
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
