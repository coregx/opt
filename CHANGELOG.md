# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

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

- **Generic `Value[T]`** — foundation type on `sql.Null[T]` with full JSON/SQL support
- **9 concrete types** — String, Int, Int32, Int16, Float, Bool, Byte, Time with optimized marshal/unmarshal
- **`Field[T]`** — three-state (absent/null/value) for PATCH API semantics
- **Functional API** — `Map`, `FlatMap`, `Equal` top-level generic functions
- **`zero/` subpackage** — alternative semantics where zero value = null (9 types)
- **`internal/`** — shared unmarshal helpers for DRY across opt and zero
- Benchmarks with zero-allocation unmarshal
- Go 1.24+ `omitzero` support via `IsZero()` on all types
- `encoding/json/v2` compatible (no changes needed)
- CI/CD: GitHub Actions (3 OS × 3 Go versions), Codecov OIDC, branch protection

[Unreleased]: https://github.com/coregx/opt/compare/v0.2.0...HEAD
[0.2.0]: https://github.com/coregx/opt/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/coregx/opt/releases/tag/v0.1.0
