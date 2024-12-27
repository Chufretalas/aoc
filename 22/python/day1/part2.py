def main():
    all_elves = []
    calories = 0
    with open("./input.txt", "r") as file:
        while True:
            line = file.readline()
            if line:
                if line == "\n":
                    all_elves.append(calories)
                    calories = 0
                else:
                    calories += int(line.removesuffix("\n"))
            else:
                break

        top_three = sorted(all_elves)[-3:]
        print(sum(top_three))

        file.close()


if __name__ == "__main__":
    main()
