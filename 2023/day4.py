import os
path = os.path.dirname(__file__)

with open(f"{path}/input/day4.txt", "r") as file:
    data = [line.strip().split(":") for line in file.readlines()]

tickets = []
for i, d in enumerate(data):
    d[1] = " ".join(d[1].split())
    d[1] = d[1].split("|")
    d[1] = [line.strip().split(" ") for line in d[1]]
    d[0] = d[0].replace("Card ", "")
    tickets.append({
        "ticket num": d[0],
        "winning nums": d[1][0],
        "my nums": d[1][1]
    })

accumulator1 = 0
for i, ticket in enumerate(tickets, 1):
    n = len([x for x in ticket["my nums"]if x in ticket["winning nums"]])
    # Using a geometric sequence r**(n-1) -> 1, 2, 4, 8, 16
    accumulator1 += int(2**(n-1))

print(accumulator1)
