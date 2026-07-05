package zero_test

import (
	"database/sql"
	"database/sql/driver"
	"encoding"
	"encoding/json"

	"github.com/coregx/opt/zero"
)

// Compile-time interface compliance checks for zero/ subpackage.
var (
	_ driver.Valuer = zero.String{}
	_ driver.Valuer = zero.Int{}
	_ driver.Valuer = zero.Int32{}
	_ driver.Valuer = zero.Int16{}
	_ driver.Valuer = zero.Float{}
	_ driver.Valuer = zero.Bool{}
	_ driver.Valuer = zero.Byte{}
	_ driver.Valuer = zero.Time{}

	_ sql.Scanner = (*zero.String)(nil)
	_ sql.Scanner = (*zero.Int)(nil)
	_ sql.Scanner = (*zero.Int32)(nil)
	_ sql.Scanner = (*zero.Int16)(nil)
	_ sql.Scanner = (*zero.Float)(nil)
	_ sql.Scanner = (*zero.Bool)(nil)
	_ sql.Scanner = (*zero.Byte)(nil)
	_ sql.Scanner = (*zero.Time)(nil)

	_ json.Marshaler   = zero.String{}
	_ json.Marshaler   = zero.Int{}
	_ json.Marshaler   = zero.Float{}
	_ json.Marshaler   = zero.Bool{}
	_ json.Unmarshaler = (*zero.String)(nil)
	_ json.Unmarshaler = (*zero.Int)(nil)
	_ json.Unmarshaler = (*zero.Float)(nil)
	_ json.Unmarshaler = (*zero.Bool)(nil)

	_ encoding.TextMarshaler   = zero.String{}
	_ encoding.TextUnmarshaler = (*zero.String)(nil)
	_ encoding.TextMarshaler   = zero.Int{}
	_ encoding.TextUnmarshaler = (*zero.Int)(nil)
)
