def partOne():
    with open("in.txt", "r") as file:
        fresh, ingredients = file.read().split("\n\n")

        fresh = [list(map(int, d.split("-"))) for d in fresh.splitlines()]
        ingredients = list(map(int, ingredients.splitlines()))

        count = 0
        for ingredient in ingredients:
            for lo, hi in fresh:
                if lo <= ingredient <= hi:
                    count += 1
                    break

        print(count)


def partTwo():
    with open("in.txt", "r") as file:
        fresh, _ = file.read().split("\n\n")

        fresh = [list(map(int, d.split("-"))) for d in fresh.splitlines()]
        fresh.sort()

        last = None
        count = 0

        for lo, hi in fresh:
            if last is None:
                last = (lo, hi)
            elif last[1] < lo:
                count += last[1] - last[0] + 1
                last = (lo, hi)
            else:
                last = (last[0], max(last[1], hi))

        count += last[1] - last[0] + 1  # type: ignore

        print(count)


if __name__ == "__main__":
    partOne()
    partTwo()
