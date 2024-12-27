import os
from functools import cmp_to_key


def parse_list(list_str: str) -> tuple[list, int]:
    new_list = []
    index = 0
    while index < len(list_str):
        char = list_str[index]

        if char == "[":
            sub_list, jump = parse_list(list_str[index+1:])
            new_list.append(sub_list)
            index += jump

        if char.isnumeric():
            number = []
            number.append(list_str[index])
            i = index + 1
            while True:
                if list_str[i].isnumeric():
                    number.append(list_str[i])
                else:
                    break
                index += 1
                i += 1
            new_list.append(int("".join(number)))

        if char == "]":
            break

        index += 1

    return new_list, index + 1


def compare_pair(left, right):
    size = len(left)

    for i in range(size):
        if i == len(right):  # left is bigger than right
            return -1

        if type(left[i]) is int and type(right[i]) is int:
            if left[i] < right[i]:
                return 1

            elif left[i] > right[i]:
                return -1

            else:
                continue

        if type(left[i]) is list and type(right[i]) is list:
            result = compare_pair(left[i], right[i])
            if type(result) is int:
                return result

        if type(left[i]) != type(right[i]):
            if type(left[i]) is int:
                result = compare_pair([left[i]], right[i])
                if type(result) is int:
                    return result

            elif type(right[i]) is int:
                result = compare_pair(left[i], [right[i]])
                if type(result) is int:
                    return result

    if len(left) < len(right):
        return 1

    return None


def main():
    input_path = os.path.join(os.path.dirname(__file__), "input.txt")
    with open(input_path, "r") as file:

        pairs: list[list] = []
        pair = []

        while True:
            line = file.readline()
            if not line:
                pairs.append(tuple(pair))
                break

            line = line.strip()

            if line == "":
                pairs.append(tuple(pair))
                pair = []
                continue

            packet, _ = parse_list(line[1:])
            pair.append(packet)

        ordered = []
        for l, r in pairs:
            ordered.append(l)
            ordered.append(r)

        ordered.append([[2]])
        ordered.append([[6]])

        ordered.sort(key=cmp_to_key(compare_pair), reverse=True)

        v1 = ordered.index([[2]]) + 1
        v2 = ordered.index([[6]]) + 1

        print(v1*v2)

        # for p in ordered:
        #     print(p)

        file.close()


if __name__ == "__main__":
    main()
