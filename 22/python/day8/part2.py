from __future__ import annotations
import os

# coord[0] == row / coord[1] == column
def calc_score(coord: tuple[int, int], trees: list[str], tree_value: int) -> int:

    if coord[0] == 0 or coord[0] == len(trees)-1:
        return 0
    elif coord[1] == 0 or coord[1] == len(trees[0])-1:
        return 0

    top = 0
    right = 0
    bottom = 0
    left = 0

    # -- check top --
    for i in range(coord[0]):  # if the tree is in row 1, i will start in 0
        top = i + 1
        search_coord = (coord[0]-1-i, coord[1])
        if int(trees[search_coord[0]][search_coord[1]]) >= tree_value:
            break

    # -- check right --
    for i in range(len(trees[0]) - coord[1] - 1):
        right = i + 1
        search_coord = (coord[0], coord[1]+1+i)
        if int(trees[search_coord[0]][search_coord[1]]) >= tree_value:
            break

    # -- check bottom --
    for i in range(len(trees) - coord[0] - 1):
        bottom = i + 1
        search_coord = (coord[0]+1+i, coord[1])
        if int(trees[search_coord[0]][search_coord[1]]) >= tree_value:
            break

    # -- check left --
    for i in range(coord[1]):
        left = i + 1
        search_coord = (coord[0], coord[1]-1-i)
        if int(trees[search_coord[0]][search_coord[1]]) >= tree_value:
            break

    return top * right * bottom * left


def main():
    input_path = os.path.join(os.path.dirname(__file__), "input.txt")
    with open(input_path, "r") as file:

        trees = []

        while True:
            line = file.readline()
            if not line:
                break

            line = line.strip("\n")

            trees.append(line)

        non_zero_values = []
        for row, line in enumerate(trees):
            for column, tree in enumerate(line):
                tree = int(tree)
                result = calc_score((row, column), trees, tree)
                if result != 0: non_zero_values.append(result)
                print((row, column), result)

        print(max(non_zero_values))
        file.close()


if __name__ == "__main__":
    main()
