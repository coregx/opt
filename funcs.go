package opt

// Map transforms the inner value if valid, returning a new Option.
// If the source is null, returns a null Option of the target type.
func Map[T, U any](v Option[T], fn func(T) U) Option[U] {
	if !v.Valid {
		var zero U
		return New(zero, false)
	}
	return From(fn(v.V))
}

// FlatMap transforms the inner value with a function that itself returns an Option.
// Enables chaining of operations that may produce null.
func FlatMap[T, U any](v Option[T], fn func(T) Option[U]) Option[U] {
	if !v.Valid {
		var zero U
		return New(zero, false)
	}
	return fn(v.V)
}

// Equal compares two Options for equality. Both must be null, or both valid with equal inner values.
func Equal[T comparable](a, b Option[T]) bool {
	return a.Valid == b.Valid && (!a.Valid || a.V == b.V)
}

// OrNull creates a valid Option if v is non-zero, null otherwise.
// Use From(v) when the zero value is meaningful, OrNull(v) when it means "not set".
func OrNull[T comparable](v T) Option[T] {
	var zero T
	return New(v, v != zero)
}

// FieldFromOption creates a present Field from an Option (lossless upcast).
// Valid Option becomes a present+valid Field. Invalid Option becomes a present+null Field (not absent).
func FieldFromOption[T any](v Option[T]) Field[T] {
	return NewField(v.V, true, v.Valid)
}
