from __future__ import annotations
from dataclasses import dataclass
from math import floor

@dataclass()
class Monkey():
    items: list[int]
    operation_str: str
    test_number: int
    true_monkey: int
    false_monkey: int
    times_tested = 0

    def throw_item(self, receiver: Monkey):
        item = self.items[0]
        self.items.pop(0)
        receiver.items.append(item)

    def test(self, item: int) -> bool:
        self.times_tested += 1

        worry_level = item
        operator, value = self.operation_str.split(" ")
        if value == "old":
            value = item
        else:
            value = int(value)

        match operator:
            case "+":
                worry_level += value
            case "*":
                worry_level *= value

        worry_level = floor(worry_level/3)

        self.items[0] = worry_level

        return worry_level%self.test_number == 0

    def __str__(self) -> str:
        return f"(items: {len(self.items)} | operation: {self.operation_str} | test: {self.test_number} | true: {self.true_monkey} | false: {self.false_monkey})"
