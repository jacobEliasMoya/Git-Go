# `git-go` / `ggo` — Codex Learning Handoff

## Status of This Document

This file is guidance and persistent project context for Codex sessions working in the `git-go` repository.

**Reading this document must not trigger implementation.** Codex should first inspect the repository, explain what currently exists, identify the learner's current checkpoint, and wait for a direct request before editing code or advancing to a major new checkpoint.

This document describes the intended destination and learning path. It is not evidence that any feature has already been implemented. The repository's actual files and behavior are the source of truth for current progress.

## Project Identity

- Working project/repository name: `git-go`
- Intended final executable name: `ggo`
- Language: Go
- Project type: beginner command-line utility
- Primary purpose: learn Go from zero by building a useful wrapper around a small Git workflow
- Secondary purpose: create a practical local tool that reinforces how compiled CLI programs interact with external commands

## Learner Context

The learner is new to Go, but is not new to programming. They already understand JavaScript, TypeScript, React, Next.js, APIs, frontend development, project structure, and production debugging.

Codex may use JavaScript or TypeScript comparisons to introduce unfamiliar Go concepts. Comparisons should help build intuition without implying that the languages behave identically.

Useful comparisons include:

- A Go function can be compared with a JavaScript/TypeScript function, while calling out Go's explicit parameter and return types.
- A Go package is loosely comparable to a module boundary, but Go package and import rules are distinct.
- `os.Args` is comparable to reading Node's `process.argv`.
- Go flag parsing can be compared with parsing CLI options in Node, while explaining the standard library's conventions.
- `exec.Command` is comparable to starting a child process from Node's `child_process`, but arguments are passed separately and execution errors must be handled explicitly.
- A compiled Go binary differs from running a JavaScript file through a runtime such as Node.

Do not assume knowledge of Go syntax, packages, modules, exported identifiers, error handling, slices, standard-library conventions, compilation, or executable installation.

## Project Purpose

The project should teach Go through a concrete utility rather than through isolated syntax exercises. The completed utility should make a familiar Git workflow shorter while giving the learner experience with:

- the structure of a minimal Go program;
- packages, imports, and `func main()`;
- CLI flags and positional arguments;
- strings and simple transformations;
- variables, constants, functions, parameters, and return values;
- slices, especially command argument slices;
- running external programs with `os/exec`;
- connecting a child process to the current terminal;
- checking and returning errors;
- early returns and readable control flow;
- building and running a Go binary;
- eventually organizing code once the single-file version is understood.

The goal is not merely to produce a working wrapper as quickly as possible. The learner should understand why the code works and be able to explain the main pieces.

## Intended User Experience

The final executable is `ggo`.

An eventual example command is:

```text
ggo -F "Testing push ability"
```

Its intended Git sequence is:

```text
git add .
git commit -m "(fix) Testing push ability."
git push
```

For this example:

- `ggo` is the compiled executable.
- `-F` represents the commit category `fix`.
- `Testing push ability` is the user-provided commit message.
- The generated commit message is `(fix) Testing push ability.`
- The utility adds a final period when needed according to the project's eventual message-formatting rules.
- Commands run sequentially.
- A later command must not run if an earlier required command fails.
- The underlying Git output should normally remain visible so the learner can understand what occurred.

The exact set and spelling of additional flags must be learned and decided incrementally. Do not invent a large flag system merely because the final tool may eventually support more commit categories.

## Behavioral Principles for the Finished Utility

Unless the learner deliberately changes the design later, the finished utility should follow these principles:

1. Accept a commit category flag and commit message from the command line.
2. Validate required input before changing the Git repository.
3. Convert the selected category and message into the project's commit-message format.
4. Run `git add .`.
5. Run `git commit -m <formatted-message>`.
6. Run `git push`.
7. Stop immediately and report a useful error if a required step fails.
8. Preserve normal Git output by wiring the child process to the terminal where appropriate.
9. Keep behavior understandable from the code; avoid hidden automation and excessive abstraction.

These are target behaviors, not permission to implement them all in one session.

## Teaching Method

Work slowly and in checkpoints. At each checkpoint, Codex should provide:

1. **What is being built now** — one small, concrete outcome.
2. **Why it matters** — how it contributes to the CLI and to broader backend/tooling skills.
3. **What Go concept it teaches** — introduce only the concepts needed at that point.
4. **The exact code for the current checkpoint** — preferably a small edit rather than a full-project replacement.
5. **How to run or test it** — include the expected observable behavior.
6. **What the learner should understand before continuing** — a short comprehension checkpoint.

