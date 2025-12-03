def partOne():
    with open("in.txt", "r") as file:
        lines = file.read()

    count = 0

    for line in lines.strip().split("\n"):
        line = list(map(int, list(line)))

        l, r = 0, 1
        max = -1
        while r < len(line):
            curr = line[l] * 10 + line[r]

            if curr > max:
                max = curr
            if line[l] < line[r]:
                l = r

            r += 1

        count += max

    print(count)


if __name__ == "__main__":
    partOne()
