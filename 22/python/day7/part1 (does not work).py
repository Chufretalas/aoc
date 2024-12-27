from __future__ import annotations
from dataclasses import dataclass, field
import os

@dataclass
class File():
    name: str
    size: int


@dataclass()
class Dir():
    name: str
    parent_dir: str
    files: list[File] = field(default_factory=list)
    sub_dirs: list[str] = field(default_factory=list)  # just the other dirs names

    def get_total_size(self, other_dirs: list[Dir]) -> int:
        size = 0
        for file in self.files:
            size += file.size
            if size > 100000: return size

        for dir_name in self.sub_dirs:
            size += other_dirs[dir_name].get_total_size(other_dirs)
            if size > 100000: break
        return size

        


def main():
    input_path = os.path.join(os.path.dirname(__file__), "input.txt")
    with open(input_path, "r") as file:

        location = ["/"]
        dirs: dict[str, Dir] = {
            "/": Dir(name="Dir", parent_dir=None)
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
                    location = location[:1]

                else:
                    location.append(command)
                    if command not in dirs.keys():
                        dirs[command] = Dir(name=command, parent_dir=location[-2])

            elif line.startswith("$ ls"):
                pass # I dont think I need this

            else:
                if line[0].isnumeric():
                    size, name = line.split(" ")
                    size = int(size)
                    dirs[location[-1]].files.append(File(name, size))

                else:
                    dirs[location[-1]].sub_dirs.append(line.split(" ")[1])

        final_sum = 0
        for name, dir in dirs.items():
            size = dir.get_total_size(dirs)
            if size <= 100_000:
                print(name, size)
                final_sum += size

        print(final_sum)
        file.close()


if __name__ == "__main__":
    main()
