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
4. What compilation and execution steps does `go run .` perform?

### Primary references

- [Getting started with Go](https://go.dev/doc/tutorial/getting-started)
- [`fmt` package](https://pkg.go.dev/fmt)
- [`go run` command](https://pkg.go.dev/cmd/go#hdr-Compile_and_run_Go_program)
- [`gofmt` command](https://pkg.go.dev/cmd/gofmt)
- [Go modules reference](https://go.dev/ref/mod)

## 2026-08-27 — World 2.1: Observe command-line arguments

Status: Argument experiment works with two user arguments; bounds validation remains intentionally unfinished.

### What I changed

- Imported `os` to read `os.Args`.
- Imported `path/filepath` to display only the executable's base name.
- Created `args := os.Args[1:]` so the new slice contains user-provided arguments and excludes the executable entry.
- Printed the first two user arguments by index.

### What the code currently assumes

The program assumes at least two user arguments exist. Accessing `args[0]` or `args[1]` without enough elements causes an index-out-of-range panic. This is recorded as the next validation problem rather than hidden by an instructor-written fix.

### Executable-path safety note

`os.Args[0]` is informational and should not be treated as trusted identity. Under `go run`, it normally refers to a temporary compiled executable. Printing that path does not execute it again. `filepath.Base` shortens the displayed value but is presentation cleanup, not a security boundary.

### Current valid invocation

```powershell
go run . alpha "hello world"
```

Expected user-argument slice:

- `args[0]` is `alpha`.
- `args[1]` is `hello world`.

### Next problem

Before treating the program as a usable CLI, validate `len(args)` before indexing and define behavior for zero, one, two, and extra user arguments.

### Primary references

- [`os.Args`](https://pkg.go.dev/os#Args)
- [`filepath.Base`](https://pkg.go.dev/path/filepath#Base)
- [Go specification: index expressions](https://go.dev/ref/spec#Index_expressions)
- [`len` built-in](https://pkg.go.dev/builtin#len)
