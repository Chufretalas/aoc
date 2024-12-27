import os

def draw_line(cicles, register, new_crt_line: list[str]):
    position = len(cicles)-1
    if position > 39:
        number = (position//40) * 40
        position = position - number
    if register-1 <= position <= register + 1:
        new_crt_line.append("#")
    else:
        new_crt_line.append(".")

    return new_crt_line

def main():
    input_path = os.path.join(os.path.dirname(__file__), "input.txt")
    with open(input_path, "r") as file:

        cicles = []
        register = 1
        crt_lines = []
        new_crt_line = []

        while True:
            line = file.readline()
            if not line:
                break

            line = line.strip()

            if line == "noop":
                cicles.append(register)
                new_crt_line = draw_line(cicles, register, new_crt_line)

                if len(new_crt_line) >= 40:
                    crt_lines.append(new_crt_line)
                    new_crt_line = []
                continue

            else:
                cicles.append(register)
                new_crt_line = draw_line(cicles, register, new_crt_line)

                if len(new_crt_line) >= 40:
                    crt_lines.append(new_crt_line)
                    new_crt_line = []

                cicles.append(register)
                new_crt_line = draw_line(cicles, register, new_crt_line)

                if len(new_crt_line) >= 40:
                    crt_lines.append(new_crt_line)
                    new_crt_line = []

                _, value = line.split(" ")
                value = int(value)
                register += value

        for l in crt_lines:
            print("".join(l))

        file.close()


if __name__ == "__main__":
    main()
