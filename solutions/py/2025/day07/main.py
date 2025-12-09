from functools import cache
import sys


def partOne(filepath):
    with open(filepath, "r") as file:
        grid = [line.strip() for line in file.readlines()]

        curr = {grid[0].index("S")}
        next = []
        splits = 0

        for row in grid:
            for c in curr:
                if row[c] == "^":
                    splits += 1
                    next.append(c - 1)
                    next.append(c + 1)
                else:
                    next.append(c)

            curr = set(next)
            next = []
            # Visualisation
            # print("".join(["|" if c in curr else row[c] for c in range(len(row))]))

        print(splits)


def partTwo(filepath):
    with open(filepath, "r") as file:
        grid = [line.strip() for line in file.readlines()]

        @cache
        def solve(row: int, col: int) -> int:
            if row >= len(grid):
                return 1
            if grid[row][col] == "^":
                return solve(row + 1, col - 1) + solve(row + 1, col + 1)
            return solve(row + 1, col)

        print(solve(0, grid[0].index("S")))


if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: python main.py <filepath>", file=sys.stderr)
        sys.exit(1)

    filepath = sys.argv[1]
    partOne(filepath)
    partTwo(filepath)
