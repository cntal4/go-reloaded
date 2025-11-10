# Usage Guide

## Command Line Interface

### Basic Syntax

```bash
go run ./cmd/textfmt <input-file> <output-file>
```

### Arguments

- `<input-file>`: Path to the input text file to be formatted
- `<output-file>`: Path where the formatted output will be written

### Exit Codes

- `0`: Success
- `1`: Error (invalid arguments, file not found, etc.)

### Examples

#### Basic Usage
```bash
go run ./cmd/textfmt input.txt output.txt
```

#### Using Built Binary
```bash
# Build first
go build ./cmd/textfmt

# Then use
./textfmt input.txt output.txt
```

#### Error Handling
```bash
# Missing arguments
go run ./cmd/textfmt
# Output: Usage: textfmt <input-file> <output-file>

# File not found
go run ./cmd/textfmt missing.txt output.txt
# Output: Error: reading input: open missing.txt: no such file or directory
```

## Transformation Rules

### Number Conversions

#### Hexadecimal to Decimal
- **Syntax**: `<number> (hex)`
- **Example**: `42 (hex)` → `66`
- **Supports**: Both uppercase and lowercase hex digits

#### Binary to Decimal
- **Syntax**: `<number> (bin)`
- **Example**: `1010 (bin)` → `10`
- **Supports**: Only 0 and 1 digits

### Case Transformations

#### Single Word
- **Uppercase**: `word (up)` → `WORD`
- **Lowercase**: `WORD (low)` → `word`
- **Capitalize**: `word (cap)` → `Word`

#### Multiple Words
- **Syntax**: `(command, N)` where N is the number of previous words
- **Examples**:
  - `hello world (up, 2)` → `HELLO WORLD`
  - `HELLO WORLD (low, 2)` → `hello world`
  - `hello world (cap, 2)` → `Hello World`

### Article Correction

- **Rule**: Changes 'a' to 'an' before words starting with vowels (a, e, i, o, u) or 'h'
- **Examples**:
  - `a apple` → `an apple`
  - `a honor` → `an honor`
  - `a book` → `a book` (unchanged)

### Quote Normalization

- **Rule**: Removes spaces immediately inside single quotes
- **Examples**:
  - `' hello '` → `'hello'`
  - `' hello world '` → `'hello world'` (preserves internal spaces)

### Punctuation Normalization

- **Rule**: Removes space before punctuation, adds single space after
- **Punctuation**: `. , ! ? : ;`
- **Examples**:
  - `hello , world` → `hello, world`
  - `wait ... what` → `wait... what`
  - `really !? yes` → `really!? yes`

## Advanced Usage

### Combining Rules

Multiple transformation rules can be applied in a single input:

```
Input:  there (cap) once was a hero with 1e (hex) coins and said ' hello world ' !
Output: There once was an hero with 30 coins and said' hello world'!
```

### Processing Order

The transformations are applied in this order:
1. Hex/Binary conversion
2. Case transformations
3. Article correction
4. Quote normalization
5. Punctuation normalization

### Error Handling

- Invalid hex/binary numbers are left unchanged
- Unknown markers are ignored
- Malformed input produces partial output without crashing
- File errors produce clear error messages

## Tips

1. **Test with small inputs first** to understand the behavior
2. **Use golden tests** to verify expected behavior: `go test ./testdata/golden`
3. **Check examples** in `docs/examples/` for reference
4. **Run with logging** to see processing steps (logs go to stderr)