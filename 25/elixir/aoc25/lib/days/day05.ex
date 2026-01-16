defmodule Aoc25.Day05 do
  defp is_in_range?({l, r}, number), do: l <= number and number <= r

  def part1() do
    {ranges, availables} =
      File.read!("inputs/day05.txt")
      |> String.split("\n")
      |> Enum.split_while(fn line -> line != "" end)
      |> then(fn {ranges, [_empty_line | availables]} -> {ranges, availables} end)

    ranges =
      Enum.map(ranges, fn range_str ->
        [l, r] = String.split(range_str, "-")
        {String.to_integer(l), String.to_integer(r)}
      end)

    Enum.reduce(availables, 0, fn available_str, acc ->
      if Enum.any?(ranges, fn range -> is_in_range?(range, String.to_integer(available_str)) end) do
        acc + 1
      else
        acc
      end
    end)
  end

  def part2() do
    ranges =
      File.read!("inputs/day05.txt")
      |> String.split("\n")
      |> Enum.split_while(fn line -> line != "" end)
      |> then(fn {ranges_str, _} -> ranges_str end)
      |> Enum.map(fn range_str ->
        [l, r] = String.split(range_str, "-")
        {String.to_integer(l), String.to_integer(r)}
      end)
      |> Enum.sort(fn {l1, _}, {l2, _} -> l1 <= l2 end)

    # Merging ranges
    merged_ranges = Enum.reduce(ranges, [], fn {crr_l, crr_r}, acc ->
      case acc do
        [] ->
          [{crr_l, crr_r}]

        [{last_l, last_r} | other_ranges] ->
          if is_in_range?({last_l, last_r}, crr_l) do
            # the new right is bigger than the old right
            if crr_r > last_r do
              # then substitute the range with the largest intersection
              # |~~~~|
              #   |~~~~|
              # |~~~~~~|
              [{last_l, crr_r} | other_ranges]
            else
              # |~~~~~~|
              #   |~~|
              # |~~~~~~|
              [{last_l, last_r} | other_ranges]
            end
          else
            # Else put this range as the new last
            # |~~~~|
            #         |~~~~|
            # |~~~~|  |~~~~|
            [{crr_l, crr_r} | acc]
          end
      end
    end)
    
    Enum.reduce(merged_ranges, 0, fn {l, r}, acc -> acc + r - l + 1 end)
  end
end
