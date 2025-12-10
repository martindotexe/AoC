import sys


def partOne(filepath):
    with open(filepath, "r") as file:
        points = [tuple(map(int, line.split(","))) for line in file.readlines()]

        pointPairs = [
            (a, b) for a in range(len(points)) for b in range(a + 1, len(points))
        ]

        largest = 0

        for a, b in pointPairs:
            x1, y1 = points[a]
            x2, y2 = points[b]

            area = (abs(x2 - x1) + 1) * (abs(y2 - y1) + 1)
            if area > largest:
                largest = area

        print(largest)


if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: python main.py <filepath>", file=sys.stderr)
        sys.exit(1)

    filepath = sys.argv[1]
    partOne(filepath)
