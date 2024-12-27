import timeit

def main():
    biggest_calories = 0
    calories = 0
    with open("./input.txt", "r") as file:
        while True:
            line = file.readline()
            if line:
                if line == "\n":
                    if calories > biggest_calories: biggest_calories = calories
                    calories = 0
                else:
                    calories += int(line.removesuffix("\n"))
            else:
                break

        print(biggest_calories)

        file.close()


if __name__ == "__main__":
    start = timeit.default_timer()
    main()
    print((timeit.default_timer() - start)*1000000, "us")
