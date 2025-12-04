C = "@"


def partOne():
    with open("in.txt", "r") as file:
        lines = [line.strip() for line in file.readlines()]

    w, h = len(lines[0]), len(lines)

    # number_grid = [[0 for d in range(w)] for _ in range(h)]
    number_grid = [0 for _ in range(w * h)]

    def xytoi(x, y):
        return x + y * w

    for y in range(h):
        for x in range(w):
            if lines[y][x] != C:
                continue
            for dx, dy in [(0, 1), (1, -1), (1, 0), (1, 1)]:
                ix, iy = x + dx, y + dy
                if ix < 0 or ix >= w or iy < 0 or iy >= h:
                    continue

                if lines[iy][ix] != C:
                    continue

                number_grid[xytoi(x, y)] += 1
                number_grid[xytoi(ix, iy)] += 1

    for i, d in enumerate(map(lambda d: "x" if d < 4 and d != 0 else ".", number_grid)):
        print(d, end="")
        if (i + 1) % w == 0:
            print()
    # for i, d in enumerate(number_grid):
    #     print(d, end="")
    #     if (i + 1) % w == 0:
    #         print()
    print(len(list(filter(lambda d: d < 4 and d != 0, number_grid))))


if __name__ == "__main__":
    partOne()
