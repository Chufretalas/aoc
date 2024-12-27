def get_score(other: str, you: str) -> int:
    choice_points = 0
    match you:
        case "X": choice_points = 1
        case "Y": choice_points = 2
        case "Z": choice_points = 3

    if (other == "A" and you == "X") or \
        (other == "B" and you == "Y") or \
            (other == "C" and you == "Z"):
        return 3 + choice_points  # draw

    elif (other == "A" and you == "Z") or \
        (other == "B" and you == "X") or \
            (other == "C" and you == "Y"):
        return 0 + choice_points  # lost

    else:
        return 6 + choice_points  # won


def main():
    total_points = 0
    with open("./input.txt", "r") as file:
        while True:
            line = file.readline()

            if not line:
                break

            other, you = line.strip().split(" ")

            total_points += get_score(other, you)

        print(total_points)
        file.close()


if __name__ == "__main__":
    main()
