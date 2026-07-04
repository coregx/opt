# Contributing to opt

Thank you for your interest in contributing to coregx/opt!

## Prerequisites

- **Go 1.24+** ([download](https://go.dev/dl/))

## Building & Testing

```bash
go build ./...                  # build all packages
go test ./...                   # run all tests
go test -v -run TestName        # run a single test
go test -cover ./...            # check coverage
go test -bench=. -benchmem ./.. # run benchmarks
```

## Code Style

- Run `go fmt ./...` before every commit
- Follow standard Go naming conventions
- Exported types and functions must have doc comments
- No external dependencies — stdlib only

## Pull Request Workflow

1. Fork and create a feature branch:
   ```bash
   git checkout -b feat/my-feature
   ```
2. Make changes, add tests
3. Verify:
   ```bash
   go fmt ./...
   go build ./...
   go test ./...
   ```
4. Push and open a pull request against `main`

Commit messages follow [Conventional Commits](https://www.conventionalcommits.org/):
```
feat: add Uint type
fix: Byte overflow on values > 255
docs: update PATCH API examples
```

## Architecture

```
opt/
├── value.go       Value[T] — generic foundation
├── field.go       Field[T] — three-state for PATCH
├── funcs.go       Map, FlatMap, Equal
├── string.go      String (+ other concrete types)
├── internal/      Shared marshal/unmarshal helpers
└── zero/          Alternative semantics (zero = null)
```

## Adding a New Type

1. Create `mytype.go` in `opt/` with the type struct embedding `Value[T]`
2. Add constructors: `NewMyType`, `MyTypeFrom`, `MyTypeFromPtr`
3. Add `Equal`, `MarshalJSON`, `UnmarshalJSON`, `MarshalText`, `UnmarshalText`
4. Create `mytype_test.go` with marshal/unmarshal/roundtrip/equal tests
5. Mirror in `zero/` with zero-is-null semantics
6. Add shared unmarshal logic to `internal/json.go` if complex

## License

By contributing, you agree that your contributions will be licensed under the [MIT License](LICENSE).
