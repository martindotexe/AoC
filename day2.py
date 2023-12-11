with open("input/day2.txt", "r") as file:
    raw_data = [line.strip() for line in file.readlines()]

data = []

for line in raw_data:
    game_num, games = line.split(":")
    game_num = int(game_num.replace("Game ", ""))
    games = [g.split(",") for g in games.split(";")]
    for i, g in enumerate(games):
        games[i] = [tuple(x.strip().split(" ")) for x in g]
    data.append({
        "game_num": game_num,
        "games": games
    })

accumulator1 = 0
accumulator2 = 0
for d in data:
    reds = [int(g[0]) for g in sum(d["games"], []) if g[1] == "red"]
    greens = [int(g[0]) for g in sum(d["games"], []) if g[1] == "green"]
    blues = [int(g[0]) for g in sum(d["games"], []) if g[1] == "blue"]
    if max(reds) <= 12 and max(greens) <= 13 and max(blues) <= 14:
        accumulator1 += d["game_num"]
    accumulator2 += max(reds) * max(greens) * max(blues)

print(f"Part One: {accumulator1}")
print(f"Part Two: {accumulator2}")
