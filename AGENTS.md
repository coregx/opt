# AGENTS.md — opt

> Optional types for Go. Generic-first, SQL+JSON, three-state PATCH support.

## What is opt

`opt` provides nullable types for Go that properly handle JSON serialization (unlike `sql.Null[T]` which marshals as `{"V":...,"Valid":...}`) and SQL scanning. Inspired by Rust's `Option<T>`, Kotlin's `T?`, and C#'s `Nullable<T>`.

## Quick Start

```go
import "github.com/coregx/opt"

name := opt.StringFrom("Alice")
age := opt.NewInt(0, false) // null

fmt.Println(name.Or("unknown")) // "Alice"
fmt.Println(age.OrZero())       // 0

data, _ := json.Marshal(struct {
    Name opt.String `json:"name"`
    Age  opt.Int    `json:"age"`
}{name, age})
// {"name":"Alice","age":null}
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
├── value.go       Value[T] — generic foundation (sql.Null[T] + JSON)
├── field.go       Field[T] — three-state (absent/null/value) for PATCH
├── funcs.go       Map, FlatMap, Equal
├── string.go      String, Int, Int32, Int16, Float, Bool, Byte, Time
├── internal/      Shared unmarshal helpers (DRY across opt + zero)
└── zero/          Alternative semantics (zero value = null)
```

## Links

- GitHub: https://github.com/coregx/opt
- Docs: https://pkg.go.dev/github.com/coregx/opt
