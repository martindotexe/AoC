import sys


def partOne(filepath):
    min = 0
    max = 99
    dial = 50
    count = 0

    with open(filepath, "r") as file:
        lines = file.readlines()
        for line in lines:
            dir: str = line[0]
            num: int = int(line[1:])

            dial = dial - num if dir == "L" else dial + num

            if dial < min or dial > max:
                dial = dial % (max + 1)

            if dial == 0:
                count += 1

    print(count)


if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: python main.py <filepath>", file=sys.stderr)
        sys.exit(1)

    filepath = sys.argv[1]
    partOne(filepath)
