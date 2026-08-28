# GoQuest Progress Tracker

## Player card

- Current title: Gopher Initiate
- Current world: World 2 — Input Scanner
- Current level: 2.1 — Observe command-line arguments
- Total verified XP: 80
- Independence rating: ★★☆
- Current streak: tracked by completed study sessions, not calendar pressure

## Current quest

### Level 2.1 — Observe command-line arguments

Objective: inspect user-provided command-line arguments and understand the executable entry, slicing, indexing, and bounds risk.

Progress:

- [x] Import `os` and inspect `os.Args`.
- [x] Exclude the executable entry with `os.Args[1:]`.
- [x] Display the executable base name with `filepath.Base`.
- [x] Read two user arguments by index.
- [x] Identify the index-out-of-range risk.
- [ ] Validate zero user arguments.
- [ ] Validate one user argument.
- [ ] Decide the policy for extra arguments.
- [ ] Explain why slicing does not guarantee later indexes exist.

Current mastery:

| Lane | Score | Evidence needed next |
|---|---:|---|
| Behavior | 1/2 | Define and handle zero, one, and extra arguments |
| Quality | 1/2 | Add bounds validation and verify the input matrix |
| Explanation | 1/2 | Explain slice length and indexing risk |
| Transfer | 1/2 | Safely vary the accepted command shape |
| **Total** | **4/8** | Continue with validation before mastery |

Known risk: the current implementation panics unless at least two user arguments are supplied.

## Earlier quest

### Level 1.1 — Start the program

Objective: create, run, and explain the smallest Go executable.

Progress:

- [x] Implement a `main` package.
- [x] Define `func main()`.
- [x] Import and use `fmt`.
- [x] Print `Hello There`.
- [x] Format the file with `gofmt`.
- [x] Run the program successfully.
- [ ] Predict and inspect an unused-import compiler error.
- [ ] Explain `package main` in your own words.
- [ ] Explain `func main()` in your own words.
- [ ] Explain why unused imports are rejected.
- [ ] Explain the compile-and-run lifecycle of `go run .`.
- [ ] Reconstruct the minimal program from an empty scratch file.

Verified XP:

| Activity | XP | Status |
|---|---:|---|
| Implement the behavior | 25 | Earned |
| Verify with tools | 15 | Earned |
| Focused documentation | 10 | Pending summary |
| Prediction | 10 | Pending |
| Compiler lab | 20 | Pending |
| Teach-back | 10 | Pending |
| Transfer task | 10 | Pending |

Mastery score:

| Lane | Score | Evidence needed next |
|---|---:|---|
| Behavior | 2/2 | Program compiles and prints expected output |
| Quality | 1/2 | Complete compiler lab and package test check |
| Explanation | 0/2 | Answer the understanding questions |
| Transfer | 0/2 | Reconstruct or vary the program independently |
| **Total** | **3/8** | Unlock requires at least 7/8 and no zero lane |

## Next unlock

World 2 — Input Scanner unlocks after Level 1.1 mastery.

First unlocked quest:

- inspect `os.Args`;
- describe the type and contents of `os.Args` using Go terminology;
- predict and observe argument indexing behavior.

## Retrieval queue

Add actual completion dates when Level 1.1 is mastered.

| Concept | 1 day | 3 days | 7 days | 14 days | Later project |
|---|---|---|---|---|---|
| `package main` and entry point | Pending | Pending | Pending | Pending | Data CLI |
| imports and exported identifiers | Pending | Pending | Pending | Pending | HTTP API |
| `go run` versus builds | Pending | Pending | Pending | Pending | `ggo` release |

## Badge cabinet

- [ ] Boot Sequence — build and explain a minimal executable.
- [ ] Argument Archaeologist — safely inspect and explain CLI input.
- [ ] Boundary Guardian — validate input without panics or side effects.
- [ ] Table Tester — write meaningful table-driven tests.
- [ ] Error Cartographer — preserve and explain failure context.
- [ ] Process Wrangler — execute and test child processes safely.
- [ ] Git Pipeline Keeper — prove ordered Git behavior with integration tests.
- [ ] Binary Smith — build and distribute `ggo` cleanly.
- [ ] Type Foundry Adept — demonstrate Go's core type and memory behavior.
- [ ] HTTP Ranger — build and test a standard-library API.
- [ ] Transaction Keeper — preserve database consistency under failure.
- [ ] Race Hunter — find and eliminate a real data race.
- [ ] Production Ready — operate and diagnose a deployed service.
- [ ] Go Hireable — complete the portfolio and final hiring boss.

## Session log template

Copy this block into `PHASE_NOTES.md` for each meaningful session:

```markdown
## YYYY-MM-DD — World.Level: Quest name

Goal:

Prediction:

What I changed:

Commands and what they did:

Tests and results:

Bug or surprise:

Optional cross-language connection, only if I requested one:

What I can explain now:

What still feels unclear:

XP earned:

Mastery score:

Next retrieval date:
```

## Progress rules

- Check a box only when there is evidence.
- Do not award explanation or transfer XP based only on working output.
- Do not erase mistakes; record the lesson they exposed.
- Reattempts can raise mastery and independence ratings.
- Asking for a hint never removes XP.
- Update this tracker before each checkpoint commit.
