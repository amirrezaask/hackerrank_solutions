defmodule Solution do
  def remainder_string(str_graphemes, n) do
    if rem(n, length(str_graphemes)) == 0 do
      ""
    else
      Enum.join(Enum.slice(str_graphemes, 0..(rem(n, length(str_graphemes)) - 1)), "")
    end
  end

  def count(str, n, char) do
    c = find_char_count(str, char)
    l = floor(n / String.length(str))
    remainder = String.graphemes(str) |> remainder_string(n) |> find_char_count(char)

    c * l + remainder
  end

  def find_char_count(str, char) do
    str
    |> String.graphemes()
    |> Enum.reduce(0, fn c, acc ->
      if c == char do
        acc + 1
      else
        acc
      end
    end)
  end
end

s = IO.read(:stdio, :line) |> String.trim()

{n, _base} = IO.read(:stdio, :line) |> Integer.parse()

s
|> Solution.count(n, "a")
|> IO.inspect()
