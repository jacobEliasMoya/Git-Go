# `git-go` Phase Notes

## Stage 1 — Checkpoint 1: Minimal executable

Date: 2026-08-26

Status: Implementation working; learner explanation pending.

### Outcome

`main.go` defines a minimal Go executable that prints:

```text
Hello There
```

### Concepts encountered

- `package main` identifies a package intended to build as an executable.
- `import "fmt"` makes Go's standard formatting package available to this file.
- `func main()` is the entry point Go invokes when the executable starts.
- `fmt.Println` calls an exported standard-library function and writes a line to standard output.
- Go requires imported packages to be used.
- `gofmt` provides one canonical source format rather than treating formatting as a project-specific style choice.

### TypeScript and JavaScript connections

- A Go import and an ECMAScript import both make another module or package's identifiers available, but Go imports are checked by the compiler and unused imports are rejected.
- `func main()` has no exact Node.js equivalent. Node normally starts evaluating the chosen JavaScript module at its top level; a Go executable starts through the specially recognized `main` function in `package main`.
- TypeScript is normally transformed into JavaScript for a JavaScript runtime. Go compiles statically typed source into native executable code before that program runs.

### Commands introduced

Run from `C:\Users\Jemoy\Golang\git-go` unless otherwise stated.

#### Inspect the installed toolchain

```powershell
go version
```

This invokes the Go toolchain and reports its installed version and target operating system/architecture.

#### Locate the active module

```powershell
go env GOMOD
```

This reads Go environment information and prints the `go.mod` file governing the current directory.

#### Apply canonical formatting

```powershell
gofmt -w main.go
```

`-w` writes the formatted source back to `main.go`.

#### Compile and run the current package

```powershell
go run .
```

`run` compiles and executes a main package; `.` selects the package in the current directory. The resulting executable is temporary rather than a permanent `ggo.exe` build artifact.

### Verification

- `go version`: Go 1.26.4 for Windows/AMD64 was found.
- `go run main.go`: printed `Hello There` and exited with status `0`.
- Final package-level validation is recorded before the checkpoint commit.

### Understanding check

Before marking this checkpoint complete, explain:

1. Why must this executable use `package main`?
2. What role does `func main()` play?
3. Why might Go reject unused imports?
4. How does `go run .` differ from running a TypeScript or JavaScript file?

### Primary references

- [Getting started with Go](https://go.dev/doc/tutorial/getting-started)
- [`fmt` package](https://pkg.go.dev/fmt)
- [`go run` command](https://pkg.go.dev/cmd/go#hdr-Compile_and_run_Go_program)
- [`gofmt` command](https://pkg.go.dev/cmd/gofmt)
- [Go modules reference](https://go.dev/ref/mod)
