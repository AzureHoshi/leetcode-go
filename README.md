# Go Algorithms & LeetCode Solutions

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Test](https://github.com/AzureHoshi/leetcode-go/actions/workflows/test.yml/badge.svg)](https://github.com/AzureHoshi/leetcode-go/actions)

A curated collection of algorithmic problem solutions and data structures implemented in **idiomatic Go**. Each solution is accompanied by **table-driven unit tests**, **performance benchmarks**, and **time/space complexity analysis** designed for engineering interviews and production-grade practice.

---

## 🚀 Features

- **Idiomatic Go**: Adheres to Go best practices and standard library idioms.
- **Table-Driven Tests**: Thorough test coverage including normal, edge, and corner cases.
- **Benchmarks & Memory Profiling**: Evaluates execution speed and memory allocations (`b.ReportAllocs()`).
- **Complexity Documented**: Explicit Big-O analysis for both time and space on every solution.
- **Automated CI**: GitHub Actions workflow automatically verifies tests on each commit.

---

## 📂 Project Structure

```text
go-algorithms/
├── .github/
│   └── workflows/
│       └── test.yml                 # Automated CI test runner
├── leetcode/
│   ├── 0001-two-sum/
│   │   ├── solution.go              # Algorithm implementation & complexity notes
│   │   └── solution_test.go         # Table-driven unit tests and benchmarks
├── .gitignore
├── go.mod
└── README.md
```

---

## 📊 Solved Problems

| # | Problem | Difficulty | Category | Time | Space | Solution |
| :---: | :--- | :---: | :--- | :---: | :---: | :---: |
| 0001 | Two Sum | `Easy` | Array / Hash Table | $O(n)$ | $O(n)$ | [Go](file:///leetcode/0001-two-sum/solution.go) |

---

## 🛠️ Getting Started & Testing

### Prerequisites
- [Go](https://go.dev/dl/) 1.21 or higher installed

### Run Unit Tests
To run all tests across all problems:
```bash
go test -v ./...
```

To run with race condition detector:
```bash
go test -v -race ./...
```

### Run Benchmarks
To run performance benchmarks and inspect memory allocation per operation:
```bash
go test -bench=. -benchmem ./...
```

---

## 📝 Conventions for Adding New Solutions

1. Create a dedicated folder under `leetcode/` using the format `XXXX-problem-name` (e.g., `0020-valid-parentheses`).
2. Implement the solution in `solution.go` with descriptive function comments explaining the approach and Big-O complexity.
3. Write comprehensive unit test cases in `solution_test.go` using Go's table-driven test pattern.
4. Add a benchmark function `BenchmarkXxx(b *testing.B)` to verify performance.
5. Update the problems table in this `README.md`.

---

## 📄 License

This repository is licensed under the [MIT License](LICENSE).
