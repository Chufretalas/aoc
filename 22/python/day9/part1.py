import os
from copy import copy

def move(head, tail, direction, value):
    for _ in range(value):
        past_position = copy(head)
        match direction:
            case "U":
                head[1] += 1
            case "D":
                head[1] -= 1
            case "R":
                head[0] += 1
            case "L":
                head[0] -= 1

        if head[0] == tail[0] or head[1] == tail[1]: # direct line
            if abs(head[0] - tail[0]) > 1 or abs(head[1] - tail[1]) > 1:
                tail = past_position

        else: # diagonal
            if (abs(head[0] - tail[0]) + abs(head[1] - tail[1])) > 2:
                tail = past_position

    return head, tail

def main():
    input_path = os.path.join(os.path.dirname(__file__), "input.txt")
    with open(input_path, "r") as file:

        all_tail_positions = [(0, 0)]
        head = [0, 0] # horizontal, vertical
        tail = [0, 0]

        while True:
            line = file.readline()
            if not line:
                break

            line = line.strip()

            direction, value = line.split(" ")
            value = int(value)
            
            head, tail = move(head, tail, direction, value)
            all_tail_positions.append(tuple(copy(tail)))

        print(head, tail)
        print(len(set(all_tail_positions)))

        file.close()


if __name__ == "__main__":
    main()
