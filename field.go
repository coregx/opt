package opt

import "encoding/json"

// Field is a three-state nullable type for PATCH API semantics.
//
// Three states:
//   - Absent:  Present=false, Valid=false → field was not in JSON ("don't touch")
//   - Null:    Present=true,  Valid=false → field was explicitly null ("set to NULL")
//   - Value:   Present=true,  Valid=true  → field has a value ("set to value")
//
// The zero value of Field is the Absent state, which is safe by default.
// When encoding/json encounters a missing field, UnmarshalJSON is never called,
// so Present stays false — distinguishing "absent" from "null".
type Field[T any] struct {
	V       T
	Present bool
	Valid   bool
}

// NewField creates a Field with all three components.
func NewField[T any](v T, present, valid bool) Field[T] {
	return Field[T]{V: v, Present: present, Valid: valid}
}

// FieldFrom creates a present, valid Field.
func FieldFrom[T any](v T) Field[T] {
	return Field[T]{V: v, Present: true, Valid: true}
}

// FieldNull creates a present but null Field ("set to NULL").
func FieldNull[T any]() Field[T] {
	return Field[T]{Present: true, Valid: false}
}

// IsAbsent returns true if the field was not present in the input.
func (f Field[T]) IsAbsent() bool {
	return !f.Present
}

// IsNull returns true if the field was present but null.
func (f Field[T]) IsNull() bool {
	return f.Present && !f.Valid
}

// IsValue returns true if the field is present and has a value.
func (f Field[T]) IsValue() bool {
	return f.Present && f.Valid
}

// Or returns the value if present and valid, otherwise the fallback.
func (f Field[T]) Or(fallback T) T {
	if !f.Valid {
		return fallback
	}
	return f.V
}

// OrZero returns the value if valid, otherwise the zero value.
func (f Field[T]) OrZero() T {
	if !f.Valid {
		var zero T
		return zero
	}
	return f.V
}

// Ptr returns a pointer to the value, or nil if absent or null.
func (f Field[T]) Ptr() *T {
	if !f.Valid {
		return nil
	}
	return &f.V
}

// ToValue converts Field to Value, collapsing absent and null into invalid.
func (f Field[T]) ToValue() Value[T] {
	return New(f.V, f.Valid)
}

// IsZero supports omitzero: absent fields are omitted from JSON output.
func (f Field[T]) IsZero() bool {
	return !f.Present
}

// MarshalJSON implements json.Marshaler.
// Absent fields should be omitted via `json:",omitzero"` tag, not marshaled.
// If marshaled explicitly: null when invalid, value when valid.
func (f Field[T]) MarshalJSON() ([]byte, error) {
	if !f.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(f.V)
}

// UnmarshalJSON implements json.Unmarshaler.
// Called only when the field IS present in JSON — so we set Present=true.
func (f *Field[T]) UnmarshalJSON(data []byte) error {
	f.Present = true
	if len(data) > 0 && data[0] == 'n' {
		f.Valid = false
		return nil
	}
	if err := json.Unmarshal(data, &f.V); err != nil {
		return err
	}
	f.Valid = true
	return nil
}
