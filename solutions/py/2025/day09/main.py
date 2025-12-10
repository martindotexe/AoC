import sys


def partOne(filepath):
    with open(filepath, "r") as file:
        points = [tuple(map(int, line.split(","))) for line in file.readlines()]

        print(
            max(
                [
                    (abs(x2 - x1) + 1) * (abs(y2 - y1) + 1)
                    for i, (x1, y1) in enumerate(points)
                    for x2, y2 in points[i:]
                ]
            )
        )


if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: python main.py <filepath>", file=sys.stderr)
        sys.exit(1)

    filepath = sys.argv[1]
    partOne(filepath)
