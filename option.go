// Package opt provides optional types for Go — a Go-idiomatic Option<T>
// with full SQL, JSON, and Text serialization support.
package opt

import (
	"database/sql"
	"encoding/json"
)

// Option is a generic nullable type backed by sql.Null[T].
// It marshals to JSON null when invalid and to the value when valid.
// For SQL, it inherits Scanner and Valuer from sql.Null[T].
type Option[T any] struct {
	sql.Null[T]
}

// New creates an Option with the given value and validity.
func New[T any](v T, valid bool) Option[T] {
	return Option[T]{sql.Null[T]{V: v, Valid: valid}}
}

// From creates an Option that is always valid.
func From[T any](v T) Option[T] {
	return New(v, true)
}

// FromPtr creates an Option from a pointer. Nil pointer results in a null Option.
func FromPtr[T any](ptr *T) Option[T] {
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

// OrZero returns the inner value if valid, otherwise the zero value of T.
func (v Option[T]) OrZero() T {
	if !v.Valid {
		var zero T
		return zero
	}
	return v.V
}

// OrElse returns the inner value if valid, otherwise calls fn and returns its result.
func (v Option[T]) OrElse(fn func() T) T {
	if !v.Valid {
		return fn()
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

// IsZero returns true when the Option is null. Supports Go 1.24+ omitzero.
func (v Option[T]) IsZero() bool {
	return !v.Valid
}

// MarshalJSON implements json.Marshaler. Encodes null when invalid.
func (v Option[T]) MarshalJSON() ([]byte, error) {
	if !v.Valid {
		return []byte("null"), nil
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
	v.Valid = true
	return nil
}
