# Roadmap

> **Optional types for Go — generic-first SQL+JSON nullable types**

---

## Current State: v0.3.1

### Core
- **`Option[T]`** — generic foundation on `sql.Null[T]` with JSON/SQL/Text support
- **9 concrete types** — String, Int, Int32, Int16, Float, Bool, Byte, Time
- **`Field[T]`** — three-state (absent/null/value) for PATCH API
- **`OrNull[T]` constructors** — zero value → null mapping (eliminates boilerplate)
- **`FieldFromOption[T]`** — lossless Option→Field conversion for PATCH workflows
- **Functional API** — Map, FlatMap, Equal
- **`zero/` subpackage** — alternative semantics (zero value = null)
- **`internal/`** — shared unmarshal helpers (DRY)
- **Compile-time interface checks** — driver.Valuer, sql.Scanner, json.Marshaler on all types

### Quality
- 82% coverage (opt), 62% coverage (zero)
- Zero-allocation unmarshal, Bool marshal <1ns
- CI: 3 OS × 3 Go versions, Codecov OIDC
- json/v2 compatible
- driver.Valuer verified on all 16 concrete types

---

## Upcoming

### v0.4.0 — Polish & Coverage

- [ ] Coverage 90%+ on both opt and zero
- [ ] Edge case tests: concurrent access, large values, unicode
- [ ] `Uint` / `Uint64` types
- [ ] `JSON` type (raw JSON storage, `json.RawMessage` wrapper)
- [ ] `Bytes` type (`[]byte` for JSONB/binary columns)
- [ ] `Field.ApplyTo(current Option[T]) Option[T]` — PATCH merge

### v0.5.0 — Integration

- [ ] `OrElse` on concrete types (not just Option[T])
- [ ] `Filter(func(T) bool) Option[T]` — keep value only if predicate matches
- [ ] pgx native type support (beyond sql.Scanner)
- [ ] sqlx `NamedExec` / `StructScan` validation
- [ ] Example: REST API with PATCH using Field[T]

### v0.6.0 — json/v2 Optimization

- [ ] `MarshalerTo` / `UnmarshalerFrom` stream-based interfaces
- [ ] Build tag: `//go:build goexperiment.jsonv2`
- [ ] Benchmark comparison: json/v1 vs json/v2
- [ ] Zero-alloc marshal (eliminate remaining allocations)

### v1.0.0 — Stable API

- [ ] API freeze
- [ ] awesome-go submission
- [ ] Migration guide from guregu/null
- [ ] Migration guide from `*T` patterns
- [ ] 95%+ coverage
- [ ] Security audit

---

## Architecture

```
opt/
├── option.go      Option[T] — generic foundation
├── field.go       Field[T] — three-state for PATCH
├── funcs.go       Map, FlatMap, Equal, OrNull, FieldFromOption
├── string.go      Concrete types (String, Int, Float, Bool, Time, ...)
├── internal/      Shared unmarshal helpers
└── zero/          Alternative semantics (zero = null)
```

---

## Design Principles

1. **Generic-first** — `Option[T]` is the foundation, concrete types are ergonomic wrappers
2. **No legacy** — clean API from scratch, no backward compatibility burden
3. **Stdlib only** — zero external dependencies
4. **Enterprise quality** — benchmarks, high coverage, CI on all platforms, compile-time interface checks
5. **Inspired by the best** — Rust `Option<T>`, Kotlin `T?`, C# `Nullable<T>`