After completing or explaining a major checkpoint, wait for confirmation such as `done`, `worked`, `got it`, `next`, or `continue` before moving to the next major checkpoint.

Do not silently advance through several checkpoints. If a tiny supporting correction is needed within the active checkpoint, it may be handled without treating it as a separate phase.

## Expected Progression

The repository may already be partway through this sequence. Codex must inspect the code before deciding where to begin, and must not recreate checkpoints already understood and working.

### Checkpoint 0 — Inspect and Re-establish Current State

- Read the repository's Go files, `go.mod`, README, learning notes, and relevant local guidance.
- Run safe, read-only checks where useful.
- Summarize what the program currently does.
- Identify the last completed concept and the smallest sensible next step.
- Note incomplete, broken, or inconsistent behavior without rewriting it immediately.

Concepts: reading an existing codebase, distinguishing current state from intended state.

### Checkpoint 1 — Minimal Go CLI Program

- Confirm or create the smallest runnable `package main` program only when requested.
- Explain `package main`, imports, and `func main()`.
- Run it with `go run`.

Concepts: program entry point, source file, compilation and execution.

### Checkpoint 2 — Read Command-Line Input

- Inspect raw command-line arguments or introduce a single basic flag.
- Show what input Go receives before adding complex behavior.

Concepts: `os.Args` and/or the standard `flag` package, strings, slices, indexes, validation.

### Checkpoint 3 — Define the First Commit-Type Flag

- Introduce the smallest useful interpretation of `-F`.
- Keep flag semantics explicit and easy to trace.
- Validate that the commit message exists.

Concepts: booleans or string flags, pointers returned by `flag` helpers if applicable, conditionals, early returns.

### Checkpoint 4 — Format the Commit Message

- Produce `(fix) Testing push ability.` from the example input.
- Add punctuation carefully and avoid duplicating a final period.
- Keep formatting in `main` until extracting it teaches a clear lesson.

Concepts: string concatenation or formatting, `strings`, helper functions, parameters and return values.

### Checkpoint 5 — Run One Harmless or Deliberately Scoped External Command

- Introduce `os/exec` with one command at a time.
- Explain the program/argument separation in `exec.Command`.
- Connect output so the command's behavior is visible.

Concepts: child processes, variadic arguments, stdout/stderr, returned errors.

Because Git commands can modify a repository or contact a remote, Codex must describe the exact command and likely effect before running it during a teaching session. Testing should use an appropriate test repository or occur only with the learner's clear approval.

### Checkpoint 6 — Create a Small Command Runner

- Extract repeated external-command setup only after the repeated code is visible and understood.
- Keep the helper narrow, such as running Git with a slice of arguments.

Concepts: reusable functions, variadic parameters or slices, error propagation.

### Checkpoint 7 — Add `git add .`

- Run and verify the staging step.
- Stop on failure.

Concepts: sequencing, side effects, error handling.

### Checkpoint 8 — Add `git commit`

- Pass the formatted message as one argument to `-m`.
- Explain why argument separation avoids manual shell quoting.
- Stop before pushing if the commit fails.

Concepts: command arguments, error propagation, control flow.

### Checkpoint 9 — Add `git push`

- Add pushing only after staging and committing are understood and tested.
- Preserve Git authentication and remote output rather than attempting to reimplement them.

Concepts: external process behavior, exit status, local versus remote side effects.

### Checkpoint 10 — Improve Input and Error Messages

- Handle missing flags, missing messages, conflicting options, and command failures.
- Keep messages concise and actionable.
- Consider exit codes only when the basic flow is already understood.

Concepts: errors as values, wrapping errors, user-facing CLI failures, `os.Exit` placement.

### Checkpoint 11 — Build and Use the `ggo` Binary

- Build the executable with the name `ggo` (or `ggo.exe` on Windows).
- Explain the difference between `go run`, `go build`, and invoking an installed binary.
- Explain PATH installation only when the learner is ready.

Concepts: compilation, binaries, operating-system conventions, PATH.

### Checkpoint 12 — Tests and Focused Refactoring

- Add tests first for pure behavior such as commit-message formatting.
- Avoid executing real Git commands in unit tests until dependency boundaries are taught deliberately.
- Refactor only when the existing single-file version is working and the reason for each boundary is clear.

