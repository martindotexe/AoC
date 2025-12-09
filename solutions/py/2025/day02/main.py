import sys


def invalidId(id: int):
    str_id = str(id)
    if len(str_id) % 2 != 0:
        return False
    return str_id[: len(str_id) // 2] == str_id[len(str_id) // 2 :]


def validPattern(pattern: str, string: str):
    string_len = len(string)
    patter_len = len(pattern)
    if string_len == 2:
        return string[0] == string[1]

    for i in range(patter_len, string_len, patter_len):
        if string[i : i + patter_len] != pattern:
            return False
    return True


def advancedInvalidId(id: int):
    str_id = str(id)
    possible_patterns: list[str] = []

    l, r = 1, len(str_id) - 1

    while l <= r:
        if len(str_id) % l == 0 and str_id[:l] == str_id[r:]:
            possible_patterns.append(str_id[:l])

        l += 1
        r -= 1

    for pattern in possible_patterns:
        if validPattern(pattern, str_id):
            return True

    return False


def partOne(filepath):
    count = 0
    with open(filepath, "r") as file:
        lines = file.read()

    for line in lines.replace("\n", "").split(","):
        l, r = map(int, line.split("-"))
        for i in range(l, r):
            if invalidId(i):
                count += i

    print(count)


def partTwo(filepath):
    count = 0
    with open(filepath, "r") as file:
        lines = file.read()

    for line in lines.replace("\n", "").split(","):
        l, r = map(int, line.split("-"))
        for i in range(l, r + 1):
            if advancedInvalidId(i):
                count += i

    print(count)


if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: python main.py <filepath>", file=sys.stderr)
        sys.exit(1)

    filepath = sys.argv[1]
    partOne(filepath)
    partTwo(filepath)
