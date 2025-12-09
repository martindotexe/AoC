import sys


def partOne(filepath):
    count = 0
    with open(filepath, "r") as file:
        for line in file.readlines():
            line = [int(d) for d in line.strip()]

            tens = max(line[:-1])
            ones = max(line[line.index(tens) + 1 :])

            count += tens * 10 + ones

        print(count)


def partTwo(filepath):
    N = 12

    count = 0
    with open(filepath, "r") as file:
        for line in file.readlines():
            line = [int(d) for d in line.strip()]

            r = 0
            for i in range(N - 1, 0, -1):
                digit = max(line[:-i])
                line = line[line.index(digit) + 1 :]
                r = r * 10 + digit

            r = r * 10 + max(line)
            count += r

        print(count)


if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: python main.py <filepath>", file=sys.stderr)
        sys.exit(1)

    filepath = sys.argv[1]
    partOne(filepath)
    partTwo(filepath)
