import os

def main():
    input_path = os.path.join(os.path.dirname(__file__), "input.txt")
    with open(input_path, "r") as file:

        # Variables here

        while True:
            line = file.readline()
            if not line:
                break

            line = line.strip()

            # line by line logic here

        # more logic here

        file.close()


if __name__ == "__main__":
    main()
