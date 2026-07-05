# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.3.0] - 2026-07-05

### Breaking Changes

- **`Value[T]` renamed to `Option[T]`** — fixes `driver.Valuer` shadow bug where all concrete types (String, Int, etc.) failed to implement `driver.Valuer` due to Go selector rules. See ADR-003 in project docs
- **`Field.ToValue()` renamed to `Field.ToOption()`** — consistent with the type rename

### Added

- **`FieldFromOption[T]()`** — lossless conversion from `Option[T]` to `Field[T]` (valid→present+valid, invalid→present+null). See ADR-002 in project docs
- **Compile-time interface checks** — `interfaces_test.go` in both `opt/` and `zero/` packages. Verifies `driver.Valuer`, `sql.Scanner`, `json.Marshaler`, `encoding.TextMarshaler` on all 16 concrete types. Prevents future shadow bugs

### Fixed

- All concrete types now correctly implement `driver.Valuer` via method promotion
- pgx v5 INSERT/Exec works without explicit forwarding methods
- `doc.go` corrected: `driver.Valuer` claim is now accurate

### Migration from v0.2.0

```
Find and replace in your code:
  opt.Value[T]        →  opt.Option[T]     (generic usage, rare)
  .Value field access  →  .Option           (embedded field, rare)
  .ToValue()          →  .ToOption()        (Field method)
```

Concrete types (opt.String, opt.Int, etc.) and constructors (New, From, FromPtr, OrNull) are **unchanged**.

## [0.2.0] - 2026-07-04

### Added

- **`OrNull[T]` generic constructor** — zero value → null mapping for any comparable type
- **Typed OrNull constructors** — `StringOrNull`, `IntOrNull`, `Int32OrNull`, `Int16OrNull`, `FloatOrNull`, `BoolOrNull`, `ByteOrNull`, `TimeOrNull`
- Eliminates boilerplate helpers in DB/API projects: `opt.StringOrNull(s)` instead of custom `optStr(s)` functions

### Documentation

- Updated README with OrNull API section
- Updated AGENTS.md with OrNull examples
- Added OrNull example tests for godoc

## [0.1.0] - 2026-07-04

### Added

- **Generic `Value[T]`** (renamed to `Option[T]` in v0.3.0) — foundation type on `sql.Null[T]` with full JSON/SQL support
- **9 concrete types** — String, Int, Int32, Int16, Float, Bool, Byte, Time with optimized marshal/unmarshal
- **`Field[T]`** — three-state (absent/null/value) for PATCH API semantics
- **Functional API** — `Map`, `FlatMap`, `Equal` top-level generic functions
- **`zero/` subpackage** — alternative semantics where zero value = null (9 types)
- **`internal/`** — shared unmarshal helpers for DRY across opt and zero
- Benchmarks with zero-allocation unmarshal
- Go 1.24+ `omitzero` support via `IsZero()` on all types
- `encoding/json/v2` compatible (no changes needed)
- CI/CD: GitHub Actions (3 OS × 3 Go versions), Codecov OIDC, branch protection

[Unreleased]: https://github.com/coregx/opt/compare/v0.3.0...HEAD
[0.3.0]: https://github.com/coregx/opt/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/coregx/opt/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/coregx/opt/releases/tag/v0.1.0
