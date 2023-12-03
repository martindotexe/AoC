with open("./input/day1.txt", "r") as file:
    inputs = []
    for line in file.readlines():
        inputs.append(line.replace("\n", ""))

numbers = "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"


def find(string, patterns):
    if string != None:
        for i, _ in enumerate(string):
            for pi, pattern in enumerate(patterns, 1):
                if str(pattern) == string[i:len(pattern)+i]:
                    return string.replace(string[i:len(pattern)+i], str(pi), 1)

    return None


def partOne(inputs):
    input_sum = 0
    for input in inputs:
        input_nums = [int(c) for c in input if c.isnumeric()]
        input_sum += input_nums[0]*10 + input_nums[-1]
    return input_sum


def partTwo(inputs):
    new_inputs = []
    for input in inputs:
        cur_string = input
        while True:
            string = find(cur_string, numbers)
            if string == None:
                new_inputs.append(cur_string)
                break
            cur_string = string
    return partOne(new_inputs)


print(f"Part One: {partOne(inputs)}")
print(f"Part Two: {partTwo(inputs)}")
