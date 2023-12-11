with open("input/day3.txt", "r") as file:
    m = [line.strip()+"." for line in file.readlines()]

# m = [
#     "....*.........................................816*588..............152",
#     "..36.290.831....374................579.536.....................408....",
#     "...........+..../...........795/.....*.*.....................%........"
# ]
# m = [x.strip()+"." for x in m]

m.insert(0, "."*len(m[0]))
m.insert(len(m), "."*len(m[0]))


def find_num_pos(m):
    # m is the map of digits and other sumbols.
    # The function finds the start and end pos of digits.
    xpos = (0, 0)
    ypos = (0, 0)
    num_pos = []

    for row, schema in enumerate(m[1:len(m)-1], 1):
        for col, _ in enumerate(schema[0:len(schema)-1]):
            if m[row][col].isdigit() and xpos == (0, 0):
                xpos = (row, col)
            if m[row][col+1].isdigit() is False and xpos != (0, 0):
                ypos = (row, col+1)
                num_pos.append((xpos, ypos))
                xpos = (0, 0)
                ypos = (0, 0)

    return num_pos


def is_adjacent_symbol(m, pos) -> bool:
    # loops over translated positions in the map.
    #          |          |
    # (-1, -1) | (-1,  0) | (-1,  1)
    # --------------------------------
    #          |          |
    # (0,  -1) |          | (0 ,  1)
    #          |          |
    # --------------------------------
    #          |          |
    # (1 , -1) | (1 ,  0) | (1 ,  1)
    #          |          |

    for y in range(pos[0][1], pos[1][1]):
        x = pos[0][0]
        for xTrans, yTrans in [(-1, -1), (-1, 0), (-1, 1), (0, -1), (0, 1), (1, -1), (1, 0), (1, 1)]:
            if m[x+xTrans][y+yTrans].isdigit() is False and m[x+xTrans][y+yTrans] != ".":
                return True
    return False


def check_adjacent_symbols(m, positions):
    nums = []
    for pos in positions:
        if is_adjacent_symbol(m, pos):
            nums.append(int(m[pos[0][0]][pos[0][1]:pos[1][1]]))
    return nums


# print("\n".join(m))
pos = find_num_pos(m)
nums = check_adjacent_symbols(m, pos)
print(sum(nums))
