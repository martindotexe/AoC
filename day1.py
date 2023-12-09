with open("input/day1.txt", "r") as file:
    lines = [line.strip() for line in file.readlines()]

nums1 = []
nums2 = []
accumilator1 = 0
accumilator2 = 0

for line in lines:
    for p, l in enumerate(line):
        if l.isdigit():
            nums1.append(l)
            nums2.append(l)
        for i, let in enumerate(["zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"]):
            if line[p:].startswith(let):
                nums2.append(i)

    accumilator1 += int(f"{nums1[0]}{nums1[-1]}")
    accumilator2 += int(f"{nums2[0]}{nums2[-1]}")
    nums1 = []
    nums2 = []

print(f"Part One: {accumilator1}")
print(f"Part Two: {accumilator2}")
