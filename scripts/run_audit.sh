#!/bin/bash

# Local Audit Script for go-reloaded
# This script runs all the checks that would be performed in CI/CD

set -e

echo "🔍 Starting Local Audit for go-reloaded..."
echo "============================================"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to print status
print_status() {
    if [ $1 -eq 0 ]; then
        echo -e "${GREEN}✅ $2${NC}"
    else
        echo -e "${RED}❌ $2${NC}"
        exit 1
    fi
}

print_info() {
    echo -e "${YELLOW}🔧 $1${NC}"
}

# Check if we're in the right directory
if [ ! -f "go.mod" ]; then
    echo -e "${RED}❌ Error: go.mod not found. Please run this script from the project root.${NC}"
    exit 1
fi

# 1. Format Check
print_info "Checking code formatting..."
UNFORMATTED=$(gofmt -s -l . | grep -v vendor || true)
if [ -n "$UNFORMATTED" ]; then
    echo -e "${RED}❌ The following files are not formatted:${NC}"
    echo "$UNFORMATTED"
    echo -e "${YELLOW}Run: gofmt -s -w .${NC}"
    exit 1
fi
print_status 0 "Code formatting check"

# 2. Go Vet
print_info "Running go vet..."
go vet ./...
print_status $? "Go vet analysis"

# 3. Build Check
print_info "Building project..."
go build ./cmd/textfmt
print_status $? "Build check"

# 4. Unit Tests
print_info "Running unit tests..."
go test ./...
print_status $? "Unit tests"

# 5. Golden Tests
print_info "Running golden tests..."
go test ./testdata/golden -v
print_status $? "Golden tests"

# 6. Race Detector
print_info "Running race detector..."
go test -race -short ./...
print_status $? "Race detector"

# 7. CLI Functionality Test
print_info "Testing CLI functionality..."
echo "hello (cap) world, a amazing test" > /tmp/audit_input.txt
./textfmt /tmp/audit_input.txt /tmp/audit_output.txt 2>/dev/null
EXPECTED="Hello world, an amazing test"
ACTUAL=$(cat /tmp/audit_output.txt)
if [ "$ACTUAL" = "$EXPECTED" ]; then
    print_status 0 "CLI functionality test"
else
    echo -e "${RED}❌ CLI test failed:${NC}"
    echo -e "Expected: $EXPECTED"
    echo -e "Actual:   $ACTUAL"
    exit 1
fi

# 8. Idempotency Test
print_info "Testing idempotency..."
./textfmt /tmp/audit_output.txt /tmp/audit_output2.txt 2>/dev/null
SECOND_RUN=$(cat /tmp/audit_output2.txt)
if [ "$ACTUAL" = "$SECOND_RUN" ]; then
    print_status 0 "Idempotency test"
else
    echo -e "${RED}❌ Idempotency test failed${NC}"
    exit 1
fi

# 9. Coverage Check
print_info "Generating coverage report..."
go test -coverprofile=coverage.out ./... > /dev/null
COVERAGE=$(go tool cover -func=coverage.out | grep total | awk '{print $3}' | sed 's/%//')
echo "Coverage: ${COVERAGE}%"
if (( $(echo "$COVERAGE >= 70" | bc -l) )); then
    print_status 0 "Coverage check (${COVERAGE}%)"
else
    echo -e "${YELLOW}⚠️  Coverage is below 70% (${COVERAGE}%)${NC}"
fi

# Cleanup
rm -f /tmp/audit_input.txt /tmp/audit_output.txt /tmp/audit_output2.txt coverage.out textfmt

echo ""
echo "============================================"
echo -e "${GREEN}🎉 All audit checks passed!${NC}"
echo ""
echo "Summary:"
echo "✅ Code formatting"
echo "✅ Static analysis (go vet)"
echo "✅ Build success"
echo "✅ Unit tests"
echo "✅ Golden tests"
echo "✅ Race condition check"
echo "✅ CLI functionality"
echo "✅ Idempotency"
echo "✅ Coverage report generated"
echo ""
echo "Ready for peer review! 🚀"