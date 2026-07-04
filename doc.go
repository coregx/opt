// Package opt provides optional types for Go — a Go-idiomatic Option<T>
// with full SQL, JSON, and Text serialization support.
//
// opt distinguishes between null and zero values: opt.IntFrom(0) is valid, not null.
// For zero-is-null semantics, use the opt/zero subpackage.
//
// # Core Types
//
// Value[T] is the generic foundation. Concrete types (String, Int, Float, Bool, Time,
// Int32, Int16, Byte) provide optimized marshaling for common SQL/JSON types.
//
// Field[T] adds three-state semantics (absent/null/value) for PATCH API support.
//
// # Functional API
//
// Map, FlatMap, and Equal are top-level generic functions for transforming
// and comparing optional values.
//
// # SQL Integration
//
// All types implement sql.Scanner and driver.Valuer through the embedded sql.Null[T],
// working seamlessly with database/sql, pgx, and other SQL drivers.
//
// # JSON Integration
//
// All types marshal to their value when valid and to null when invalid.
// Unmarshal accepts null JSON values. Int and Float also accept string-encoded numbers.
// Use the omitzero struct tag to omit null fields from JSON output.
package opt
