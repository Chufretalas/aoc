defmodule Aoc25.Day07 do
  defp selective_tachyon_merge(original_map, new_map, insert_map) do
    Enum.reduce(insert_map, new_map, fn {coord, value}, new_map ->
      char = Map.get(original_map, coord)

      if char == "^" do
        Map.put(new_map, coord, char)
      else
        Map.put(new_map, coord, value)
      end
    end)
  end

  def part1() do
    {num_rows, num_cols, original_puzzle} = Aoc25.Utils.parse_map_matrix("inputs/day07.txt")

    iterator =
      for r <- 1..(num_rows - 1), c <- 0..(num_cols - 1) do
        {r, c}
      end

    # copy the first line
    new_puzzle =
      for c <- 0..(num_cols - 1), into: %{} do
        {{0, c}, Map.get(original_puzzle, {0, c})}
      end

    {final_puzzle, splits} =
      Enum.reduce(iterator, {new_puzzle, 0}, fn {r, c}, {new_puzzle, splits} ->
        char = Map.get(original_puzzle, {r, c})
        # if there's a tachyon beam above
        if Map.get(new_puzzle, {r - 1, c}) in ["S", "|"] do
          if char == "^" do
            {
              selective_tachyon_merge(original_puzzle, new_puzzle, %{
                {{r, c - 1}, "|"},
                {{r, c}, char},
                {{r, c + 1}, "|"}
              }),
              splits + 1
            }
          else
            {Map.put(new_puzzle, {r, c}, "|"), splits}
          end
        else
          current_val = Map.get(new_puzzle, {r, c})

          if current_val == "|" do
            {new_puzzle, splits}
          else
            {Map.put(new_puzzle, {r, c}, char), splits}
          end
        end
      end)

    Aoc25.Utils.debug_print_map_matrix(num_rows, num_cols, final_puzzle)

    splits
  end

  defp update_count(map, key, amount) do
    Map.update(map, key, amount, fn current -> current + amount end)
  end

  # (DP, Pascal Triangle)
  # The ideia is to go row by row and keep track of how many different paths lead to each new path
  # the new paths created by a split inherits the number of paths of the source
  # if multiple splits lead to the same path, that new path combines the multiple inhetances
  #   S : S = 1
  #  / \
  # A   B : A = 1, B = 1
  #  \ /
  #   C : C = A + B = 2 
  #   |
  #   D : D = 2
  def part2() do
    {num_rows, _num_cols, grid} = Aoc25.Utils.parse_map_matrix("inputs/day07.txt")

    {{start_r, start_c}, _} = Enum.find(grid, fn {_, char} -> char == "S" end)

    # %{ column_index => count_of_timelines }
    initial_timelines = %{start_c => 1}

    final_timelines =
      Enum.reduce(start_r..(num_rows - 1), initial_timelines, fn r, current_timelines ->
        Enum.reduce(current_timelines, %{}, fn {c, count}, next_acc ->
          char = Map.get(grid, {r, c})

          case char do
            "^" ->
              next_acc
              |> update_count(c - 1, count)
              |> update_count(c + 1, count)

            _ ->
              update_count(next_acc, c, count)
          end
        end)
      end)

    final_timelines
    |> Map.values()
    |> Enum.sum()
  end
end
