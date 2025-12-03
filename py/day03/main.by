def partOne():
    count = 0
    with open("in.txt", "r") as file:
        for line in file.readlines():
            line = [int(d) for d in line.strip()]

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