Concepts: table-driven tests, pure functions, package boundaries, dependency seams.

### Later Possibilities — Not Current Requirements

Only consider these after the core utility is understood:

- more commit-type flags;
- help and usage improvements;
- a dry-run mode;
- selecting files instead of always staging `.`;
- configuration;
- cross-platform installation scripts;
- richer tests around command execution;
- splitting the program into packages.

Do not treat this list as a backlog that should be implemented automatically.

## Codex Operating Instructions

At the beginning of a new work session:

1. Read this file and any repository-level `AGENTS.md`.
2. Inspect the repository before proposing edits. At minimum, look for `go.mod`, all `.go` files, README/learning documents, tests, and current Git status.
3. Treat existing code as the learner's work. Preserve it where reasonable.
4. State what is currently implemented, what appears unfinished, and what the next smallest checkpoint should be.
5. Do not infer completion from this handoff document.
6. Do not implement anything merely because this document was opened.

When asked to make a change:

- Work only on the requested or mutually established checkpoint.
- Prefer a focused patch over replacing an entire file.
- Explain new syntax before or alongside its first use.
- Point out compiler errors in plain language and show how the error relates to the current concept.
- Format changed Go files with `gofmt`.
- Run proportionate checks such as `go test ./...`, `go vet ./...`, `go run`, or `go build` when they match the current checkpoint.
- Do not execute `git add`, `git commit`, or `git push` merely to test the utility against the learner's real repository without explicit approval and a clear explanation of the effect.
- Report what was changed and how it was verified.
- End major checkpoints with a short understanding check, then wait.

## Command-Teaching Protocol

Terminal commands are part of the learning material, not invisible implementation details. This applies especially to Go and Git commands.

Before running a command during a teaching session, Codex should explain, in proportion to its complexity and risk:

- the working directory in which it will run;
- the executable and important arguments or flags;
- why the command is relevant to the current checkpoint;
- the expected output or filesystem, build, Git, or remote effect;
- whether the command is read-only, creates generated files, changes the working tree or index, creates history, or affects a remote.

After running it, Codex should summarize what the result means, call out warnings or errors, and show the learner how to reproduce the useful check. Repeated commands may receive a shorter reminder once the learner has demonstrated understanding, but new flags and materially different effects must still be explained.

Commands that teach core workflow—such as `go run`, `go test`, `go test -race`, `go vet`, `gofmt`, `go build`, `git status`, `git diff`, `git add`, `git commit`, and `git push`—should be unpacked rather than presented as unexplained recipes.

The existing safety requirement remains: do not run repository-changing or remote-changing Git commands without the learner's authorization and a clear statement of the target and effect.

## Documentation-Reference Protocol

Technical teaching and code review should include direct links to the most relevant primary documentation. Do not give only a documentation homepage when a specific language feature, standard-library package, Go command, Git command, or PowerShell cmdlet has its own reference page.

Prefer these sources:

- `https://go.dev` for the Go specification, official tutorials, toolchain documentation, release notes, and diagnostics;
- `https://pkg.go.dev` for exact standard-library package, type, function, and method references;
- `https://git-scm.com/docs` for exact Git command and concept references;
- `https://learn.microsoft.com/powershell` for PowerShell syntax and cmdlet behavior in this Windows environment.

Link the primary page that directly supports the explanation. Add closely related official references when they help connect the current concept to testing, tooling, security, performance, or a prerequisite, but avoid burying the learner in an undifferentiated link dump. Briefly state what each linked reference is useful for.

## Cross-Language Transfer Protocol

Use the learner's JavaScript, TypeScript, React, Next.js, and API experience to build accurate mental models for Go. When introducing a concept, explain the nearest familiar idea when useful, then state the important differences and where the analogy stops working.

Useful comparisons include modules versus packages, arrays versus slices, objects versus structs, TypeScript structural typing versus Go interfaces, exceptions or rejected promises versus explicit Go errors, Node child processes versus `os/exec`, async JavaScript versus goroutines, and request cancellation versus Go `context`.

Do not force a comparison when Go's model is clearer on its own. Correct transferred habits that produce non-idiomatic Go, such as premature abstraction, class-shaped designs, overly broad interfaces, hidden error handling, or assuming JavaScript reference behavior. Also connect concepts to SQL, HTTP, operating systems, testing, and distributed systems when those relationships strengthen practical engineering judgment.

