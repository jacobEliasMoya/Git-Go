# Go Learning Workspace

This repository is a hands-on path toward employable Go skills. The first project is `git-go`, a small Git workflow CLI whose eventual executable will be named `ggo`.

The learner writes the implementation. Codex acts primarily as a teacher and reviewer: it may inspect code, explain concepts, identify defects, suggest focused exercises, and run safe checks. It should not silently complete later checkpoints or replace the learning process with a finished solution.

## Curriculum

The learning experience is organized as **GoQuest**, an original gamified curriculum with quests, XP, mastery gates, boss battles, testing progression, Go-first explanations, and spaced retrieval. Cross-language comparisons are included only when the learner explicitly requests them.

- [`git-go/LEARNING_ROADMAP.md`](git-go/LEARNING_ROADMAP.md): the complete course and skill tree.
- [`git-go/PROGRESS.md`](git-go/PROGRESS.md): current quest, XP, mastery score, retrieval queue, and badges.
- [`git-go/PHASE_NOTES.md`](git-go/PHASE_NOTES.md): session explanations and checkpoint evidence.

## Current state

The learner is in World 1, Level 1.1: building and explaining a minimal Go executable. The implementation runs; the compiler lab, teach-back, and transfer task remain before the next level unlocks.

The longer-lived Codex session guidance is in [`GIT_GO_CODEX_HANDOFF.md`](GIT_GO_CODEX_HANDOFF.md).

## Learning rules

- Attempt each checkpoint before requesting an implementation.
- Ask for code review after saving the relevant files.
- Explain important code in your own words before advancing.
- Test success paths, edge cases, and failure paths.
- Prefer the Go standard library until a dependency solves a demonstrated need.
- Keep commits small enough to show the learning progression.
- Treat terminal, Go, and Git commands as part of the curriculum: explain them before use and interpret their results afterward.
