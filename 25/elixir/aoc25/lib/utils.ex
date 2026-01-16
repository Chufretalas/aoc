defmodule Aoc25.Utils do
  @spec parse_map_matrix(String.t()) :: {integer(), integer(), map()}
  def parse_map_matrix(input_path) do
    lines =
      File.stream!(input_path)
      |> Stream.map(&String.trim/1)
      |> Enum.to_list()

    num_lines = length(lines)
    num_cols = String.length(hd(lines))

    puzzle =
      for {line, r} <- Enum.with_index(lines),
          {char, c} <- Enum.with_index(String.graphemes(line)),
          into: %{} do
        {{r, c}, char}
      end

    {num_lines, num_cols, puzzle}
  end

  def debug_print_map_matrix(num_rows, num_cols, matrix) do
    for r <- 0..(num_rows - 1) do
      for c <- 0..(num_cols - 1) do
        Map.get(matrix, {r, c})
      end
    end
    |> IO.inspect()
  end
end
