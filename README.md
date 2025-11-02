# Text Auto-Formatter (Go)

A simple command-line tool written in Go that performs automatic text completion, editing and autocorrection according to a set of transformation rules (hex/bin conversion, case changes, punctuation normalization, quotes, and article correction).

This README documents: setup, architecture, how to run the program and tests, an AI-agent-driven workflow for entry-level developers, and a granular, ordered Agile task list that follows Test-Driven Development (TDD). It also includes instructions for auditors (peer reviewers).

---

## Project Overview

The program accepts two arguments:

```
go run . <input-file> <output-file>
```

It reads the input file, applies transformations, and writes the result to the output file.

Transformation rules supported (high-level):

- Convert `(<hex>)` and `(<bin>)` markers to decimal representation of the preceding word.
- Case conversions: `(up)`, `(low)`, `(cap)` and their numeric variants `(up, N)` etc., affecting previous N words.
- Punctuation normalization for `. , ! ? : ;` and grouped punctuation such as `...` and `!?`.
- Single-quote normalization: `'` marks must enclose content tightly with no internal spaces and preserve multi-word quotes.
- Article correction: turn `a` into `an` when the following word begins with a vowel or `h`.

---

## Architecture

This project uses a **Pipeline** architecture where each transformation rule is implemented as a separate, testable stage (filter) that receives and returns a sequence of tokens or text. This makes the code modular, easy to test, and extendable.

Each stage (example): `tokenize -> hex/bin conversion -> case transformations -> article correction -> punctuation normalization -> quotes normalization -> stringify`.

---

## Repo layout (recommended)

```
/ (module root)
  cmd/
    textfmt/    # main package (cmd/textfmt/main.go)
  internal/
    pipeline/   # pipeline orchestration code + stage interfaces
    rules/      # implementations for each rule (hex, bin, up, low, cap, punctuation, quotes, article)
  testdata/
    golden/     # golden test input and expected outputs (from golden test set)
  pkg/
    tokenizer/  # optional tokenizer package
  go.mod
  README.md
  analysis_document.md    # provided
  golden_test_set.md      # provided
```

---

## How to run

1. Build/run directly with `go run ./cmd/textfmt <infile> <outfile>`.
2. Run unit tests: `go test ./...`.
3. Regression/golden tests: `go test ./... -run Golden` or a dedicated `integration_test.go` that reads files from `testdata/golden`.

---

## Agents & Roles

This project uses an **AI agent + human entry-level developer** workflow. The intended roles are:

- **Senior Architect (you / this README)**: defines architecture, writes high-level tasks, reviews auditor reports.
- **Entry-level Developer**: implements tasks one-by-one following TDD and uses the agent for boilerplate scaffolding and suggestions.
- **AI Agent**: auto-generates unit-test skeletons, suggests function stubs, converts simple regex into code, and runs static checks locally when asked by the developer.
- **Auditor (peer-reviewer)**: other students act as graders who run the golden tests, read code, and provide feedback.

Agents must not push code directly; they should produce candidate changes the developer applies.

---

## Testing Strategy

- Unit tests for each rule and edge cases.
- Integration tests that run the full pipeline on golden inputs and compare to expected outputs.
- Property tests for idempotency where appropriate (running formatter twice should not change correct output further).

---

## Task List (Ordered, TDD-first)

_Important_: Each task below is small and self-contained. Follow the order. For each task: first write failing tests, then implement enough code to make tests pass, then validate.

### Task 1 — Project skeleton & CI
- **TDD (write tests)**: Add `go.mod`, create `cmd/textfmt/main_test.go` with tests that assert the CLI returns an error when arguments are missing.
- **Implementation goal**: Create `cmd/textfmt/main.go` skeleton that parses CLI args and returns proper exit codes.
- **Validation**: `go test ./cmd/textfmt -run TestCLIArgs` passes.

### Task 2 — Tokenizer basics
- **TDD**: Write unit tests for a tokenizer that splits text into words, punctuation, and markers (e.g., `(hex)`, `(up, 2)`), and preserves whitespace boundaries as metadata.
- **Implementation goal**: Implement `tokenizer.Tokenize(string) ([]Token, error)` producing tokens for words, punctuation and markers.
- **Validation**: Tokenizer tests pass and represent sample lines from the golden set.

### Task 3 — Pipeline stage interface
- **TDD**: Add tests for the pipeline orchestration: given a list of stages, the pipeline must apply them in order and produce expected tokens.
- **Implementation goal**: Implement `pipeline.New(...stages)` and `pipeline.Run(tokens)`.
- **Validation**: Pipeline tests pass and a mock stage can transform tokens.

### Task 4 — Hexadecimal conversion rule
- **TDD**: Write unit tests for `rules.HexRule` that ensure `1E (hex)` becomes `30` and that lowercase hex like `1e` works.
- **Implementation goal**: Implement a rule that finds the `(hex)` marker and converts the previous word from hex to decimal, replacing the token and removing the marker.
- **Validation**: Rule unit tests pass and a simple integration test confirms correct replacement.

### Task 5 — Binary conversion rule
- **TDD**: Tests for `rules.BinRule` that `10 (bin)` → `2`, `1111 (bin)` → `15`, and that invalid binary returns an error or leaves token unchanged (specify chosen behavior in test).
- **Implementation goal**: Implement binary conversion analogous to hex.
- **Validation**: `go test ./internal/rules -run TestBinRule` passes and integration confirms conversions.

### Task 6 — Case rules: (up), (low), (cap) single-word
- **TDD**: Tests for single-word `(up)`, `(low)`, `(cap)` behavior.
- **Implementation goal**: Implement simple case transformations for the preceding single word.
- **Validation**: Tests pass and `it (cap)` → `It`.

### Task 7 — Case rules: numeric variants `(cap, N)`, `(up, N)`, `(low, N)`
- **TDD**: Tests covering multi-word transformations.
- **Implementation goal**: Implement logic to find previous N word tokens and apply transformation.
- **Validation**: Tests pass and the golden examples with numbered variants match expected results.

### Task 8 — Case precedence & overlapping markers
- **TDD**: Tests for interactions where multiple markers appear close together.
- **Implementation goal**: Ensure deterministic rule order in pipeline.
- **Validation**: Tests pass and examples behave correctly.

### Task 9 — Article correction rule (a → an)
- **TDD**: Tests where `a` is followed by words starting with vowels and `h`.
- **Implementation goal**: Implement `rules.ArticleRule`.
- **Validation**: Tests pass and golden examples show `a honor` → `an honor`.

### Task 10 — Quotation normalization (`'` rule)
- **TDD**: Tests ensuring quotes attach tightly to their content: `' awesome '` → `'awesome'`.
- **Implementation goal**: Implement quote trimming logic.
- **Validation**: Tests pass and match golden cases.

### Task 11–20 — (Remaining tasks: punctuation, integration, idempotency, CLI, documentation, CI)
(Refer to full README in project for details.)

---

## Auditor Instructions

1. Clone the repo.
2. Run `go test ./...` — all unit tests should pass.
3. Run `make golden` or `scripts/run_golden.sh` to validate golden test cases.
4. For each failing case, create an issue and assign it to the auditee. Use the provided golden test input and expected output to reproduce.

---

## Notes & best practices

- Keep each rule small, single-responsibility, and well-documented.
- Prefer pure functions for rule transformations for easy testing.
- Add table-driven tests for edge cases.
- When in doubt, add a golden test that captures the desired behavior.

---

## Acknowledgements

Inputs used to produce this README: the project analysis and golden test set provided by the assignment.
