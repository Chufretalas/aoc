defmodule Aoc25.Day06 do
  defp compute_operation("+", operands), do: Enum.sum(operands)
  defp compute_operation("*", operands), do: Enum.product(operands)

  def part1() do
    # operations = [["+", 1, 2, 3], ["*", 4, 5, 6]]
    operations =
      File.stream!("inputs/day06.txt")
      |> Stream.map(&String.trim/1)
      |> Stream.map(&String.split(&1, ~r/\s+/))
      |> Enum.reverse()
      |> Enum.zip()
      |> Enum.map(&Tuple.to_list/1)
      |> Enum.map(fn [operator | operands] ->
        [operator | Enum.map(operands, &String.to_integer/1)]
      end)

    Enum.reduce(operations, 0, fn [operator | operands], acc ->
      acc + compute_operation(operator, operands)
    end)
  end

  # I'm so sorry for writing this mess, but it works and it's fast, so I am not going to bother with refactoring 🙇‍♂️
  def part2() do
    [operators_line | operand_lines_inverted] =
      File.read!("inputs/day06.txt")
      |> String.split("\n", trim: true)
      |> Enum.map(&String.reverse/1)
      |> Enum.reverse()

    boundaries =
      for {char, index} <- Enum.with_index(String.graphemes(operators_line)),
          char == "+" or char == "*",
          do: index

    operand_lines = Enum.reverse(operand_lines_inverted)

    {split_ranges, _} =
      Enum.map_reduce(boundaries, 0, fn boundary, start_index ->
        {start_index..boundary, boundary + 1}
      end)

    operands =
      operand_lines
      |> Enum.map(fn line ->
        Enum.map(split_ranges, fn range ->
          String.slice(line, range)
        end)
        |> Enum.map(&String.reverse/1)
      end)
      |> Enum.zip()
      |> Enum.map(&Tuple.to_list/1)
      |> Enum.map(fn numbers_str ->
        # ["  1", "  22", "333"]
        Enum.map(numbers_str, &String.graphemes/1)
        # [["", "", "1"], ["", "2", "2"], ["3", "3", "3"]]
        |> Enum.zip()
        # [{"", "", "3"}, {"", "2", "3"}, {"1", "2", "3"}]
        |> Enum.map(&Tuple.to_list/1)
        |> Enum.reject(&Enum.all?(&1, fn e -> e in ["", " "] end)) # Remove columns that are all whitespace
        |> Enum.map(fn list ->
          Enum.join(list)
          |> String.trim()
          |> String.to_integer()
        end)
      end)

    operators =
      String.split(operators_line, ~r/\s+/)
      |> Enum.reject(&(&1 == ""))

    operations = Enum.zip(operators, operands)

    Enum.reduce(operations, 0, fn {operator, operands}, acc ->
      acc + compute_operation(operator, operands)
    end)
  end
end
