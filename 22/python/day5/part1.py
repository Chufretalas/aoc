from dataclasses import dataclass
from os import path
import re


@dataclass
class Instruction():
    amount: int
    origin: int
    to: int


def parse_stacks(stacks_raw: list[str]) -> list[list[str]]:
    # saves the stacks from 0 to 8 (1 - 9 in the instructions)
    stacks = [[], [], [], [], [], [], [], [], []]
    for line in stacks_raw:
        start_index = 0  # starting for searching "["
        stack_to_fill = 0  # wich stack it is supposed to fill with content
        end_line = False  # if the contents from the line are done
        # loops the process of finds "[" -> fill the gaps -> fill the letter
        while True:
            index = 0  # declares the index where "[" is found
            try:
                index = line[start_index:].index("[")
            except:  # if none "[" is found, it means that there's no more letters in the line
                # fills the remaining spaces with "-" (blanks)
                for _ in range(9-stack_to_fill):
                    stacks[stack_to_fill].append("-")
                    stack_to_fill += 1
                    end_line = True
            if end_line:
                break
            empty_stacks = index/4  # finds how many stacks are missing from the last spot searched
            # if there's no missing stack, then the spaces (1) divided by 4 is lesser than one (0.25)
            if not (empty_stacks < 1):
                empty_stacks = int(empty_stacks)
                for _ in range(empty_stacks):  # fills the empty spaces
                    stacks[stack_to_fill].append("-")
                    stack_to_fill += 1
            if stack_to_fill > 8:
                break
            # fills the letter that was found in the beginning with the "["
            stacks[stack_to_fill].append(line[index+start_index+1])
            stack_to_fill += 1
            # updates the start index so the next search starts from the "]"
            start_index = start_index + index + 2

    for stack in stacks:
        stack.reverse()
        while "-" in stack:
            stack.remove("-")

    return stacks


def parse_instruction(line: str) -> Instruction:
    result = re.search(r"move\s(\d*)\sfrom\s(\d*)\sto\s(\d*)", line)
    return Instruction(int(result.group(1)), int(result.group(2)), int(result.group(3)))


def move_stacks(instruction: Instruction, stacks: list[list[str]]) -> list[list[str]]:
    load = stacks[instruction.origin-1][-instruction.amount:]
    stacks[instruction.origin-1] = stacks[instruction.origin -
                                          1][:len(stacks[instruction.origin-1])-instruction.amount]
    load.reverse()
    stacks[instruction.to-1].extend(load)
    return stacks


file_path = path.join(path.dirname(__file__), "input.txt")

with open(file_path) as file:

    done_reading_satcks = False

    stacks_raw = []

    stacks = []

    while True:  # reading stacks

        line = file.readline()
        if not line:
            break
        else:
            line = line.strip("\n")

        if line.startswith(" 1"):
            file.readline()
            line = file.readline()
            line = line.strip("\n")
            stacks = parse_stacks(stacks_raw)
            done_reading_satcks = True

        if not done_reading_satcks:
            stacks_raw.append(line)

        else:
            instruction = parse_instruction(line)
            stacks = move_stacks(instruction, stacks)

    for stack in stacks:
        print(stack[-1])
