import os

def main():
    input_path = os.path.join(os.path.dirname(__file__), "input.txt")
    with open(input_path, "r") as file:

        cicles = [1]
        register = 1

        while True:
            line = file.readline()
            if not line:
                break

            line = line.strip()

            if line == "noop":
                cicles.append(register)
                continue

            else:
                cicles.append(register)
                cicles.append(register)
                _, value = line.split(" ")
                value = int(value)
                register += value

        final = cicles[20]*20 + cicles[60]*60 + cicles[100]*100 + cicles[140]*140 + cicles[180]*180 + cicles[220]*220
        print(final)

        file.close()


if __name__ == "__main__":
    main()
