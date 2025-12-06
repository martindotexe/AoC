def partOne():
    with open("in.txt", "r") as file:
        input = file.read().splitlines()

        operators = input[-1].split()
        numbers = [list(map(int, digits.split())) for digits in input[:-1]]

        aggr = 0
        for col in range(len(numbers[0])):
            count = 0
            operator = operators[col]
            for row in range(len(numbers)):
                curr = numbers[row][col]
                if operator == "*":
                    count = count * curr if count != 0 else curr
                elif operator == "+":
                    count += numbers[row][col]
                elif operator == "-":
                    count -= numbers[row][col]
            aggr += count

        print(aggr)


if __name__ == "__main__":
    partOne()
