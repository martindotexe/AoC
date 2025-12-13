import sys


def distance(p1: list[int], p2: list[int]):
    # Don't need the squre root since we are just comparing distances.
    # Slower genral n-dimensions version: return sum((a - b) ** 2 for a, b in zip(p1, p2))
    return (p1[0] - p2[0]) ** 2 + (p1[1] - p2[1]) ** 2 + (p1[2] - p2[2]) ** 2


def partOne(points, pairs):
    n = 1000

    parent = list(range(len(points)))

    def root(x):
        if parent[x] == x:
            return x
        parent[x] = root(parent[x])
        return parent[x]

    def merge(a, b):
        parent[root(a)] = root(b)

    for a, b in pairs[:n]:
        merge(a, b)

    sizes = [0] * len(points)

    for i in range(len(points)):
        sizes[root(i)] += 1

    sizes.sort(reverse=True)

    print(sizes[0] * sizes[1] * sizes[2])


def partTwo(points, pairs):
    parent = list(range(len(points)))

    def root(x):
        if parent[x] == x:
            return x
        parent[x] = root(parent[x])
        return parent[x]

    def merge(a, b):
        parent[root(a)] = root(b)

    circuits = len(points)

    for a, b in pairs:
        if root(a) == root(b):
            continue
        merge(a, b)
        circuits -= 1
        if circuits == 1:
            print(points[a][0] * points[b][0])
            break


if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: python main.py <filepath>", file=sys.stderr)
        sys.exit(1)

    filepath = sys.argv[1]
    with open(filepath, "r") as file:
        points = [list(map(int, line.split(","))) for line in file.readlines()]
        pairs = [(a, b) for a in range(len(points)) for b in range(a + 1, len(points))]
        # Expensive sort done once for performance
        pairs.sort(key=lambda x: distance(points[x[0]], points[x[1]]))

        partOne(points, pairs)
        partTwo(points, pairs)
