def partOne():
    with open("in.txt", "r") as file:
        fresh, ingredients = file.read().split("\n\n")

        fresh = [tuple(map(int, d.split("-"))) for d in fresh.strip().split("\n")]
        ingredients = [int(d) for d in ingredients.strip().split("\n")]

        count = 0
        for ingredient in ingredients:
            for f in fresh:
                if f[0] <= ingredient <= f[1]:
                    count += 1
                    break

        print(count)


if __name__ == "__main__":
    partOne()
