# Advent of Code Solutions in Go
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/martindotexe/AoC/go.yml?branch=2023&style=flat&logo=adventofcode&logoSize=auto&label=2023&labelColor=%230E0E24)

This repository contains my solutions to [Advent of Code](https://adventofcode.com/) challenges implemented in Go. 
Each year's solutions are maintained in a separate branch.

## Solutions by Year

- [2023](https://github.com/martindotexe/AoC/tree/2023): 2 ⭐ out of 50 possible stars.

## Project Structure

The different years are separated into their own branches. 
```
puzzles /
├── day01/
│   ├── solution.go       # Solution code
│   ├── solution_test.go  # Tests
└── ...
```

## Running Solutions

To run a specific day's solution:

```bash
go run main.go -d XX
```

To run tests:

```bash
go test ./...
```

## Branches

- `main` - Contains this README
- `2023` - Solutions for Advent of Code 2023
