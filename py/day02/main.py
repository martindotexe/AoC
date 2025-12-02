def invalidId(id: int):
    str_id = str(id)
    if len(str_id) % 2 != 0:
        return False
    return str_id[:len(str_id)//2] == str_id[len(str_id)//2:]

def validPattern(pattern: str, string: str):
    for i in range(0, len(string), len(pattern)):
        if string[i:i+len(pattern)] != pattern:
            return False
    return True

def advancedInvalidId(id: int):
    str_id = str(id)
    possible_patterns: list[str] = []

    l, r = 1, len(str_id)-1

    while l <= r:
        if str_id[:l] == str_id[r:]:
            possible_patterns.append(str_id[:l])

        l+=1
        r-=1


    for pattern in possible_patterns:
        if validPattern(pattern, str_id):
            return True

    return False

def partOne():
    count = 0
    with open("in.txt", "r") as file:
        lines = file.read()
    
    for line in lines.replace("\n", "").split(","):
        l, r = map(int, line.split("-"))
        for i in range(l, r):
            if invalidId(i):
                count += i

    print(count)

def partTwo():
    count = 0
    with open("in.txt", "r") as file:
        lines = file.read()
    
    for line in lines.replace("\n", "").split(","):
        l, r = map(int, line.split("-"))
        for i in range(l, r+1):
            if advancedInvalidId(i):
                count += i

    print(count)

if __name__ == "__main__":
    partOne()
    partTwo()
