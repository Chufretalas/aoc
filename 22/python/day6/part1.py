import os

def main():
    input_path = os.path.join(os.path.dirname(__file__), "input.txt")
    with open(input_path, "r") as file:
        stream = file.readline()
        stream = stream.strip("\n")
        sequence = []
        for index, letter in enumerate(stream):
            sequence.append(letter)
            if len(sequence) > 4:
                sequence.pop(0)

            if len(sequence) == 4 and len(set(sequence)) == 4:
                print(index+1)
                break

if __name__ == "__main__":
    main()