Type-system comparisons deserve special attention. Contrast TypeScript's mostly erased compile-time types with Go's statically compiled types and concrete runtime representations. Revisit inference, zero values, assignability, conversions, structs, implicit interface satisfaction, pointers and `nil`, `any`, generics, the absence of ordinary union types, and value versus reference-like behavior as those topics arise. Use small prediction exercises to expose differences rather than relying only on verbal analogies.

If existing code takes a different reasonable approach, explain the tradeoff before changing it. Preserve working choices unless they block the learning objective, introduce incorrect behavior, or the learner asks for a rewrite.

## Code-Comment Philosophy

Comments should help the learner understand **why** code works, especially where Go behavior is unfamiliar.

Good comments may explain:

- why `package main` and `func main()` are required for an executable;
- why `exec.Command("git", "commit", "-m", message)` passes arguments separately;
- why child-process stdout and stderr are connected to the current terminal;
- why the function returns immediately after an error;
- why a message-formatting helper returns a new string;
- why a command runner returns an error instead of exiting deep inside the helper.

Avoid comments that merely translate syntax into English:

```go
// Set name to Jacob.
name := "Jacob"
```

Prefer a concise explanation of a non-obvious decision:

```go
// Pass each Git argument separately so Go handles spaces in the commit message
// without us building and quoting a shell command string.
cmd := exec.Command("git", "commit", "-m", message)
```

Comments should be plentiful enough to support a beginner, but not so numerous that every obvious line is narrated. If a concept deserves a long explanation, explain it in the teaching response or learning notes and keep the source comment concise.

Do not erase the learner's useful comments merely to impose a different style. Correct comments that are inaccurate or misleading, and explain why.

## Scope Boundaries

For the core learning version, prefer:

- one `main.go` file until there is a concrete teaching reason to split it;
- the Go standard library;
- direct, readable control flow;
- small functions introduced after their purpose is visible;
- explicit error checking;
- behavior that can be exercised from the terminal.

Avoid introducing the following prematurely:

- third-party CLI frameworks such as Cobra;
- dependency-injection frameworks;
- elaborate command, service, or repository layers;
- configuration frameworks;
- logging frameworks;
- concurrency or goroutines;
- platform-specific shell scripts as the main implementation;
- automatic semantic-versioning or release pipelines;
- broad abstractions designed for hypothetical future features;
- a full rewrite when a small edit would preserve the learner's progress.

Do not turn the project into a production-grade Git client. The utility may delegate authentication, repository discovery, hooks, and remote behavior to the installed `git` executable.

## Safety and Git Side Effects

This project intentionally learns operations that can change local and remote repositories. Distinguish code authoring from actually invoking those operations.

- Inspecting files, building, formatting, and running unit tests are normally safe.
- `git add .` changes the staging area.
- `git commit` creates repository history and may invoke hooks.
- `git push` contacts a remote and changes shared state.

Before end-to-end execution, tell the learner exactly which repository will be affected and which commands will run. Prefer a disposable test repository for early integration testing. Never hide or bypass authentication, hooks, branch protections, or Git errors.

## Definition of a Successful Learning Checkpoint

A checkpoint is complete when:

- its narrow behavior works;
- the relevant code is formatted;
- an appropriate check has passed;
- the learner has been shown how to reproduce the result;
- the new Go concept has been explained in beginner-friendly language;
- the learner can reasonably describe what the important lines do;
- Codex has not advanced into the next major checkpoint without confirmation.

## Definition of the Core Project's Eventual Success

The core project is successful when the learner understands and can maintain a simple Go CLI that:

- builds to an executable named `ggo`;
- accepts the established flag/message input;
- formats the intended commit message;
- runs `git add .`, `git commit -m <message>`, and `git push` in order;
- exposes Git output;
- stops and reports errors correctly;
- has focused tests for pure logic where appropriate;
- remains simple enough for the learner to explain without relying on a framework.

The learning outcome is more important than minimizing the number of sessions or lines of code.

## Session-Start Reminder for Codex

Use this sequence whenever resuming work:

> Inspect first. Describe the repository's actual state. Identify the smallest next learning checkpoint. Ask or wait for authorization to implement it. Teach only that checkpoint, verify it, and wait for confirmation before continuing.
