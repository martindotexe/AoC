import sys
from itertools import combinations


def partOne(filepath):
    with open(filepath, "r") as file:
        lines = [line.split() for line in file.readlines()]
        total = 0

        for target, *buttons, _ in lines:
            buttons = [set(map(int, button[1:-1].split(","))) for button in buttons]
            target = {index for index, light in enumerate(target[1:-1]) if light == "#"}

            for count in range(1, len(buttons) + 1):
                for attempt in combinations(buttons, count):
                    lights: set[int] = set()
                    for button in attempt:
                        lights ^= button
                    if target == lights:
                        total += count
                        break
                else:
                    continue
                break

        print(total)


if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: python main.py <filepath>", file=sys.stderr)
        sys.exit(1)

    filepath = sys.argv[1]
    partOne(filepath)
