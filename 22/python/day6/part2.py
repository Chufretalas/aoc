import os

def main():
    input_path = os.path.join(os.path.dirname(__file__), "input.txt")
    with open(input_path, "r") as file:
        stream = file.readline()
        stream = stream.strip("\n")
        sequence = []
        for index, letter in enumerate(stream):
            sequence.append(letter)
            if len(sequence) > 14:
                sequence.pop(0)

            if len(sequence) == 14 and len(set(sequence)) == 14:
                print(index+1)
                break

if __name__ == "__main__":
    main()