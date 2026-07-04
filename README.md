<h1 align="center">opt</h1>

<p align="center">
  <strong>Optional types for Go — a Go-idiomatic Option&lt;T&gt;</strong><br>
  Full SQL + JSON integration. Three-state PATCH support. Zero dependencies beyond stdlib.
</p>

<p align="center">
  <a href="https://github.com/coregx/opt/actions"><img src="https://github.com/coregx/opt/actions/workflows/ci.yml/badge.svg" alt="CI"></a>
  <a href="https://pkg.go.dev/github.com/coregx/opt"><img src="https://pkg.go.dev/badge/github.com/coregx/opt.svg" alt="Go Reference"></a>
  <a href="https://goreportcard.com/report/github.com/coregx/opt"><img src="https://goreportcard.com/badge/github.com/coregx/opt" alt="Go Report Card"></a>
  <a href="https://github.com/coregx/opt/blob/main/LICENSE"><img src="https://img.shields.io/badge/License-MIT-yellow.svg" alt="License"></a>
  <a href="https://github.com/coregx/opt"><img src="https://img.shields.io/badge/Go-1.24+-00ADD8" alt="Go 1.24+"></a>
</p>

---

## Why opt?

Go's `sql.Null[T]` handles SQL but [will never get JSON support](https://github.com/golang/go/issues/68375) (closed as infeasible). Pointers (`*string`) work but have overhead and awkward ergonomics. `opt` fills this gap permanently.

```go
type User struct {
    Name  opt.String `json:"name"`
    Email opt.String `json:"email,omitzero"`
    Age   opt.Int    `json:"age"`
}
// {"name":"Alice","age":null}  — Email omitted (omitzero), Age is explicit null
```

## Features

- **Generic foundation** — `Value[T]` works with any type via `sql.Null[T]`
- **9 concrete types** — String, Int, Int32, Int16, Float, Bool, Byte, Time
- **Three-state `Field[T]`** — distinguish absent / null / value for PATCH APIs
- **Functional API** — `Map`, `FlatMap`, `Equal` for composable transformations
- **Zero dependencies** — only Go stdlib (`database/sql`, `encoding/json`)
- **SQL-ready** — `Scanner`/`Valuer` via `sql.Null[T]`, works with pgx, database/sql
- **JSON-ready** — proper null marshaling, `omitzero` support (Go 1.24+)
- **json/v2 compatible** — works with `encoding/json/v2` without changes
- **`zero/` subpackage** — alternative semantics where zero value = null
- **Benchmarked** — zero-allocation unmarshal, Bool marshal in <1ns

## Installation

```bash
go get github.com/coregx/opt
```

**Requires Go 1.24+**

## Quick Start

```go
package main

import (
    "encoding/json"
    "fmt"

    "github.com/coregx/opt"
)

func main() {
    // Create optional values
    name := opt.StringFrom("Alice")
    age := opt.NewInt(0, false) // null

    // Safe access with fallbacks
    fmt.Println(name.Or("unknown")) // "Alice"
    fmt.Println(age.Or(18))         // 18

    // JSON marshaling — just works
    type User struct {
        Name opt.String `json:"name"`
        Age  opt.Int    `json:"age"`
    }
    data, _ := json.Marshal(User{Name: name, Age: age})
    fmt.Println(string(data)) // {"name":"Alice","age":null}
}
```

## API

### Constructors

```go
opt.From(value)          // Always valid: opt.From("hello"), opt.From(42)
opt.New(value, valid)    // Explicit: opt.New("", false) → null
opt.FromPtr(ptr)         // From pointer: nil → null, &v → valid
opt.OrNull(value)        // Zero value → null: opt.OrNull("") → null, opt.OrNull(42) → valid

// Type-specific shortcuts
opt.StringFrom("hello")      opt.StringOrNull("")     // "" → null
opt.IntFrom(42)              opt.IntOrNull(0)         // 0 → null
opt.FloatFrom(3.14)          opt.FloatOrNull(0.0)     // 0 → null
opt.BoolFrom(true)           opt.BoolOrNull(false)    // false → null
opt.TimeFrom(t)              opt.TimeOrNull(t)        // zero time → null
opt.ByteFrom(0x42)           opt.ByteOrNull(0)        // 0 → null
```

`From` = value is always valid. `OrNull` = zero value means "not set" → null. Choose based on your semantics.

### Value Access

```go
v.Or(fallback)           // Value or fallback
v.OrZero()               // Value or zero value of T
v.OrElse(func() T)       // Value or lazy-computed fallback
v.Ptr()                  // *T or nil
v.IsZero()               // true when null (for omitzero)
```

### Functional

```go
opt.Map(v, func(T) U) Value[U]              // Transform if valid
opt.FlatMap(v, func(T) Value[U]) Value[U]    // Chain optional operations
opt.Equal(a, b)                               // Nil-safe comparison
```

### Three-State Field (PATCH API)

```go
type PatchUser struct {
    Name  opt.Field[string] `json:"name,omitzero"`
    Email opt.Field[string] `json:"email,omitzero"`
}

// {}                    → Name.IsAbsent()=true  — don't touch
// {"name": null}        → Name.IsNull()=true    — set to NULL
// {"name": "Alice"}     → Name.IsValue()=true   — set to "Alice"
```

### Zero Subpackage

```go
import "github.com/coregx/opt/zero"

s := zero.StringFrom("")   // Invalid — empty string = null
i := zero.IntFrom(0)       // Invalid — zero = null
data, _ := json.Marshal(i) // "0" (not "null")
```

| Behavior | `opt` | `opt/zero` |
|----------|-------|-----------|
| `From("")` | Valid (empty string) | Invalid (empty = null) |
| `From(0)` | Valid (zero int) | Invalid (zero = null) |
| Marshal null | `null` | `""` / `0` / `false` |

## SQL Usage

```go
// Works with database/sql
var user struct {
    Name opt.String
    Age  opt.Int
}
db.QueryRow("SELECT name, age FROM users WHERE id=$1", id).Scan(&user.Name, &user.Age)

// Also works with pgx, sqlx, and other drivers
```

## Comparison with Alternatives

| Feature | opt | guregu/null | `*T` (pointer) | `sql.Null[T]` |
|---------|-----|-----------|----------------|---------------|
| Generic `Value[T]` | **Full** | Partial (MarshalText commented out) | N/A | No JSON |
| Three-state (PATCH) | **`Field[T]`** | No | No | No |
| Map / FlatMap | **Yes** | No | No | No |
| OrElse (lazy) | **Yes** | No | No | No |
| JSON marshal | **Yes** | Yes | Yes | Broken (`{"V":...,"Valid":...}`) |
| SQL Scanner/Valuer | **Yes** | Yes | Yes | Yes |
| omitzero | **Yes** | Yes | No | No |
| Zero-is-null variant | **`zero/`** | `zero/` | No | No |
| json/v2 compatible | **Yes** | Yes | Yes | No |
| Legacy code | **None** | v1→v6 | N/A | N/A |

## Benchmarks

```
BenchmarkBoolMarshalJSON     0.85 ns/op    0 allocs
BenchmarkBoolUnmarshalJSON   2.1 ns/op     0 allocs
BenchmarkIntMarshalJSON      48 ns/op      2 allocs
BenchmarkIntUnmarshalJSON    193 ns/op     1 alloc
BenchmarkStringUnmarshalJSON 137 ns/op     0 allocs
BenchmarkStructMarshalJSON   876 ns/op     9 allocs
BenchmarkStructUnmarshalJSON 1116 ns/op    2 allocs
```

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## License

MIT License — see [LICENSE](LICENSE) for details.
