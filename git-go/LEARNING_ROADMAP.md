# Employable Go Learning Roadmap

## Purpose

This roadmap uses small, finished projects to build deep Go knowledge and practical backend engineering skills. Completing code is not enough: each checkpoint requires explanation, testing, and later recall without copying.

Codex should review the learner's attempt before offering a complete implementation. Hints should progress from questions, to conceptual guidance, to pseudocode, and only then to focused code when requested or necessary.

## Completion standard for every checkpoint

A checkpoint is complete only when the learner can:

- run the behavior and describe the observed result;
- explain the important syntax and control flow in their own words;
- identify at least one relevant failure or edge case;
- write or understand the tests appropriate to the checkpoint;
- run formatting, tests, and static checks appropriate to the current stage;
- make a focused commit that records the milestone.

Useful checks will eventually include `gofmt`, `go test`, `go test -race`, `go vet`, fuzz tests, benchmarks, and integration tests. They should be introduced when their purpose is concrete rather than all at once.

## Command literacy

Terminal fluency is an explicit learning goal. Whenever a new Go, Git, or shell command is introduced, record and understand:

- which directory it runs from;
- what program is invoked;
- what each meaningful argument or flag changes;
- what output, files, Git state, process, or remote state it may affect;
- what success and common failure look like;
- how to inspect or reverse its effect when reversal is possible.

The learner should eventually be able to type the common development commands from memory and explain them rather than relying on copied command sequences.

## Primary documentation

Use exact pages from these primary sources throughout the roadmap:

