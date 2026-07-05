# AGENTS.md — opt

> Optional types for Go. Generic-first, SQL+JSON, three-state PATCH support.

## What is opt

`opt` provides nullable types for Go that properly handle JSON serialization (unlike `sql.Null[T]` which marshals as `{"V":...,"Valid":...}`) and SQL scanning. Inspired by Rust's `Option<T>`, Kotlin's `T?`, and C#'s `Nullable<T>`.

## Quick Start

```go
import "github.com/coregx/opt"

name := opt.StringFrom("Alice")       // always valid
city := opt.StringOrNull(userCity)     // "" → null, "Moscow" → valid
age := opt.NewInt(0, false)            // explicit null

fmt.Println(name.Or("unknown")) // "Alice"
fmt.Println(city.OrZero())      // "" if null
fmt.Println(age.OrZero())       // 0

data, _ := json.Marshal(struct {
    Name opt.String `json:"name"`
    City opt.String `json:"city"`
    Age  opt.Int    `json:"age"`
}{name, city, age})
// {"name":"Alice","city":null,"age":null}
```

## Build & Test

```bash
go build ./...
go test ./...
go test -bench=. -benchmem ./...
```

## Architecture

```
opt/
├── option.go      Option[T] — generic foundation (sql.Null[T] + JSON)
├── field.go       Field[T] — three-state (absent/null/value) for PATCH
├── funcs.go       Map, FlatMap, Equal, OrNull, FieldFromOption
├── string.go      String, Int, Int32, Int16, Float, Bool, Byte, Time
├── internal/      Shared unmarshal helpers (DRY across opt + zero)
└── zero/          Alternative semantics (zero value = null)
```

## Links

- GitHub: https://github.com/coregx/opt
- Docs: https://pkg.go.dev/github.com/coregx/opt
