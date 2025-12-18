defmodule Days.Day03 do
  def part1() do
    File.stream!("inputs/day03.txt")
    |> Stream.map(&String.trim/1)
    |> Stream.map(&String.graphemes/1)
    |> Stream.map(&Enum.with_index/1)
    |> Stream.map(fn indexed_chars ->
      # This returns teh left-most max value in the sequence
      {left, left_index} = Enum.max_by(indexed_chars, fn {value, _idx} -> value end)

      len = length(indexed_chars)

      if left_index == len - 1 do
        # The last item is the biggest, so search the list, minus the last item, for the new (actual) left
        {actual_left, _} =
          Enum.max_by(Enum.slice(indexed_chars, 0..(left_index - 1)), fn {value, _idx} ->
            value
          end)

        String.to_integer(actual_left <> left)
      else
        {right, _} =
          Enum.max_by(Enum.slice(indexed_chars, (left_index + 1)..len), fn {value, _idx} ->
            value
          end)

        String.to_integer(left <> right)
      end
    end)
    |> Enum.sum()
  end

  def part2() do
    File.stream!("inputs/day03.txt")
    |> Stream.map(&String.trim/1)
    |> Stream.map(&String.graphemes/1)
    |> Stream.map(&Enum.with_index/1)
    |> Stream.map(fn indexed_chars ->
      len = length(indexed_chars)

      initial_last_idx = -1

      {digits_list, _final_idx} =
        Enum.map_reduce(12..1//-1, initial_last_idx, fn nums_left, last_chosen_idx ->
          start_search = last_chosen_idx + 1 # Start searching after the last picked number
          end_search = len - nums_left # Leave enough numbers to complete the required sequence

          window = Enum.slice(indexed_chars, start_search..end_search)

          {max_value, found_idx} = Enum.max_by(window, fn {value, _} -> value end)
        
          # The first element will go to the mapped list and the second will update the accumulator
          {max_value, found_idx}
        end)

      digits_list
      |> Enum.join("")
      |> String.to_integer()
    end)
    |> Enum.sum()
  end
end
