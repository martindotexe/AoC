# Advent of Code Solutions
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/martindotexe/aoc/go.yml?style=flat&logo=adventofcode&logoSize=auto&label=Tests&labelColor=%230f0f23)

This repository contains my solutions to [Advent of Code](https://adventofcode.com/) challenges.

## Solutions by Year

- 2024: 11/50 ⭐
- 2023: 7/50 ⭐
- 2022: 1/50 ⭐
- 2021: 5/50 ⭐
## Project Structure

The different years are separated into their own branches. 
```
2024/
├── day01/
│   ├── solution.go       # Solution code
│   ├── solution_test.go  # Tests
└── ...
```

## Running Solutions

To run a specific day's solution:

```bash
go run main.go -y XXXX -d XX
```

To run tests:

```bash
go test ./...
```
