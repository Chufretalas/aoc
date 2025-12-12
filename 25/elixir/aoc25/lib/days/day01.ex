defmodule Aoc25.Day01 do
  def part1 do
    File.stream!("inputs/day01.txt")
    |> Enum.map(&String.trim/1)
    |> Enum.map(&String.split_at(&1, 1))
    |> Enum.reduce({50, 0}, fn {direction, amount_str}, {dial, zero_passes} ->
      amount = String.to_integer(amount_str)

      new_dial =
        case direction do
          "L" -> dial - amount
          "R" -> dial + amount
        end

      new_zero_passes =
        if rem(new_dial, 100) == 0 do
          zero_passes + 1
        else
          zero_passes
        end

      {new_dial, new_zero_passes}
    end)
    |> elem(1)
  end

  def part2 do
    File.stream!("inputs/day01.txt")
    |> Enum.map(&String.trim/1)
    |> Enum.map(&String.split_at(&1, 1))
    |> Enum.reduce({50, 0}, fn {direction, amount_str}, {dial, total_zero_passes} ->
      amount = String.to_integer(amount_str)

      {new_dial, additional_zero_passes} =
        Enum.reduce(0..(amount - 1), {dial, 0}, fn _, {dial, zero_passes} ->
          new_dial =
            case direction do
              "L" -> dial - 1
              "R" -> dial + 1
            end

          new_zero_passes =
            if rem(new_dial, 100) == 0 do
              zero_passes + 1
            else
              zero_passes
            end

          {new_dial, new_zero_passes}
        end)

      {new_dial, total_zero_passes + additional_zero_passes}
    end)
    |> elem(1)
  end
end
