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

### Task 11 — Punctuation: single punctuation normalization
**Goal**  
Attach single punctuation marks `. , ! ? : ;` to the previous token (no space before) and ensure exactly one space after, unless punctuation ends a line.

**TDD (write tests)**  
- Tests that `"word ,next"` → `"word, next"`.
- Tests that punctuation at end-of-line is handled correctly.
- Edge cases: consecutive punctuation treated specially in next task.

**Implementation**  
- Create `pkg/processors/punctuation_basic.go` implementing `PunctuationBasicProcessor`.
- Processor should normalize token spacing and attach punctuation to the left word.

**Validation**  
- `go test ./pkg/processors -run TestPunctuationBasic` passes.

---

### Task 12 — Punctuation: grouped punctuation handling
**Goal**  
Detect grouped punctuation sequences like `...`, `!?`, `!!`, `??` and format them as a single punctuation token with no preceding space and a single following space.

**TDD (write tests)**  
- Tests for `"I was thinking ... You"` → `"I was thinking... You"`.
- Tests for `"Really !? no"` → `"Really!? no"`.
- Ensure `...` remains intact and not split into three `.` tokens.

**Implementation**  
- `PunctuationGroupedProcessor` that merges consecutive punctuation tokens into a single grouped token.

**Validation**  
- `go test ./pkg/processors -run TestPunctuationGrouped` passes.

---

### Task 13 — Quotes processor (single quotes)
**Goal**  
Normalize single-quote tokens so that ` ' word ' ` becomes `'word'` and multi-word quotes keep internal spaces (`'a phrase here'`). Ensure punctuation next to quotes is handled correctly.

**TDD (write tests)**  
- Tests for `' awesome '` → `'awesome'`.
- Tests for `' I am the man '` → `'I am the man'`.
- Tests for quotes with punctuation: `'hello ,'` → `'hello,'`.

**Implementation**  
- `QuoteProcessor` that finds paired Quote tokens, trims internal edges, and preserves internal spacing.

**Validation**  
- `go test ./pkg/processors -run TestQuoteProcessor` passes.

---

### Task 14 — Integration Tests
**File:** `tasks/14_integration_tests.md`  
**Goal:**  
Combine all existing processors (Tokenizer, HexBin, Case, Article, Quote, Punctuation) into a single functional pipeline.  
Ensure token flow, transformations, and order integrity are maintained end-to-end.

**TDD Steps:**  
- Write integration tests that tokenize, process, and stringify text back to readable form.  
- Compare pipeline output to known-correct (“golden”) text.  
- Test invalid marker handling gracefully.

**Implementation Notes:**  
- Implement pipeline composition under `internal/pipeline/pipeline.go`.  
- Ensure each processor conforms to the `Processor` interface.  
- Add benchmarks for performance regression checks.

**Validation:**  
`go test ./internal/pipeline -run TestIntegrationPipeline -v`  
should pass all golden examples.

---

### Task 15 — Idempotency & Property Tests
**File:** `tasks/15_idempotency_test.md`  
**Goal:**  
Guarantee the formatter is idempotent — multiple runs yield identical results.

**TDD Steps:**  
- Add property tests: `format(format(text)) == format(text)`.  
- Generate multiple random small strings and ensure stability.  
- Add benchmarks for round-trip consistency.

**Implementation Notes:**  
- Implement test utilities under `testutils/`.  
- Focus on processors that mutate text (Case, Punctuation).

**Validation:**  
`go test ./testutils -run TestIdempotency`  
should pass and show no diffs between first and second runs.

---

### Task 16 — CLI Integration
**File:** `tasks/16_cli_integration.md`  
**Goal:**  
Integrate file reading/writing in `cmd/textfmt/main.go` with command-line argument handling.

**TDD Steps:**  
- Mock file I/O and run CLI with temp files.  
- Validate correct exit codes and stdout/stderr messages.  
- Test missing input/output args and invalid file names.

**Implementation Notes:**  
- Use `os.ReadFile` and `os.WriteFile`.  
- Leverage existing pipeline runner to process content.

**Validation:**  
`go test ./cmd/textfmt -run TestCLIIntegration -v`  
and manual:  
`go run . sample.txt result.txt` → expected output file matches golden text.

---

### Task 17 — Error Handling
**File:** `tasks/17_error_handling.md`  
**Goal:**  
Add safe error handling, logging, and fallback logic across processors and CLI.

**TDD Steps:**  
- Test behavior for invalid markers, malformed tokens, unreadable files.  
- Ensure pipeline returns partial but stable output rather than crashing.  
- Verify proper `os.Exit(1)` on unrecoverable errors.

**Implementation Notes:**  
- Implement a `logger` helper under `internal/logger`.  
- Use Go’s `errors.Join` or `fmt.Errorf` with `%w` for traceable wrapping.

**Validation:**  
`go test ./internal/logger -v`  
`go run . missing.txt result.txt` → prints meaningful error without panic.

---

### Task 18 — Golden Tests
**File:** `tasks/18_golden_tests.md`  
**Goal:**  
Validate final formatting behavior against golden input/output text files.

**TDD Steps:**  
- Load golden pairs from `docs/golden test set.md`.  
- Compare processed text to expected outputs.  
- Output diffs in a unified diff format when mismatches occur.

**Implementation Notes:**  
- Store test fixtures under `testdata/golden/`.  
- Use `t.Helper()` and `cmp.Diff()` for clean output.

**Validation:**  
`go test ./testdata/golden -v`  
shows all golden tests passing with no diff output.

---

### Task 19 — Documentation & Developer Help
**File:** `tasks/19_docs_help.md`  
**Goal:**  
Finalize README examples, usage, and quick developer setup documentation.

**TDD Steps:**  
- Validate that `README.md` code blocks execute successfully.  
- Ensure `go run . sample.txt result.txt` matches documented examples.  
- Spellcheck and linkcheck documentation.

**Implementation Notes:**  
- Add `docs/examples/` folder with sample input/output.  
- Include command-line usage guide and pipeline diagram.

**Validation:**  
`markdownlint README.md`  
and visual README rendering verified on GitHub preview.

---

### Task 20 — CI & Peer Review Automation
**File:** `tasks/20_ci_peer_review.md`  
**Goal:**  
Automate testing, linting, and peer auditing in CI/CD.

**TDD Steps:**  
- CI should run `go test ./...` + `go vet ./...`.  
- Include a golden test sweep.  
- Verify that CI status badges appear in README.

**Implementation Notes:**  
- Add `.github/workflows/ci.yml` with build/test/lint steps.  
- Include a `scripts/run_audit.sh` for local auditing workflow.  
- Peer review guide in `docs/audit_checklist.md`.

**Validation:**  
All GitHub Actions checks pass and PR template enforces peer audit completion.

