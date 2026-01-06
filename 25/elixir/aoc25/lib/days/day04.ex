defmodule Days.Day04 do
  def part1() do
    {num_lines, num_cols, puzzle} =
      File.stream!("inputs/day04.txt")
      |> Stream.map(&String.trim/1)
      |> Stream.map(&String.graphemes/1)
      |> Stream.with_index()
      |> Enum.reduce({0, 0, %{}}, fn {split_line, line_idx}, {num_lines, _num_cols, puzzle} ->
        {num_cols, puzzle_line} =
          Enum.with_index(split_line)
          |> Enum.reduce({0, %{}}, fn {crr, col_idx}, {num_cols, puzzle_line} ->
            {num_cols + 1, Map.put(puzzle_line, {line_idx, col_idx}, crr)}
          end)

        {num_lines + 1, num_cols, Map.merge(puzzle, puzzle_line)}
      end)

    rolls =
      for l <- 0..(num_lines - 1),
          c <- 0..(num_cols - 1),
          puzzle[{l, c}] == "@",
          do: {l, c}

    Enum.reduce(rolls, 0, fn {roll_l, roll_c}, accessible_rolls ->
      adjacent_rolls =
        for l <- (roll_l - 1)..(roll_l + 1),
            c <- (roll_c - 1)..(roll_c + 1),
            l >= 0 && l < num_lines && c >= 0 && c < num_cols && (l != roll_l || c != roll_c) do
          {l, c}
        end

        # Generates all adjacent positions for the roll
        |> Enum.count(&(puzzle[&1] == "@"))

      if adjacent_rolls < 4 do
        accessible_rolls + 1
      else
        accessible_rolls
      end
    end)
  end

  defp isAccessible?(num_lines, num_cols, puzzle, roll_l, roll_c) do
    adjacent_rolls =
      for l <- (roll_l - 1)..(roll_l + 1),
          c <- (roll_c - 1)..(roll_c + 1),
          l >= 0 && l < num_lines && c >= 0 && c < num_cols && (l != roll_l || c != roll_c) do
        {l, c}
      end
      |> Enum.count(&(puzzle[&1] == "@"))

    adjacent_rolls < 4
  end

  @spec roll_remover(integer(), integer(), map()) :: {map(), integer()}
  defp roll_remover(num_lines, num_cols, puzzle) do
    rolls =
      for l <- 0..(num_lines - 1),
          c <- 0..(num_cols - 1),
          puzzle[{l, c}] == "@" && isAccessible?(num_lines, num_cols, puzzle, l, c),
          do: {l, c}

    removed_rolls = length(rolls)

    new_puzzle =
      Map.new(puzzle, fn {key, value} ->
        if Enum.member?(rolls, key) do
          {key, "."}
        else
          {key, value}
        end
      end)

    {new_puzzle, removed_rolls}
  end

  defp part2_loop(num_lines, num_cols, puzzle, total_removed_rolls \\ 0) do
    {new_puzzle, removed_rolls} = roll_remover(num_lines, num_cols, puzzle)

    if removed_rolls == 0 do
      total_removed_rolls
    else
      part2_loop(num_lines, num_cols, new_puzzle, total_removed_rolls + removed_rolls)
    end
  end

  def part2() do
    {num_lines, num_cols, puzzle} =
      File.stream!("inputs/day04.txt")
      |> Stream.map(&String.trim/1)
      |> Stream.map(&String.graphemes/1)
      |> Stream.with_index()
      |> Enum.reduce({0, 0, %{}}, fn {split_line, line_idx}, {num_lines, _num_cols, puzzle} ->
        {num_cols, puzzle_line} =
          Enum.with_index(split_line)
          |> Enum.reduce({0, %{}}, fn {crr, col_idx}, {num_cols, puzzle_line} ->
            {num_cols + 1, Map.put(puzzle_line, {line_idx, col_idx}, crr)}
          end)

        {num_lines + 1, num_cols, Map.merge(puzzle, puzzle_line)}
      end)

    part2_loop(num_lines, num_cols, puzzle)
  end
end
