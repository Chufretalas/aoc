# inspired by u/junefish on reddit

from __future__ import annotations
from dataclasses import dataclass
from typing import Union
import os

@dataclass()
class Dir():
    name: str
    parent_dir: Union[None, str]
    size: int = 0


def main():
    input_path = os.path.join(os.path.dirname(__file__), "input.txt")
    with open(input_path, "r") as file:

        location = ["/"]
        dirs: dict[str, Dir] = {
            "/": Dir(name="/", parent_dir=None)
        }

        while True:
            line = file.readline()
            if not line:
                break

            line = line.strip("\n")

            if line.startswith("$ cd"):
                command = line.split(" ")[-1]

                if command == "..":
                    location.pop()

                elif command == "/":
                    location = ["/"]

                else:
                    location.append(command)
                    if command not in dirs.keys():
                        dirs["/".join(location)] = Dir(name=command, parent_dir=location[-2])

            else:
                if line[0].isdigit():
                    size, name = line.split(" ")
                    size = int(size)
                    for i in range(len(location)):
                        dirs["/".join(location[:i+1])].size += size

        final_sum = 0
        for name, dir in dirs.items():
            size = dir.size
            if size <= 100_000:
                print(name, size)
                final_sum += size

        print(final_sum)
        file.close()


if __name__ == "__main__":
    main()
