package opt

// Map transforms the inner value if valid, returning a new Value.
// If the source is null, returns a null Value of the target type.
func Map[T, U any](v Value[T], fn func(T) U) Value[U] {
	if !v.Valid {
		var zero U
		return New(zero, false)
	}
	return From(fn(v.V))
}

// FlatMap transforms the inner value with a function that itself returns a Value.
// Enables chaining of operations that may produce null.
func FlatMap[T, U any](v Value[T], fn func(T) Value[U]) Value[U] {
	if !v.Valid {
		var zero U
		return New(zero, false)
	}
	return fn(v.V)
}

// Equal compares two Values for equality. Both must be null, or both valid with equal inner values.
func Equal[T comparable](a, b Value[T]) bool {
	return a.Valid == b.Valid && (!a.Valid || a.V == b.V)
}

// OrNull creates a valid Value if v is non-zero, null otherwise.
// Use From(v) when the zero value is meaningful, OrNull(v) when it means "not set".
func OrNull[T comparable](v T) Value[T] {
	var zero T
	return New(v, v != zero)
}
