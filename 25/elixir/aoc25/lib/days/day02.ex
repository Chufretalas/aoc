defmodule Days.Day02 do
  def part1() do
    File.read!("inputs/day02.txt")
    |> String.split(",")
    |> Enum.reduce(0, fn range_str, total_sum ->
      [low, high] = String.split(range_str, "-")

      total_sum +
        Enum.reduce(String.to_integer(low)..String.to_integer(high), 0, fn candidate,
                                                                           partial_sum ->
          candidate_str = Integer.to_string(candidate)
          candidate_str_len = String.length(candidate_str)

          if rem(candidate_str_len, 2) == 1 do
            partial_sum
          else
            {left, right} = String.split_at(candidate_str, div(candidate_str_len, 2))

            if left == right do
              partial_sum + candidate
            else
              partial_sum
            end
          end
        end)
    end)
  end

  defp split_by_length(str, len) do
    str
    |> String.graphemes()
    |> Enum.chunk_every(len)
    |> Enum.map(&Enum.join/1)
  end

  def part2() do
    File.read!("inputs/day02.txt")
    |> String.split(",")
    |> Enum.reduce(0, fn range_str, total_sum ->
      [low, high] = String.split(range_str, "-")

      # Checking the ids inside the ranges
      total_sum +
        Enum.reduce(
          String.to_integer(low)..String.to_integer(high),
          0,
          fn candidate, partial_sum ->
            candidate_str = Integer.to_string(candidate)
            candidate_str_len = String.length(candidate_str)

            # Cheking if by slicing the candidate at every possible interval leads to all slices being equal at at least on instance
            if candidate_str_len < 2 do
              partial_sum
            else
              is_invalid_id =
                1..div(candidate_str_len, 2)
                # The interval must lead to equal-sized slices
                |> Enum.filter(fn chunk_size ->
                  chunk_size > 0 && rem(candidate_str_len, chunk_size) == 0
                end)
                |> Enum.any?(fn chunck_size ->
                  split = split_by_length(candidate_str, chunck_size)
                  [first_elem | _] = split
                  # If every slice is equal to the first
                  Enum.all?(split, &(&1 == first_elem))
                end)

              if is_invalid_id do
                partial_sum + candidate
              else
                partial_sum
              end
              
            end
          end
        )
    end)
  end
end
