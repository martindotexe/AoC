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
            aggr += count

        print(aggr)


def partTwo():
    def solve(problem: list[str]):
        count = 0
        operator = problem[-1].strip()
        problem = problem[:-1]
        for col in range(len(problem[0]) - 1, -1, -1):
            curr = int("".join([problem[i][col] for i in range(len(problem))]))
            if operator == "*":
                count = count * curr if count != 0 else curr
            elif operator == "+":
                count += curr

        return count

    with open("in.txt", "r") as file:
        input = file.read().splitlines()

        aggr = 0

        l, r = 0, 1
        problem: list[str] = []
        while r < len(input[0]):
            problem = [i[l:r] for i in input]
            if len(list(filter(lambda c: c != " ", [i[r] for i in input]))) == 0:
                aggr += solve(problem)
                l, r = r + 1, r + 1
            r += 1
        problem = [i[l : r + 1] for i in input]
        aggr += solve(problem)

        print(aggr)


if __name__ == "__main__":
    partOne()
    partTwo()
