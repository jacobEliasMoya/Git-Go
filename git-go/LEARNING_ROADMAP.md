# GoQuest: An Employable Go Curriculum

## Mission

GoQuest is a project-based curriculum for becoming employable with Go. It combines short lessons, hands-on quests, testing, deliberate debugging, Go-first explanations, spaced retrieval, and portfolio projects.

The goal is not to collect completed files. The goal is to develop skills that can be recalled, explained, tested, and transferred to unfamiliar problems.

The learner writes the implementation. Codex acts as instructor, reviewer, debugger, and curriculum guide. Codex should not silently complete a quest that the learner is meant to practice.

## Course artifacts

- `LEARNING_ROADMAP.md`: course rules, skill tree, levels, and mastery requirements.
- `PROGRESS.md`: current quest, XP, mastery scores, review schedule, and unlocks.
- `PHASE_NOTES.md`: explanations, commands, discoveries, mistakes, and checkpoint results.
- Git history: evidence of small, understandable milestones.
- Tests: executable evidence that behavior and failure handling work.

## The quest loop

Every quest follows the same seven-step loop:

1. **Brief** — define one observable outcome and the concepts it introduces.
2. **Predict** — state what the code or command should do before running it.
3. **Build** — implement the smallest version personally.
4. **Verify** — format, compile, run, and inspect the result.
5. **Test** — exercise normal behavior, boundaries, and failures appropriate to the level.
6. **Explain** — describe the important lines and decisions without reading a prepared answer.
7. **Transfer** — solve a small variation that was not copied from the lesson.

A quest is not mastered merely because the program printed the expected output.

## XP system

Each standard quest is worth 100 XP:

| Activity | XP | Evidence |
|---|---:|---|
| Read the focused references | 10 | Summarize the relevant idea |
| Predict behavior | 10 | Record a prediction before execution |
| Implement the behavior | 25 | Working learner-written code |
| Verify with tools | 15 | Relevant commands and interpreted output |
| Test or deliberately break it | 20 | Tests, edge cases, or compiler experiment |
| Teach it back | 10 | Explanation in the learner's own words |
| Complete the transfer task | 10 | Uncopied variation or application |

Boss battles are worth 200 XP. Retrieval challenges are worth 25 XP. Portfolio releases are worth 300 XP.

XP provides motivation and a history of effort. It never overrides a failed mastery gate.

## Mastery gate

Every checkpoint is scored from 0–2 in four lanes:

| Lane | 0 | 1 | 2 |
|---|---|---|---|
| Behavior | Missing or broken | Happy path works | Required behavior and boundaries work |
| Quality | Unsafe or unchecked | Formatted and basically verified | Appropriate tests and static checks pass |
| Explanation | Cannot yet explain | Explains syntax | Explains behavior, tradeoffs, and failure paths |
| Transfer | Cannot modify it | Completes a guided variation | Solves an unfamiliar variation independently |

Unlock requirement:

- at least 7/8 total;
- no lane may be 0;
- required tests must pass;
- the learner must make or approve the checkpoint commit.

## Independence rating

Mastery and independence are tracked separately:

- ★☆☆ — completed with direct guidance or focused code.
- ★★☆ — completed with conceptual hints or pseudocode.
- ★★★ — completed from the brief and documentation without implementation help.

There is no penalty for asking for help. A later reattempt can improve the independence rating.

## Hint ladder

When blocked, request the smallest helpful rung:

1. official documentation link;
2. diagnostic question;
3. conceptual explanation;
4. another Go example or compiler experiment;
5. test case or failing example;
6. pseudocode;
7. focused code fragment;
8. full solution only when explicitly requested or when the lesson is no longer serving as practice.

Codex should normally begin at the lowest useful rung.

## Retrieval schedule

Important concepts return after approximately:

- 1 day;
- 3 days;
- 7 days;
- 14 days;
- one later project.

A review should require recall, prediction, debugging, or application. Rereading notes alone does not count as retrieval practice.

## Go-first explanations

Teach each concept using Go's vocabulary, rules, examples, compiler behavior, and official documentation. Do not introduce comparisons to TypeScript, JavaScript, or another language unless the learner explicitly asks for one.

When requested, a comparison should address one focused question, identify where the analogy breaks, and then return to the Go model. The target is an independent understanding of idiomatic Go.

