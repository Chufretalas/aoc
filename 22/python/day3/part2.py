def get_char_value(char: str):

    if char.isupper():
        return ord(char)-38  # -64 +26

    else:
        return ord(char)-96


def find_common(sacks: list[str]) -> str:
    for char in sacks[0]:
        if char in sacks[1] and char in sacks[2]:
            return char


def main():

    total_points = 0

    with open("./input.txt", "r") as file:
        flag = False
        while not flag:
            sacks = []
            for _ in range(3):
                sack = file.readline()
                if sack:
                    sacks.append(sack.strip())
                else:
                    flag = True
                    break

            if flag:
                break

            common = find_common(sacks)
            points = get_char_value(common)
            total_points += points
        file.close()

    print(total_points)


if __name__ == "__main__":
    main()
