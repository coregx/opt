package opt_test

import (
	"database/sql"
	"database/sql/driver"
	"encoding"
	"encoding/json"

	"github.com/coregx/opt"
)

// Compile-time interface compliance checks.
// If any of these fail, a concrete type lost an interface
// through embedding changes or field shadowing (see ADR-003).
var (
	_ driver.Valuer = opt.String{}
	_ driver.Valuer = opt.Int{}
	_ driver.Valuer = opt.Int32{}
	_ driver.Valuer = opt.Int16{}
	_ driver.Valuer = opt.Float{}
	_ driver.Valuer = opt.Bool{}
	_ driver.Valuer = opt.Byte{}
	_ driver.Valuer = opt.Time{}

	_ sql.Scanner = (*opt.String)(nil)
	_ sql.Scanner = (*opt.Int)(nil)
	_ sql.Scanner = (*opt.Int32)(nil)
	_ sql.Scanner = (*opt.Int16)(nil)
	_ sql.Scanner = (*opt.Float)(nil)
	_ sql.Scanner = (*opt.Bool)(nil)
	_ sql.Scanner = (*opt.Byte)(nil)
	_ sql.Scanner = (*opt.Time)(nil)

	_ json.Marshaler   = opt.String{}
	_ json.Marshaler   = opt.Int{}
	_ json.Marshaler   = opt.Int32{}
	_ json.Marshaler   = opt.Int16{}
	_ json.Marshaler   = opt.Float{}
	_ json.Marshaler   = opt.Bool{}
	_ json.Marshaler   = opt.Byte{}
	_ json.Marshaler   = opt.Time{}
	_ json.Unmarshaler = (*opt.String)(nil)
	_ json.Unmarshaler = (*opt.Int)(nil)
	_ json.Unmarshaler = (*opt.Int32)(nil)
	_ json.Unmarshaler = (*opt.Int16)(nil)
	_ json.Unmarshaler = (*opt.Float)(nil)
	_ json.Unmarshaler = (*opt.Bool)(nil)
	_ json.Unmarshaler = (*opt.Byte)(nil)
	_ json.Unmarshaler = (*opt.Time)(nil)

	_ encoding.TextMarshaler   = opt.String{}
	_ encoding.TextUnmarshaler = (*opt.String)(nil)
	_ encoding.TextMarshaler   = opt.Int{}
	_ encoding.TextUnmarshaler = (*opt.Int)(nil)
	_ encoding.TextMarshaler   = opt.Int32{}
	_ encoding.TextUnmarshaler = (*opt.Int32)(nil)
	_ encoding.TextMarshaler   = opt.Int16{}
	_ encoding.TextUnmarshaler = (*opt.Int16)(nil)
	_ encoding.TextMarshaler   = opt.Float{}
	_ encoding.TextUnmarshaler = (*opt.Float)(nil)
	_ encoding.TextMarshaler   = opt.Bool{}
	_ encoding.TextUnmarshaler = (*opt.Bool)(nil)
	_ encoding.TextMarshaler   = opt.Byte{}
	_ encoding.TextUnmarshaler = (*opt.Byte)(nil)
	_ encoding.TextMarshaler   = opt.Time{}
	_ encoding.TextUnmarshaler = (*opt.Time)(nil)
)
