import sys
import re


def partOne(filepath):
    with open(filepath, "r") as file:
        lines = file.read().split("\n\n")[-1].splitlines()
        total = 0

        for line in lines:
            x, y, *counts = list(map(int, re.findall(r"\d+", line)))
            if (x // 3) * (y // 3) >= sum(counts):
                total += 1
        print(total)


if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: python main.py <filepath>", file=sys.stderr)
        sys.exit(1)

    filepath = sys.argv[1]
    partOne(filepath)
