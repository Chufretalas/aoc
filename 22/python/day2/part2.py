def get_score(other: str, result: str) -> int:
    result_score = 0
    you = ""

    if result == "X":  # lose
        result_score = 0
        match other:
            case "A": you = "Z"
            case "B": you = "X"
            case "C": you = "Y"

    elif result == "Y":  # draw
        result_score = 3
        match other:
            case "A": you = "X"
            case "B": you = "Y"
            case "C": you = "Z"

    elif result == "Z":  # win
        result_score = 6
        match other:
            case "A": you = "Y"
            case "B": you = "Z"
            case "C": you = "X"

    choice_points = 0
    match you:
        case "X": choice_points = 1
        case "Y": choice_points = 2
        case "Z": choice_points = 3

    return result_score + choice_points


def main():
    total_points = 0
    with open("./input.txt", "r") as file:
        while True:
            line = file.readline()

            if not line:
                break

            other, result = line.strip().split(" ")

            total_points += get_score(other, result)

        print(total_points)
        file.close()


if __name__ == "__main__":
    main()
