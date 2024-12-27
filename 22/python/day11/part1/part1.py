import os
from Monkey import Monkey
from parse_monkey import parse_monkey

def main():
    input_path = os.path.join(os.path.dirname(__file__), "input.txt")
    with open(input_path, "r") as file:

        lines = []
        monkeys: list[Monkey] = []

        while True:
            line = file.readline()
            if not line:
                break

            line = line.strip()

            if line != "":
                lines.append(line)

            if len(lines) == 6:
                new_monkey = parse_monkey(lines)
                monkeys.append(new_monkey)
                lines = []

        for _ in range(20): # rounds
            for monkey in monkeys:
                items = monkey.items.copy()
                for item in items:
                    result = monkey.test(item)
                    if result:
                        monkey.throw_item(monkeys[monkey.true_monkey])
                    else:
                        monkey.throw_item(monkeys[monkey.false_monkey])

        monkey_business_values = []
        for m in monkeys:
            monkey_business_values.append(m.times_tested)
        monkey_business_values.sort()
        
        monkey_business = monkey_business_values[-1] * monkey_business_values[-2]

        print(monkey_business)

        file.close()


if __name__ == "__main__":
    main()
