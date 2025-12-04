C = "@"


def partOne():
    with open("in.txt", "r") as file:
        lines = [line.strip() for line in file.readlines()]

    w, h = len(lines[0]), len(lines)

    neighbor_grid = [-1 for _ in range(w * h)]

    def xytoi(x, y):
        return x + y * w

    for y in range(h):
        for x in range(w):
            if lines[y][x] != C:
                continue

            neighbor_grid[xytoi(x, y)] += 1

            for dx, dy in [(0, 1), (1, -1), (1, 0), (1, 1)]:
                ix, iy = x + dx, y + dy
                if ix < 0 or ix >= w or iy < 0 or iy >= h:
                    continue

                if lines[iy][ix] != C:
                    continue

                neighbor_grid[xytoi(x, y)] += 1
                neighbor_grid[xytoi(ix, iy)] += 1

    print(len(list(filter(lambda d: d < 4 and d >= 0, neighbor_grid))))


if __name__ == "__main__":
    partOne()