- [Go documentation](https://go.dev/doc/) for official language and toolchain guides;
- [How to Write Go Code](https://go.dev/doc/code) for modules, packages, builds, and the initial testing model;
- [Go command reference](https://pkg.go.dev/cmd/go) for exact `go` subcommands and flags;
- [Go standard library](https://pkg.go.dev/std) for package and API references;
- [Git reference](https://git-scm.com/docs) for Git concepts and commands;
- [PowerShell documentation](https://learn.microsoft.com/powershell/) for the Windows shell used by this workspace.

Each checkpoint should link to the exact relevant page—for example, the specific package or Git command—not only this general list. Related official references should be included when they clarify a prerequisite or a meaningful connection without overwhelming the current lesson.

## Cross-language connections

Use prior JavaScript and TypeScript knowledge as a bridge. For each useful comparison:

1. identify the familiar concept;
2. explain the shared purpose;
3. state the Go-specific behavior;
4. identify where the analogy breaks;
5. test the difference with a small prediction or experiment when practical.

Pay particular attention to package boundaries, slices and backing arrays, value and pointer semantics, structural interfaces, explicit errors, process execution, goroutines, channels, and `context`. The objective is to transfer durable programming knowledge while learning to write idiomatic Go rather than transliterated TypeScript.

Maintain a recurring TypeScript-to-Go type map as concepts are encountered—not as an upfront vocabulary dump. It should cover inference, primitive types, arrays and slices, object shapes and structs, structural interfaces, `any`, generics, pointers, `nil`, zero values, conversions, and Go's lack of TypeScript-style union types. Every mapping must include the important mismatch, especially that TypeScript types are generally erased before JavaScript executes while Go uses static types to compile native code with concrete representations.

## Stage 1 — Build `ggo` from first principles

### Checkpoint 1: Minimal executable

Outcome: create the smallest Go program that runs and prints an intentional message.

Learn:

- source files, packages, imports, and the program entry point;
- the difference between `go run` and a compiled executable;
- how the compiler reports syntax and import errors.

Evidence:

- run the program;
- deliberately predict what each required line does;
- format it and confirm it still runs.

### Checkpoint 2: Observe command-line arguments

Outcome: inspect exactly what the operating system passes to the program.

Learn:

- strings, slices, indexing, length, and bounds safety;
- the executable path versus user-provided arguments;
- why quoted terminal input becomes one argument.

Evidence:

- try no arguments, one argument, and a quoted multi-word argument;
- explain every element received by the program.

### Checkpoint 3: Validate the first command shape

Outcome: accept only the initial `-F` workflow and reject incomplete or unknown input before any Git operation.

Learn:

- conditionals and early returns;
- validation order;
- user-facing errors versus programmer errors;
- raw argument parsing before introducing the `flag` package.

Evidence:

- exercise valid, missing, extra, and unknown input;
- show that invalid input causes no repository changes.

### Checkpoint 4: Format a commit message

Outcome: turn a validated message into the established fix-commit format without duplicating terminal punctuation.

Learn:

- typed function parameters and return values;
- pure functions;
- string operations and small design decisions;
- why pure behavior is the easiest place to begin testing.

Evidence:

- create table-driven unit tests covering ordinary and edge-case messages;
- explain the Arrange/Act/Assert structure even if the test is written compactly.

### Checkpoint 5: Run one external command safely

Outcome: use Go to start one harmless, visible child process before modifying Git state.

Learn:

- programs versus shell command strings;
- standard input, output, and error streams;
- returned errors and process exit status;
- why arguments should be passed separately.

Evidence:

- observe success and a deliberately failing command;
- explain what Go knows when the child process fails.

### Checkpoint 6: Introduce a Git command boundary

Outcome: isolate command execution so behavior can be tested without pushing to a real remote.

Learn:

- small helpers, slices or variadic arguments, and error propagation;
- dependency seams without a framework;
- when an interface helps and when it is premature.

Evidence:

- test argument construction separately from real Git execution;
- confirm errors stop later steps.

### Checkpoints 7–9: Stage, commit, then push

Outcomes, introduced one at a time:

1. run `git add .`;
2. run `git commit -m <formatted-message>`;
3. run `git push` only after the earlier steps succeed.

Learn:

- ordered side effects;
- wrapping errors with useful context;
- terminal wiring and authentication boundaries;
- why partial failure must be designed explicitly.

Evidence:

- use a disposable local repository for integration tests;
- test a failure at each stage and prove later commands do not run;
- never exercise push against a real remote merely for automated testing.

### Checkpoint 10: Produce and use the binary

Outcome: build `ggo`, invoke it as an executable, and understand how it is discovered through `PATH`.

Learn:

- builds, binaries, exit codes, and cross-platform naming;
- version information and reproducible build commands;
- generated artifacts versus tracked source.

## Stage 2 — Go language depth

Build small exercises or extend a suitable project to cover:

- arrays, slices, and maps, including allocation and mutation;
- structs, methods, pointer receivers, and value semantics;
- interfaces defined at the point of use;
- error creation, wrapping, inspection, and sentinel errors;
- `defer`, resource ownership, and cleanup;
- readers, writers, files, JSON, time, and environment configuration;
- package design, exported identifiers, documentation, and internal packages;
- generics only after ordinary functions and interfaces are comfortable.

Testing depth:

- table-driven tests and subtests;
- test helpers and temporary directories;
- golden files where output stability matters;
- fuzz tests for parsers and formatting invariants;
- benchmarks and allocation reports for measured questions.

Exit project: a second CLI that consumes files or structured input and has meaningful unit and integration tests.

## Stage 3 — Standard-library HTTP service

Build a service primarily with `net/http` before selecting a framework.

Learn:

- handlers, routing, middleware, and request lifecycles;
- JSON contracts, validation, status codes, and consistent errors;
- `context`, timeouts, cancellation, and graceful shutdown;
- configuration, structured logging, and dependency construction;
- unit tests with `httptest` and end-to-end HTTP integration tests.

Exit project: a documented API with tests for success, malformed input, authorization boundaries, timeouts, and shutdown behavior.

## Stage 4 — SQL and PostgreSQL

Learn:

- relational modeling, keys, constraints, and indexes;
- parameterized queries and scanning results;
- transactions, isolation, rollback, and idempotency;
- migrations and compatibility between application and schema versions;
- connection pools, cancellation, and database failure handling;
- integration testing against an actual PostgreSQL instance.

Exit project: extend the HTTP service with durable storage and tests that prove transactional behavior and constraint enforcement.

## Stage 5 — Concurrency and systems behavior

Learn:

- goroutines, channels, mutexes, and ownership;
- bounded worker pools, backpressure, cancellation, and timeouts;
- common deadlocks, data races, goroutine leaks, and safe shutdown;
- selecting sequential code when concurrency adds no value;
- the race detector, execution traces, CPU profiles, and memory profiles.

Exit project: a bounded concurrent worker or network service with cancellation, load tests, race-free tests, and documented throughput tradeoffs.

## Stage 6 — Production engineering

Learn:

- module and dependency management;
- linting, CI, reproducible builds, and release artifacts;
- containers and local development environments;
- logs, metrics, traces, health checks, and operational debugging;
- authentication, authorization, secrets, input limits, and secure defaults;
- profiling before optimization and load testing with explicit hypotheses.

Exit standard: deploy and operate a service, diagnose an injected failure, and explain the relevant reliability and security tradeoffs.

## Stage 7 — Hiring preparation

Portfolio evidence should include:

- the finished `ggo` CLI with focused history and tests;
- a production-shaped HTTP and PostgreSQL service;
- one concurrency-oriented tool or service;
- CI, clear READMEs, architecture notes, and reproducible local setup;
- written explanations of tradeoffs, failures encountered, and lessons learned.

Interview preparation should cover:

- explaining slices, maps, interfaces, pointers, errors, and goroutines;
- debugging unfamiliar code and failing tests aloud;
- implementing small features without relying on memorized scaffolds;
- HTTP, SQL, concurrency, testing, and basic distributed-systems reasoning;
- reviewing code for correctness, maintainability, performance, and security.

## Learning cadence

Use a repeating cycle:

1. Attempt the checkpoint without a complete solution.
2. Predict the behavior before running the program.
3. Run it and compare the result with the prediction.
4. Request a review focused on correctness and Go idioms.
5. Fix the findings personally.
6. Add tests for discovered failure modes.
7. Explain the implementation from memory.
8. Revisit the concept after several days in a different problem.

Every few checkpoints, complete a no-assistance exercise. These retrieval checks reveal whether knowledge is becoming usable independently rather than merely recognizable.
