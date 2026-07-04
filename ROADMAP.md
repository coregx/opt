# Roadmap

> **Optional types for Go — generic-first SQL+JSON nullable types**

---

## Current State: v0.1.0

### Core
- **`Value[T]`** — generic foundation on `sql.Null[T]` with JSON/SQL/Text support
- **9 concrete types** — String, Int, Int32, Int16, Float, Bool, Byte, Time
- **`Field[T]`** — three-state (absent/null/value) for PATCH API
- **Functional API** — Map, FlatMap, Equal
- **`zero/` subpackage** — alternative semantics (zero value = null)
- **`internal/`** — shared unmarshal helpers (DRY)

### Quality
- 82% coverage (opt), 62% coverage (zero)
- Zero-allocation unmarshal, Bool marshal <1ns
- CI: 3 OS × 3 Go versions, Codecov OIDC
- json/v2 compatible

---

## Upcoming

### v0.2.0 — Polish & Coverage

- [ ] Coverage 90%+ on both opt and zero
- [ ] Edge case tests: concurrent access, large values, unicode
- [ ] `Uint` / `Uint64` types
- [ ] `JSON` type (raw JSON storage, `json.RawMessage` wrapper)
- [ ] `Bytes` type (`[]byte` for JSONB/binary columns)

### v0.3.0 — Functional & Integration

- [ ] `OrElse` on concrete types (not just Value[T])
- [ ] `Filter(func(T) bool) Value[T]` — keep value only if predicate matches
- [ ] pgx native type support (beyond sql.Scanner)
- [ ] sqlx `NamedExec` / `StructScan` validation
- [ ] Example: REST API with PATCH using Field[T]

### v0.4.0 — json/v2 Optimization

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
├── value.go       Value[T] — generic foundation
├── field.go       Field[T] — three-state for PATCH
├── funcs.go       Map, FlatMap, Equal
├── string.go      Concrete types (String, Int, Float, Bool, Time, ...)
├── internal/      Shared unmarshal helpers
└── zero/          Alternative semantics (zero = null)
```

---

## Design Principles

1. **Generic-first** — `Value[T]` is the foundation, concrete types are ergonomic wrappers
2. **No legacy** — clean API from scratch, no backward compatibility burden
3. **Stdlib only** — zero external dependencies
4. **Enterprise quality** — benchmarks, high coverage, CI on all platforms
5. **Inspired by the best** — Rust `Option<T>`, Kotlin `T?`, C# `Nullable<T>`
