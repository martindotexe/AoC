def partOne():
    with open("in.txt", "r") as file:
        grid = [line.split() for line in file.readlines()]
        cols = list(zip(*grid))

        aggr = 0
        for *col, op in cols:
            aggr += eval(op.join(col))

        print(aggr)


def partTwo():
    with open("in.txt", "r") as file:
        grid = [line.strip("\n") for line in file.readlines()]
        cols = list(zip(*grid))

        groups = []
        group = []

        for col in cols:
            if set(col) == {" "}:
                groups.append(group)
                group = []
            else:
                group.append(col)

        groups.append(group)

        aggr = 0
        for group in groups:
            aggr += eval(group[0][-1].join(["".join(line) for *line, _ in group]))

        print(aggr)


if __name__ == "__main__":
    partOne()
    partTwo()
