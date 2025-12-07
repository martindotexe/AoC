def partOne():
    with open("in.txt", "r") as file:
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
            print("".join(["|" if c in curr else row[c] for c in range(len(row))]))

        print(splits)


if __name__ == "__main__":
    partOne()
