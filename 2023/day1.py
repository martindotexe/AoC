import os
path = os.path.dirname(__file__)

with open(f"{path}/input/day1.txt", "r") as file:
    lines = [line.strip() for line in file.readlines()]

nums1 = []
nums2 = []
accumulator1 = 0
accumulator2 = 0

for line in lines:
    for p, l in enumerate(line):
        if l.isdigit():
            nums1.append(l)
            nums2.append(l)
        for i, let in enumerate(["zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"]):
            if line[p:].startswith(let):
                nums2.append(i)

    accumulator1 += int(f"{nums1[0]}{nums1[-1]}")
    accumulator2 += int(f"{nums2[0]}{nums2[-1]}")
    nums1 = []
    nums2 = []

print(f"Part One: {accumulator1}")
print(f"Part Two: {accumulator2}")