# Skill tree

## World 0 — Base Camp: Toolchain and workflow

### Level 0.1: Locate the toolchain

Quest:

- identify the installed Go version and platform;
- locate the active `go` executable;
- explain what an exit code represents.

Commands:

- `go version`
- PowerShell `Get-Command go`

Docs:

- [Download and install Go](https://go.dev/doc/install)
- [PowerShell command discovery](https://learn.microsoft.com/powershell/scripting/discover-powershell)

### Level 0.2: Understand the module boundary

Quest:

- locate `go.mod`;
- explain the module directive and Go version;
- distinguish a module, package, directory, and repository.

Commands:

- `go env GOMOD`
- `go list ./...`

Docs:

- [Go modules reference](https://go.dev/ref/mod)
- [`go list`](https://pkg.go.dev/cmd/go#hdr-List_packages_or_modules)

### Level 0.3: Control Git state

Quest:

- explain working tree, staging index, commit, branch, and remote;
- inspect a change before staging;
- prepare and approve a focused commit.

Docs:

- [`git status`](https://git-scm.com/docs/git-status)
- [`git diff`](https://git-scm.com/docs/git-diff)
- [`git add`](https://git-scm.com/docs/git-add)
- [`git commit`](https://git-scm.com/docs/git-commit)
- [`git push`](https://git-scm.com/docs/git-push)

Boss battle: clone the repository into a clean location, verify the module, run the project, make a harmless documentation change, inspect it, and discard or commit it intentionally.

## World 1 — Boot Sequence: Minimal Go executable

### Level 1.1: Start the program

Main quest:

- create a `main` package;
- define the program entry point;
- print one intentional line.

Concepts:

- package declarations;
- imports;
- functions;
- exported identifiers;
- standard output;
- compilation before execution.

Verification:

- `gofmt -w main.go`
- `go run .`
- `go test ./...` as a build check even before test files exist.

Compiler lab:

- temporarily create an unused import;
- predict the compiler response;
- run the program and interpret the error;
- restore the valid source.

Transfer task:

- change the output without changing the program structure;
- explain why the import is still required.

Docs:

- [Getting started tutorial](https://go.dev/doc/tutorial/getting-started)
- [`fmt` package](https://pkg.go.dev/fmt)
- [`go run`](https://pkg.go.dev/cmd/go#hdr-Compile_and_run_Go_program)
- [`gofmt`](https://pkg.go.dev/cmd/gofmt)

Boss battle: reconstruct the minimal executable from an empty file without looking at the prior version.

## World 2 — Input Scanner: Arguments, strings, and slices

### Level 2.1: Observe `os.Args`

Main quest:

- print the raw argument slice;
- identify the executable path;
- observe no arguments, one argument, and a quoted multi-word argument.

Concepts:

- variables and short declarations;
- strings;
- slices;
- indexes;
- length;
- terminal quoting.

Optional cross-language comparison, only when explicitly requested:

- compare `os.Args` with Node's `process.argv`;
- compare a Go slice with a JavaScript array while identifying type and memory differences.

Safety lab:

- predict what happens when indexing an absent element;
- trigger the failure deliberately;
- read the panic;
- add a safe length check.

Docs:

- [`os.Args`](https://pkg.go.dev/os#Args)
- [Go slices introduction](https://go.dev/blog/slices-intro)
- [Go specification: index expressions](https://go.dev/ref/spec#Index_expressions)

### Level 2.2: Name and display inputs

Main quest:

- assign the flag and message to meaningfully named variables;
- print them separately;
- explain type inference.

Transfer task:

- display the number of user-provided arguments without counting the executable path.

Boss battle: write a small argument inspector from memory that safely describes every received argument.

## World 3 — Gatekeeper: Validation and control flow

### Level 3.1: Reject incomplete input

Main quest:

- reject missing arguments before indexing them;
- produce a concise usage message;
- stop execution with an early return.

Concepts:

- `if`;
- comparisons;
- early returns;
- validation order;
- safe boundaries.

### Level 3.2: Accept only `-F`

Main quest:

- accept the initial fix flag;
- reject unknown flags;
- keep repository-changing commands out of the program.

Test matrix:

- no arguments;
- only `-F`;
- unknown flag;
- valid `-F` and message;
- extra arguments.

Optional cross-language comparison, only when explicitly requested:

- compare early-return validation with request validation in an API route;
- contrast Go's lack of truthy/falsy coercion with JavaScript.

Docs:

- [Go specification: if statements](https://go.dev/ref/spec#If_statements)
- [Effective Go: control structures](https://go.dev/doc/effective_go#control-structures)

Boss battle: implement the validation again from a written behavior table rather than existing code.

## World 4 — Message Forge: Functions, strings, and first tests

### Level 4.1: Extract a pure formatter

Main quest:

- turn the user message into `(fix) Message.`;
- avoid duplicating final punctuation;
- return a new string from a helper.

Concepts:

- typed parameters;
- return types;
- pure functions;
- string operations;
- naming and function boundaries.

### Level 4.2: Write table-driven tests

Test cases:

- ordinary message;
- message already ending in a period;
- empty message policy;
- surrounding whitespace policy;
- punctuation cases chosen deliberately.

Concepts:

- `_test.go` files;
- `testing.T`;
- table-driven tests;
- subtests;
- expected versus actual output;
- regression tests.

Commands:

- `go test ./...`
- `go test -v ./...`
- `go test -run <pattern> ./...`

Optional cross-language comparison, only when explicitly requested:

- compare a Go table-driven test with parameterized Jest or Vitest tests;
- compare explicit return types with TypeScript annotations.

Docs:

- [Add a test tutorial](https://go.dev/doc/tutorial/add-a-test)
- [`testing` package](https://pkg.go.dev/testing)
- [`strings` package](https://pkg.go.dev/strings)

Boss battle: receive a failing hidden-style test case, diagnose the invariant, fix the formatter, and add a regression test.

## World 5 — Process Portal: Errors and external commands

### Level 5.1: Run one harmless command

Main quest:

- create a child process with `os/exec`;
- pass the program and arguments separately;
- connect output to the terminal;
- observe success and failure.

Concepts:

- processes;
- arguments versus shell strings;
- standard input/output/error;
- exit status;
- returned errors.

### Level 5.2: Propagate useful errors

Main quest:

- return command errors from a helper;
- add context without losing the original error;
- stop the workflow on failure.

Optional cross-language comparison, only when explicitly requested:

- compare `exec.Command` with Node's `child_process`;
- compare explicit `error` values with thrown exceptions and rejected promises.

Docs:

- [`os/exec`](https://pkg.go.dev/os/exec)
- [`errors`](https://pkg.go.dev/errors)
- [`fmt.Errorf`](https://pkg.go.dev/fmt#Errorf)

Boss battle: run a known-success command and a known-failure command, then explain every output stream and error value.

## World 6 — Git Pipeline: Safe side effects

### Level 6.1: Create a command boundary

Main quest:

- separate Git argument construction from process execution;
- record commands in tests without running real Git;
- prove later steps stop after a failure.

### Level 6.2: Stage in a disposable repository

Main quest:

- create a temporary Git repository;
- run the staging step there;
- inspect the index afterward.

### Level 6.3: Commit safely

Main quest:

- pass the formatted message as one argument to `-m`;
- preserve Git output;
- stop when the commit fails.

### Level 6.4: Push only after success

Main quest:

- introduce pushing last;
- test sequencing without contacting a real remote;
- perform a real push only with explicit authorization.

Testing:

- fake-runner unit tests;
- temporary-directory integration tests;
- disposable Git repository tests;
- failure injection at every step.

Docs:

- [`os.MkdirTemp`](https://pkg.go.dev/os#MkdirTemp)
- [`testing.T.TempDir`](https://pkg.go.dev/testing#T.TempDir)
- [`git init`](https://git-scm.com/docs/git-init)
- [`git add`](https://git-scm.com/docs/git-add)
- [`git commit`](https://git-scm.com/docs/git-commit)
- [`git push`](https://git-scm.com/docs/git-push)

Boss battle: demonstrate with tests that push cannot occur after an add or commit failure.

## World 7 — Release Gate: Build and distribute `ggo`

Main quest:

- build `ggo.exe`;
- explain `go run`, `go build`, and `go install`;
- keep generated binaries out of Git;
- run the command from another disposable repository;
- document installation and removal.

Quality gate:

- `gofmt` clean;
- unit tests pass;
- integration tests pass;
- `go vet ./...` passes;
- README contains reproducible examples;
- errors are actionable;
- no real remote is changed during automated tests.

Docs:

- [`go build`](https://pkg.go.dev/cmd/go#hdr-Compile_packages_and_dependencies)
- [`go install`](https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies)
- [How to Write Go Code](https://go.dev/doc/code)

Portfolio release: tag and document the first maintainable `ggo` version.

# Core Go campaign

## World 8 — Type Foundry

Quests:

- primitive types, constants, inference, and conversions;
- arrays versus slices;
- slice length, capacity, append, copy, and backing arrays;
- maps and the comma-ok idiom;
- structs and field access;
- zero values and `nil`.

Boss battle: build an in-memory issue tracker with typed statuses, slices, maps, validation, and table-driven tests.

Docs:

- [Go Tour](https://go.dev/tour/)
- [Go specification: types](https://go.dev/ref/spec#Types)
- [Slices usage and internals](https://go.dev/blog/slices-intro)

## World 9 — Method Mountains

Quests:

- methods;
- value receivers;
- pointer receivers;
- mutation and copying;
- constructors as ordinary functions;
- embedding and composition.

Boss battle: model a small domain, predict which mutations persist, and prove the behavior with tests.

Docs:

- [Effective Go: methods](https://go.dev/doc/effective_go#methods)
- [Go specification: method declarations](https://go.dev/ref/spec#Method_declarations)

## World 10 — Interface Crossing

Quests:

- implicit interface satisfaction;
- interfaces defined by consumers;
- small interfaces;
- type assertions and type switches;
- `any` and when not to use it;
- dependency seams for testing.

Boss battle: replace one concrete dependency with a minimal consumer-owned interface without producing an abstraction layer maze.

Docs:

- [Go specification: interface types](https://go.dev/ref/spec#Interface_types)
- [Effective Go: interfaces](https://go.dev/doc/effective_go#interfaces_and_types)

## World 11 — Error Caverns

Quests:

- creating errors;
- wrapping errors;
- `errors.Is` and `errors.As`;
- sentinel and typed errors;
- errors at package boundaries;
- logging versus returning;
- `defer` and cleanup.

Boss battle: debug a multi-layer failure while preserving enough context for the caller to make a decision.

Docs:

- [Working with Errors in Go](https://go.dev/blog/go1.13-errors)
- [`errors` package](https://pkg.go.dev/errors)
- [Effective Go: defer](https://go.dev/doc/effective_go#defer)

# Portfolio campaigns

## World 12 — Data Explorer CLI

Build a second CLI that reads JSON, CSV, logs, or repository data and produces a useful report.

Required skills:

- `io.Reader` and `io.Writer`;
- files and paths;
- JSON or CSV encoding;
- configuration;
- package design;
- dependency injection through ordinary functions and small interfaces.

Testing ladder:

- pure unit tests;
- temporary files and directories;
- malformed input;
- golden output where justified;
- fuzzing for parser invariants;
- benchmarks only for an explicit performance question.

Boss battle: implement a new input case from documentation and tests without instructor code.

Docs:

- [`io`](https://pkg.go.dev/io)
- [`os`](https://pkg.go.dev/os)
- [`encoding/json`](https://pkg.go.dev/encoding/json)
- [`encoding/csv`](https://pkg.go.dev/encoding/csv)
- [Go fuzzing tutorial](https://go.dev/doc/tutorial/fuzz)

## World 13 — HTTP Province

Build a service primarily with `net/http` before adopting a framework.

Quests:

- handlers and routing;
- JSON request and response contracts;
- validation and consistent errors;
- middleware;
- request-scoped context;
- cancellation and timeouts;
- graceful shutdown;
- structured logging;
- configuration.

Tests:

- handler unit tests;
- `httptest` requests and recorders;
- full HTTP integration tests;
- malformed bodies;
- unknown resources;
- cancellation and timeout behavior.

Boss battle: add a new endpoint from a behavioral specification with tests and no implementation hint.

Docs:

- [`net/http`](https://pkg.go.dev/net/http)
- [`net/http/httptest`](https://pkg.go.dev/net/http/httptest)
- [`context`](https://pkg.go.dev/context)
- [JSON and Go](https://go.dev/blog/json)

## World 14 — PostgreSQL Depths

Extend the HTTP service with durable storage.

Quests:

- relational modeling;
- keys, constraints, and indexes;
- parameterized queries;
- scanning rows;
- transactions and rollback;
- migrations;
- connection pools;
- cancellation;
- idempotency;
- schema/application compatibility.

Tests:

- integration tests against PostgreSQL;
- constraint failures;
- rollback behavior;
- concurrent update behavior;
- migration from an older schema.

Boss battle: diagnose and fix a consistency bug using a transaction and a regression test.

Docs:

- [`database/sql`](https://pkg.go.dev/database/sql)
- [Executing transactions](https://go.dev/doc/database/execute-transactions)
- [Accessing relational databases](https://go.dev/doc/database/)

## World 15 — Concurrency Citadel

Quests:

- goroutines;
- channels;
- `select`;
- mutexes;
- ownership;
- bounded worker pools;
- backpressure;
- cancellation;
- deadlocks;
- goroutine leaks;
- data races.

Tests and tools:

- deterministic concurrency tests;
- `go test -race ./...`;
- controlled timeouts;
- load tests;
- execution traces;
- CPU and memory profiles.

Boss battle: build a bounded worker service that shuts down cleanly, leaks no goroutines, and passes the race detector.

Docs:

- [Go concurrency patterns](https://go.dev/blog/pipelines)
- [Go memory model](https://go.dev/ref/mem)
- [Data race detector](https://go.dev/doc/articles/race_detector)
- [`sync` package](https://pkg.go.dev/sync)

## World 16 — Production Keep

Quests:

- module and dependency maintenance;
- CI;
- linting and static analysis;
- reproducible builds;
- containers;
- environment configuration and secrets;
- health checks;
- logs, metrics, and traces;
- authentication and authorization;
- input and resource limits;
- dependency auditing;
- profiling before optimization.

Boss battle: deploy a service, inject a failure, diagnose it using operational signals, and write a prevention or mitigation test.

## World 17 — Hiring Arena

Portfolio releases:

- tested `ggo` CLI;
- file/data CLI;
- PostgreSQL-backed HTTP service;
- concurrency-oriented worker or service;
- clear READMEs, architecture notes, CI, and reproducible setup.

Interview quests:

- explain slices, maps, structs, interfaces, pointers, errors, and goroutines;
- debug unfamiliar Go code aloud;
- review code for correctness, maintainability, performance, and security;
- implement small features without memorized scaffolding;
- discuss HTTP, SQL, concurrency, and testing tradeoffs;
- compare Go and TypeScript accurately without overstating similarities.

Final boss: receive an unfamiliar but bounded backend task, clarify requirements, implement it, test it, review the diff, and explain the design under time constraints.

# Testing mastery ladder

Testing capabilities unlock in this order:

1. compile and run;
2. deliberate compiler-error experiments;
3. unit tests for pure functions;
4. table-driven tests and subtests;
5. fake or recorded dependency boundaries;
6. temporary files and repositories;
7. HTTP integration tests;
8. database integration tests;
9. fuzz tests;
10. race detection;
11. benchmarks and profiling;
12. CI quality gates.

Coverage percentage is diagnostic information, not the objective. Tests should protect important behavior, boundaries, invariants, and failure paths.

# Instructor protocol

For each new quest, Codex should provide:

1. the observable objective;
2. why the skill matters professionally;
3. prerequisite check;
4. focused official references;
5. a second Go example or experiment when useful;
6. a prediction prompt;
7. acceptance criteria;
8. commands with concise explanations;
9. a code review after the learner's attempt;
10. a mastery score and next unlock.

Routine commands should be explained fully on first use and briefly on later repetition. New flags, side effects, destructive operations, and remote effects always require explanation.

# Definition of course completion

The curriculum is complete when the learner can independently:

- build and maintain multiple Go programs;
- use the standard library confidently;
- design clear packages and boundaries;
- write layered unit and integration tests;
- diagnose compiler, runtime, race, database, and HTTP failures;
- build a production-shaped service;
- explain tradeoffs during code review;
- reason about Go directly and write idiomatic Go;
- present tested projects and reason effectively in hiring interviews.
