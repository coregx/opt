// Package opt provides optional types for Go — a Go-idiomatic Option<T>
// with full SQL, JSON, and Text serialization support.
package opt

import (
	"database/sql"
	"encoding/json"
)

// Value is a generic nullable type backed by sql.Null[T].
// It marshals to JSON null when invalid and to the value when valid.
// For SQL, it inherits Scanner and Valuer from sql.Null[T].
type Value[T any] struct {
	sql.Null[T]
}

// New creates a Value with the given value and validity.
func New[T any](v T, valid bool) Value[T] {
	return Value[T]{sql.Null[T]{V: v, Valid: valid}}
}

// From creates a Value that is always valid.
func From[T any](v T) Value[T] {
	return New(v, true)
}

// FromPtr creates a Value from a pointer. Nil pointer results in a null Value.
func FromPtr[T any](ptr *T) Value[T] {
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

// OrZero returns the inner value if valid, otherwise the zero value of T.
func (v Value[T]) OrZero() T {
	if !v.Valid {
		var zero T
		return zero
	}
	return v.V
}

// OrElse returns the inner value if valid, otherwise calls fn and returns its result.
func (v Value[T]) OrElse(fn func() T) T {
	if !v.Valid {
		return fn()
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

// IsZero returns true when the Value is null. Supports Go 1.24+ omitzero.
func (v Value[T]) IsZero() bool {
	return !v.Valid
}

// MarshalJSON implements json.Marshaler. Encodes null when invalid.
func (v Value[T]) MarshalJSON() ([]byte, error) {
	if !v.Valid {
		return []byte("null"), nil
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
	v.Valid = true
	return nil
}
